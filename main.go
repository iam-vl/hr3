package main

import (
	"context"
	"flag"

	"github.com/gofiber/fiber/v2"
	"github.com/iam-vl/hr3/api"
	"github.com/iam-vl/hr3/db"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	DBURI    = "mongodb://localhost:27017"
	DBNAME   = "hr3"
	USERCOLL = "users"
)

var config = fiber.Config{
	ErrorHandler: func(c *fiber.Ctx, err error) error {
		return c.JSON(map[string]string{"error": err.Error()})
	},
}

func main() {

	listenAddr := flag.String("listenAddr", ":1111", "The listen address of the API server")
	flag.Parse()

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(DBURI))
	logf(err)

	// bcp.InsertData(client)

	uh := api.NewUserHandler(db.NewMongoUserStore(client)) // user handler

	app := fiber.New(config) // add config for errors
	apiv1 := app.Group("/api/v1")

	apiv1.Get("/user", uh.HandleGetUsers)
	apiv1.Post("/user", uh.HandlePostUser)
	apiv1.Get("/user/:id", uh.HandleGetUser)
	apiv1.Delete("/user/:id", uh.HandleDeleteUser)
	apiv1.Put("/user/:id", uh.HandlePutUser) // update

	app.Listen(*listenAddr)
}
