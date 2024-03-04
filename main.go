package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/iam-vl/hr3/api"
	"github.com/iam-vl/hr3/types"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	DBURI    = "mongodb://localhost:27017"
	DBNAME   = "hr3"
	USERCOLL = "users"
)

func main() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(DBURI))
	logf(err)

	ctx := context.Background()
	coll := client.Database(DBNAME).Collection(USERCOLL)

	user := types.User{
		FirstName: "Vassily",
		LastName:  "La",
	}

	res, err := coll.InsertOne(ctx, user)
	logf(err)

	line()
	fmt.Println(res)

	listenAddr := flag.String("listenAddr", ":5000", "The listen address of the API server")
	flag.Parse()

	app := fiber.New()
	apiv1 := app.Group("/api/v1")

	apiv1.Get("/user", api.HandleGetUsers)
	apiv1.Get("/user/:id", api.HandleGetUser)
	app.Listen(*listenAddr)
}
