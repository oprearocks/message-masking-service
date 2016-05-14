package handlers

import (
	"github.com/ant0ine/go-json-rest/rest"
	"net/http"
	"strings"
)

type Message struct {
	Locale string
	Text   string
}

func MaskSensitiveData(w rest.ResponseWriter, r *rest.Request) {
	message := Message{}
	err := r.DecodeJsonPayload(&message)

	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	processed_message := mask_sensitive_data(message.Text)
	w.WriteJson(
		&Message{
			Locale: message.Locale,
			Text:   processed_message,
		},
	)
}

func mask_sensitive_data(s string) string {
	return strings.ToUpper(s)
}
