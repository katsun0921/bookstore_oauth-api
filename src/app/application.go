package app

import (
  "github.com/gin-gonic/gin"
  "github.com/katsun0921/bookstore_oauth-api/src/domain/access_token"
  "github.com/katsun0921/bookstore_oauth-api/src/http"
  "github.com/katsun0921/bookstore_oauth-api/src/repository/db"
)

var (
  router = gin.Default()
)

func StartApplication() {
  atHandler := http.NewAccessTokenHandler(access_token.NewService(db.NewRepository()))

  router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)
  router.Run(":8080")
}
