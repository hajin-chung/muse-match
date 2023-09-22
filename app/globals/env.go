package globals

import (
	"os"
	"reflect"
	"strings"
)

type EnvType struct {
	REST_API_KEY string
	REDIRECT_URI string
	DB_URL       string
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
		key := parts[0]
		value := parts[1]

		reflect.ValueOf(&Env).Elem().FieldByName(key).SetString(value)
	}

	return nil
}
