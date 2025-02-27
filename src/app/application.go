package app

import (
	"github.com/dralos/bookstore_oauth-api/clients/cassandra"
	"github.com/dralos/bookstore_oauth-api/src/domain/access_token"
	"github.com/dralos/bookstore_oauth-api/src/http"
	"github.com/dralos/bookstore_oauth-api/src/repository/db"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApp() {

	session, err := cassandra.GetSession()
	if err != nil {
		panic(err)
	}
	session.Close()

	atHandler := http.NewHandler(access_token.NewService(db.NewRepository()))

	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)
	router.POST("/oauth/access_token", atHandler.Create)
	router.PUT("/oauth/access_token", atHandler.Update)

	router.Run(":8080")
}
