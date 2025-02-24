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
	db := db.NewDb(conf)
	router := http.NewServeMux()

	//Repositories
	linksRepository := links.NewLinkRepository(db)

	//Handler
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{Config: conf})
	links.NewLinkHandler(router, links.LinkHandlerDeps{LinkRepository: linksRepository})

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	server.ListenAndServe()
}
