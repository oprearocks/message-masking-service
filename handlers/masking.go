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
	MaskString string
}

// MaskSensitiveData is the route handler that responds whenever the `/mask` route
// has been called with valid data
func MaskSensitiveData(w rest.ResponseWriter, r *rest.Request) {

	message := Message{}
	err := r.DecodeJsonPayload(&message)
	var maskString = "(hidden)"

	if message.MaskString != "" {
		maskString = message.MaskString
	}

	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	processedMessage := maskSensitiveData(message.Text, persistence.Expressions, maskString)
	w.WriteJson(
		&Message{
			Locale:     message.Locale,
			Text:       processedMessage,
			MaskString: maskString,
		},
	)
}

func maskSensitiveData(s string, expressionMap map[string]string, maskString string) string {
	for _, value := range expressionMap {
		s = applyExpression(s, value, maskString)
	}

	return s
}

func applyExpression(s string, expression string, maskString string) string {
	re := regexp.MustCompile(expression)
	return re.ReplaceAllString(s, maskString)
}
