package handlers

import (
	"go-message-masking/persistence"
	"net/http"
	"regexp"

	"github.com/ant0ine/go-json-rest/rest"
)

// Message is a code representation of the data sent by the API user through the wire
type Message struct {
	Locale     string
	Text       string
	MaskSymbol string
}

// MaskSensitiveData is the route handler that responds whenever the `/mask` route
// has been called with valid data
func MaskSensitiveData(w rest.ResponseWriter, r *rest.Request) {

	message := Message{}
	err := r.DecodeJsonPayload(&message)
	var maskSymbol = "X"

	if message.MaskSymbol != "" {
		maskSymbol = message.MaskSymbol
	}

	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	processedMessage := maskSensitiveData(message.Text, persistence.Expressions, maskSymbol)
	w.WriteJson(
		&Message{
			Locale:     message.Locale,
			Text:       processedMessage,
			MaskSymbol: maskSymbol,
		},
	)
}

func maskSensitiveData(s string, expressionMap map[string]string, maskSymbol string) string {
	for _, value := range expressionMap {
		s = applyExpression(s, value, maskSymbol)
	}

	return s
}

func applyExpression(s string, expression string, maskSymbol string) string {
	re := regexp.MustCompile(expression)
	return re.ReplaceAllStringFunc(s, func(str string) string {
		var maskedValue string
		if len(maskSymbol) > 1 {
			return maskSymbol
		}

		for range str {
			maskedValue += maskSymbol
		}
		return maskedValue
	})
}
