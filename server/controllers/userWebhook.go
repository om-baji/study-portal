package controllers

import (
	"net/http"
	"os"

	"github.com/om-baji/utils"
	svix "github.com/svix/svix-webhooks/go"
)

func postUser(w http.ResponseWriter, r *http.Response) {

	secret := os.Getenv("WEBHOOK_SECRET")

	svix_id := r.Header.Get("svix-id")
	svix_timestamp := r.Header.Get("svix-timestamp")
	svix_signature := r.Header.Get("svix-signature")

	headers := http.Header{}

	headers.Set("svix-id", svix_id)
	headers.Set("svix-timestamp", svix_timestamp)
	headers.Set("svix-signature", svix_signature)

	wh, err := svix.NewWebhook(secret)
	if err != nil {
		utils.ToJSON(w, 403, utils.Response{
			Message: "Webhook validation failed!",
			Code:    403,
		})
	}

	wh.Verify([]byte(secret), headers)

}
