package main

import (
	"bookstore_oauth-api/src/app"
	mongo_con "bookstore_oauth-api/src/datasources/mongo"
	"context"
	"fmt"
)

func main() {
	app.StartApplication()
	defer func() {
		fmt.Println("I'm closing the connection.")
		if err := mongo_con.GetClient().Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
}
