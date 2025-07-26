package memoirnotification

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	expo "github.com/oliveroneill/exponent-server-sdk-golang/sdk"
)

type NotificationRequest struct {
	Token     []string `json:"token"`
	Title     string   `json:"title"`
	Body      string   `json:"body"`
	URLScheme string   `json:"urlScheme"`
}

func SendNotification2(w http.ResponseWriter, r *http.Request) {
	param := NotificationRequest{}

	if err := json.NewDecoder(r.Body).Decode(&param); err != nil {
		fmt.Println("error1:", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Printf("params: %s\n", strings.Join(param.Token, ","))

	to := []expo.ExponentPushToken{}
	for _, token := range param.Token {
		pushToken, err := expo.NewExponentPushToken(token)
		if err != nil {
			fmt.Println("error2:", err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		to = append(to, pushToken)

	}

	client := expo.NewPushClient(nil)

	for _, token := range to {
		_, err := client.Publish(
			&expo.PushMessage{
				To:       []expo.ExponentPushToken{token},
				Body:     param.Body,
				Data:     map[string]string{"urlScheme": param.URLScheme},
				Sound:    "default",
				Title:    param.Title,
				Priority: expo.DefaultPriority,
			},
		)
		if err != nil {
			// 現状複数でPush通知すると「error3: Mismatched response length. Expected * receipts but only received *」のエラーになるので単発で送信する
			fmt.Println("error3:", err.Error())
		}
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
