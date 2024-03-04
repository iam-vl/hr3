package bcp

import (
	"context"
	"log"

	"github.com/iam-vl/hr3/db"
	"github.com/iam-vl/hr3/types"
	"go.mongodb.org/mongo-driver/mongo"
)

func InsertData(client *mongo.Client) {
	ctx := context.Background()
	coll := client.Database(db.DBNAME).Collection("users")

	user := types.User{
		FirstName: "Vassily",
		LastName:  "La",
	}

	_, err := coll.InsertOne(ctx, user)
	if err != nil {
		log.Fatal(err)
	}

	// var vl types.User
	// err = coll.FindOne(ctx, bson.M{}).Decode(&vl)
	// logf(err)

	// line()
	// fmt.Println(vl)
}
