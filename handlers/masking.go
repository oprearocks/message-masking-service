package handlers

import (
	"go-message-masking/persistence"
	"net/http"
	"regexp"

	"github.com/ant0ine/go-json-rest/rest"
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

	processed_message := maskSensitiveData(message.Text, persistence.Expressions)
	w.WriteJson(
		&Message{
			Locale: message.Locale,
			Text:   processed_message,
		},
	)
}

func maskSensitiveData(s string, expressionMap map[string]string) string {
	for _, value := range expressionMap {
		s = applyExpression(s, value)
	}

	return s
}

func applyExpression(s string, expression string) string {
	re := regexp.MustCompile(expression)
	return re.ReplaceAllStringFunc(s, func(str string) string {
		var mask string
		for range str {
			mask += "X"
		}
		return mask
	})
}
