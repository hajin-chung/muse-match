package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"musematch/app/globals"
	"musematch/app/models"
	"musematch/app/queries"
	"musematch/app/utils"
	"net/http"
	"net/url"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func LoginController(c *fiber.Ctx) error {
	kakaoHref := fmt.Sprintf(
		"https://kauth.kakao.com/oauth/authorize?client_id=%s&redirect_uri=%s&response_type=code",
		globals.Env.REST_API_KEY,
		globals.Env.REDIRECT_URI,
	)

	return c.Render("pages/login", fiber.Map{
		"Title": "로그인",
		"Kakao": fiber.Map{
			"Href": kakaoHref,
		},
	}, "layout")
}

func LogoutController(c *fiber.Ctx) error {
	// TODO: handle error
	sess, _ := globals.Store.Get(c)
	sess.Destroy()

	return c.Redirect("/")
}

type JWT struct {
	Token string                 `json:"id_token"`
	X     map[string]interface{} `json:"-"`
}

type Token struct {
	Sub      string `json:"sub"`
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
	Picture  string `json:"picture"`
}

func KakaoCallbackController(c *fiber.Ctx) error {
	// TODO: handle error
	sess, err := globals.Store.Get(c)

	u, _ := url.Parse(c.OriginalURL())

	code := u.Query().Get("code")
	tokenBody := map[string]string{
		"grant_type":   "authorization_code",
		"client_id":    globals.Env.REST_API_KEY,
		"redirect_uri": globals.Env.REDIRECT_URI,
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

	var jwt JWT
	data, err := io.ReadAll(res.Body)
	err = json.Unmarshal(data, &jwt)

	var token Token
	utils.DecodeJWT(jwt.Token, &token)

	// TODO: handle err
	user, err := queries.GetUserBySub(token.Sub)
	if (user == models.User{}) { // create user
		user = models.User{
			Id:          utils.CreateId(),
			Name:        token.Nickname,
			Email:       token.Email,
			Sub:         token.Sub,
			Picture:     token.Picture,
			Description: "",
			History:     "",
		}
		// TODO: handle err
		err = queries.CreateUser(&user)
		if err != nil {
			log.Println("ERROR on create user", err)
		}
	}
	log.Printf("%s\n", user.Id)
	sess.Set("id", user.Id)
	sess.Save()

	return c.Redirect("/")
}
