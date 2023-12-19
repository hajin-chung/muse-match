package utils

import (
	"bytes"
	"encoding/json"
	"log"
	"musematch/globals"
	"net/http"
)

func SlackSendMessage(message string) error {
	if globals.Env.MODE != "PROD" {
		log.Println(message)
		return nil
	}
	body := map[string]string{"text": message}
	json, _ := json.Marshal(body)
	_, err := http.Post(globals.Env.SLACK_WEBHOOK, "application/json", bytes.NewBuffer(json))
	return err
}
