package main

import (
	"context"
	"flag"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/iam-vl/hr3/api"
	"github.com/iam-vl/hr3/db"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var config = fiber.Config{
	ErrorHandler: func(c *fiber.Ctx, err error) error {
		return c.JSON(map[string]string{"error": err.Error()})
	},
}

func main() {

	listenAddr := flag.String("listenAddr", ":1111", "The listen address of the API server")
	flag.Parse()

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(db.DBURI))
	if err != nil {
		log.Fatal(err)
	}

	// bcp.InsertData(client)

	// Data stores and handlers
	var (
		hotelStore = db.NewMongoHotelStore(client)
		roomStore  = db.NewMongoRoomStore(client, hotelStore)
		userStore  = db.NewMongoUserStore(client)
		store      = &db.Store{
			Hotel: hotelStore,
			Room:  roomStore,
			User:  userStore,
		}

		// uh = api.NewUserHandler(db.NewMongoUserStore(client))
		uh = api.NewUserHandler(userStore)
		hh = api.NewHotelHandler(store)

		app   = fiber.New(config) // add config for errors
		apiv1 = app.Group("/api/v1")
	)

	// User handlers
	apiv1.Get("/user", uh.HandleGetUsers)
	apiv1.Post("/user", uh.HandlePostUser)
	apiv1.Get("/user/:id", uh.HandleGetUser)
	apiv1.Delete("/user/:id", uh.HandleDeleteUser)
	apiv1.Put("/user/:id", uh.HandlePutUser) // update

	// Hotel handlers
	apiv1.Get("/hotels", hh.HandleGetHotels)
	apiv1.Get("/hotels/:id", hh.HandleGetHotel)
	apiv1.Get("/hotel/:id/rooms", hh.HandleGetRooms)

	app.Listen(*listenAddr)
}
