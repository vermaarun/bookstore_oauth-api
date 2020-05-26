package app

import (
	"github.com/gin-gonic/gin"
	"github.com/vermaarun/bookstore_oauth-api/src/domain/access_token"
	"github.com/vermaarun/bookstore_oauth-api/src/http"
	"github.com/vermaarun/bookstore_oauth-api/src/repository/db"
)

var (
	router = gin.Default()
)

// Application  => handler => service => repository

func StartApplication() {
	dbRepository := db.NewRepository()
	atService := access_token.NewService(dbRepository)
	atHandler := http.NewHandler(atService)

	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)
	router.POST("/oauth/access_token", atHandler.Create)

	router.Run(":8080")
}
