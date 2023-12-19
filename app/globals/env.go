package globals

import (
	"log"
	"os"
	"reflect"
	"strings"
)

type EnvType struct {
	MODE          string // DEV | PROD
	REST_API_KEY  string // kakao oauth rest api key
	REDIRECT_URI  string // kakao oauth redirect uri
	DB_URL        string // sqlite db url
	ACCESS_KEY    string // aws access key
	SECRET_KEY    string // aws secret key
	BUCKET_NAME   string // aws s3 bucket name
	ADMIN         string // admin session secret token
	SLACK_WEBHOOK string // slack webhook url
	LOG_FILE      string // log file path
}

var Env EnvType

func LoadEnv() error {
	data, err := os.ReadFile(".env")
	if err != nil {
		return err
	}
	content := string(data)
	for _, line := range strings.Split(content, "\n") {
		parts := strings.Split(line, "=")
		if len(parts) != 2 {
			continue
		}
		key := parts[0]
		value := parts[1]
		log.Printf("%s=\"%s\"", key, value)

		reflect.ValueOf(&Env).Elem().FieldByName(key).SetString(value)
	}

	return nil
}
