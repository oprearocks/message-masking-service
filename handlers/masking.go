package handlers

import (
	"github.com/ant0ine/go-json-rest/rest"
	"go-message-masking/persistence"
	"net/http"
	"regexp"
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

	processed_message := mask_sensitive_data(message.Text, persistence.Expressions)
	w.WriteJson(
		&Message{
			Locale: message.Locale,
			Text:   processed_message,
		},
	)
}

func mask_sensitive_data(s string, expressions_map map[string]string) string {
	for _, value := range expressions_map {
		s = apply_expression(s, value)
	}

	return s
}

func apply_expression(s string, expression string) string {
	re := regexp.MustCompile(expression)
	return re.ReplaceAllStringFunc(s, func(str string) string {
		var mask string
		for range str {
			mask += "X"
		}
		return mask
	})
}
