package main

import (
	"github.com/ant0ine/go-json-rest/rest"
	"go-message-masking/handlers"
	"log"
	"net/http"
)

func main() {
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	api.Use(&rest.AuthBasicMiddleware{
		Realm: "test zone",
		Authenticator: func(userId string, password string) bool {
			if userId == "admin" && password == "admin" {
				return true
			}
			return false
		},
	})
	router, err := rest.MakeRouter(
		rest.Post("/mask", handlers.MaskSensitiveData),
	)

	if err != nil {
		log.Fatal(err)
	}

	api.SetApp(router)

	log.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", api.MakeHandler()))
}
