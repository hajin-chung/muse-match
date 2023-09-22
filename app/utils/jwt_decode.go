package utils

import (
	"encoding/base64"
	"encoding/json"
	"strings"
)

func DecodeJWT(jwt string, v any) error {
	parts := strings.Split(jwt, ".")
	jwtBytes, err := base64.StdEncoding.DecodeString(parts[1])
	if err != nil {
		return err
	}
	err = json.Unmarshal(jwtBytes, v)
	if err != nil {
		return err
	}
	return nil
}
