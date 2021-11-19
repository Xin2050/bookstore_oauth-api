package app

import (
	mongo_con "bookstore_oauth-api/src/datasources/mongo"
	"bookstore_oauth-api/src/domain/access_token"
	"bookstore_oauth-api/src/http"
	"bookstore_oauth-api/src/repository/db"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	dbClient := mongo_con.GetClient()
	if dbClient == nil {
		panic("dbError, please check the MongoDB")
	}
	atHandler := http.NewHandler(access_token.NewService(db.NewRepository()))
	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)
	router.Run(":3001")

}
