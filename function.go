package memoirnotification

import (
	"encoding/json"
	"net/http"

	expo "github.com/oliveroneill/exponent-server-sdk-golang/sdk"
)

type RequestParam struct {
	Token     string `json:"token"`
	Title     string `json:"title"`
	Body      string `json:"body"`
	URLScheme string `json:"urlScheme"`
}

func SendNotification(w http.ResponseWriter, r *http.Request) {
	param := RequestParam{}

	if err := json.NewDecoder(r.Body).Decode(&param); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	pushToken, err := expo.NewExponentPushToken(param.Token)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	to := []expo.ExponentPushToken{pushToken}

	client := expo.NewPushClient(nil)

	_, err = client.Publish(
		&expo.PushMessage{
			To:       to,
			Body:     param.Body,
			Data:     map[string]string{"urlScheme": param.URLScheme},
			Sound:    "default",
			Title:    param.Title,
			Priority: expo.DefaultPriority,
		},
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
