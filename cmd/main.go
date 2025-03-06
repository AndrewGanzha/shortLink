package main

import (
	"learnProject/configs"
	"learnProject/internal/auth"
	"learnProject/internal/links"
	"learnProject/internal/user"
	"learnProject/pkg/db"
	"learnProject/pkg/middleware"
	"net/http"
)

func main() {
	conf := configs.LoadConfig()
	db := db.NewDb(conf)
	router := http.NewServeMux()

	//Repositories
	linksRepository := links.NewLinkRepository(db)
	userRepository := user.NewUserRepository(db)

	//Handler
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{Config: conf})
	links.NewLinkHandler(router, links.LinkHandlerDeps{LinkRepository: linksRepository})
	user.NewUserHandler(router, user.UserHandlerDeps{UserRepository: userRepository})

	//Middlewares

	stack := middleware.Chain(
		middleware.CORS,
		middleware.Logging,
	)

	server := http.Server{
		Addr:    ":8080",
		Handler: stack(router),
	}
	server.ListenAndServe()
}
