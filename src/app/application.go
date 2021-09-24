package app

import (
  "github.com/gin-gonic/gin"
  "github.com/katsun0921/bookstore_oauth-api/src/http"
  "github.com/katsun0921/bookstore_oauth-api/src/repository/db"
  "github.com/katsun0921/bookstore_oauth-api/src/repository/rest"
  "github.com/katsun0921/bookstore_oauth-api/src/services/access_token"
)

var (
	router = gin.Default()
)

func StartApplication() {
	atHandler := http.NewAccessTokenHandler(access_token.NewService(rest.NewRepository(), db.NewRepository()))

	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)
	router.POST("/oauth/access_token", atHandler.Create)

	err := router.Run(":8080")
	if err != nil {
		return
	}
}
