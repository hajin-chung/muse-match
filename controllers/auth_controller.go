package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"musematch/globals"
	"musematch/models"
	"musematch/queries"
	"musematch/utils"
	"musematch/views/pages"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func LoginController(c *fiber.Ctx) error {
	kakaoUrl := fmt.Sprintf(
		"https://kauth.kakao.com/oauth/authorize?client_id=%s&redirect_uri=%s&response_type=code",
		globals.Env.KAKAO_API_KEY,
		globals.Env.KAKAO_REDIRECT_URI,
	)

	randomState := rand.Intn(100)
	naverUrl := fmt.Sprintf(
		"https://nid.naver.com/oauth2.0/authorize?client_id=%s&redirect_uri=%s&state=%d&response_type=code",
		globals.Env.NAVER_CLIENT_ID,
		globals.Env.NAVER_REDIRECT_URI,
		randomState,
	)

	googleUrl := ""

	page := pages.Login("this is title", kakaoUrl, naverUrl, googleUrl)
	return utils.Render(c, page)
}

type KakaoJWT struct {
	Token string                 `json:"id_token"`
	X     map[string]interface{} `json:"-"`
}

type KakaoToken struct {
	Sub      string `json:"sub"`
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
	Picture  string `json:"picture"`
}

func KakaoCallbackController(c *fiber.Ctx) error {
	sess, _ := globals.Store.Get(c)

	u, _ := url.Parse(c.OriginalURL())

	code := u.Query().Get("code")
	tokenBody := map[string]string{
		"grant_type":   "authorization_code",
		"client_id":    globals.Env.KAKAO_API_KEY,
		"redirect_uri": globals.Env.KAKAO_REDIRECT_URI,
		"code":         code,
	}
	tokenEncoded := ""
	for key, value := range tokenBody {
		tokenEncoded += fmt.Sprintf("%s=%s&", url.QueryEscape(key), url.QueryEscape(value))
	}
	tokenEncoded = strings.TrimRight(tokenEncoded, "&")

	res, err := http.Post(
		"https://kauth.kakao.com/oauth/token",
		"application/x-www-form-urlencoded",
		bytes.NewBuffer([]byte(tokenEncoded)),
	)
	if err != nil {
		log.Println("Warning kakao token post req error")
		return err
	}
	defer res.Body.Close()

	var jwt KakaoJWT
	data, _ := io.ReadAll(res.Body)
	_ = json.Unmarshal(data, &jwt)

	var token KakaoToken
	utils.DecodeJWT(jwt.Token, &token)

	// TODO: implement this
	// TODO: handle err
	user, _ := queries.GetUserBySub(token.Sub)
	if user == nil { // create user
		user = &models.User{
			Id:          utils.CreateId(),
			Name:        token.Nickname,
			Email:       token.Email,
			Sub:         token.Sub,
			Picture:     token.Picture,
			Description: "",
			Note:        "",
			InstagramId: "",
			FacebookId:  "",
			TwitterId:   "",
		}
		// TODO: handle err
		err = queries.CreateUser(user)
		if err != nil {
			log.Println("ERROR on create user: ", err)
			return c.Redirect("/")
			// return err
		}
	}
	sess.Set("id", user.Id)
	sess.Save()

	return c.Redirect("/")
}

type NaverToken struct {
	AccessToken string `json:"access_token"`
}

type NaverUser struct {
	Email   string `json:"email"`
	Name    string `json:"name"`
	Id      string `json:"id"`
	Picture string `json:"profile_image"`
}

type NaverAPI struct {
	User NaverUser `json:"response"`
}

func NaverCallbackController(c *fiber.Ctx) error {
	u, _ := url.Parse(c.OriginalURL())

	code := u.Query().Get("code")
	randomState := rand.Intn(100)
	fields := map[string]string{
		"grant_type":    "authorization_code",
		"client_id":     globals.Env.NAVER_CLIENT_ID,
		"client_secret": globals.Env.NAVER_SECRET,
		"redirect_uri":  globals.Env.NAVER_REDIRECT_URI,
		"code":          code,
		"state":         strconv.Itoa(randomState),
	}
	queryString := ""
	for key, value := range fields {
		queryString += fmt.Sprintf("%s=%s&", url.QueryEscape(key), url.QueryEscape(value))
	}
	queryString = strings.TrimRight(queryString, "&")
	apiUrl := "https://nid.naver.com/oauth2.0/token?grant_type=authorization_code&" + queryString

	res, err := http.Get(apiUrl)
	if err != nil {
		log.Println("Warning naver token get req error")
		return err
	}
	defer res.Body.Close()

	var token NaverToken
	data, _ := io.ReadAll(res.Body)
	_ = json.Unmarshal(data, &token)

	// fetch user info
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "https://openapi.naver.com/v1/nid/me", nil)
	req.Header.Set("Authorization", "Bearer "+token.AccessToken)
	res, err = client.Do(req)
	if err != nil {
		log.Println("failed to fetch user info naver openapi")
	}

	var apiResponse NaverAPI
	data, _ = io.ReadAll(res.Body)
	_ = json.Unmarshal(data, &apiResponse)

	return nil
}
