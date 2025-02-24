package main

import (
	"learnProject/configs"
	"learnProject/internal/auth"
	"learnProject/internal/links"
	"learnProject/pkg/db"
	"net/http"
)

func main() {
	conf := configs.LoadConfig()
	_ = db.NewDb(conf)
	router := http.NewServeMux()

	//Handler
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{Config: conf})
	links.NewLinkHandler(router, links.LinkHandlerDeps{Config: conf})

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	server.ListenAndServe()
}
