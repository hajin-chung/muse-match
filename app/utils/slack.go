package utils

import (
	"bytes"
	"encoding/json"
	"musematch/app/globals"
	"net/http"
)

func SlackSendMessage(message string) error {
	body := map[string]string{"text": message}
	json, _ := json.Marshal(body)
	_, err := http.Post(globals.Env.SLACK_WEBHOOK, "application/json", bytes.NewBuffer(json))
	return err
}
