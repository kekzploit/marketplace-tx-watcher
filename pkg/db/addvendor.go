package db

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func AddVendor(uri string, image string, title string, description string, secret string, storeType string, hash string) bool {
	// instantiate mongodb client
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	collection := client.Database("marketplace").Collection("vendors")

	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := collection.InsertOne(ctx, bson.D{
		{Key: "image", Value: image},
		{Key: "title", Value: title},
		{Key: "description", Value: description},
		{Key: "socials", Value: bson.D{
			{Key: "twitter", Value: ""},
			{Key: "discord", Value: ""},
			{Key: "telegram", Value: ""},
			{Key: "discord", Value: ""},
		}},
		{Key: "secret", Value: secret},
		{Key: "type", Value: storeType},
		{Key: "hash", Value: hash}})

	id := res.InsertedID

	if id != "" {
		return true
	} else {
		return false
	}
}
