package globals

import (
	"log"
	"os"
	"reflect"
	"strings"
)

type EnvType struct {
	REST_API_KEY string
	REDIRECT_URI string
	DB_URL       string
	ACCOUNT_ID   string
	ACCESS_KEY   string
	SECRET_KEY   string
	BUCKET_NAME  string
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
		log.Printf("%s=%s", key, value)

		reflect.ValueOf(&Env).Elem().FieldByName(key).SetString(value)
	}

	return nil
}
