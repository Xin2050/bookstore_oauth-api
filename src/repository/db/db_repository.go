package db

import (
	mongo_con "bookstore_oauth-api/src/datasources/mongo"
	"bookstore_oauth-api/src/domain/access_token"
	"bookstore_oauth-api/src/utils/errors"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type DbRepository interface {
	GetById(string) (*access_token.AccessToken, *errors.RestError)
}

type dbRepository struct {
}

func NewRepository() DbRepository {
	return &dbRepository{}
}

func (r *dbRepository) GetById(string) (*access_token.AccessToken, *errors.RestError) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	collection := mongo_con.GetDatabase().Collection("ratings")

	res, err := collection.InsertOne(ctx, bson.D{{"name", "pi"}, {"value", 3.14159}})
	if err != nil {
		panic(err)
	}
	id := res.InsertedID.(primitive.ObjectID)
	token := access_token.GetNewAccessToken()
	token.AccessToken = id.String()
	return token, nil
}
