package main

import (
	"context"
	"fmt"
	"log"

	"github.com/iam-vl/hr3/db"
	"github.com/iam-vl/hr3/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	ctx := context.Background()
	fmt.Println("Connecting mongo...")
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(db.DBURI))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Creating collections...")
	hotelStore := db.NewMongoHotelStore(client)
	roomStore := db.NewMongoRoomStore(client, hotelStore)

	fmt.Println("Seeding database...")
	hotel := types.Hotel{
		Name:     "Grand dauphin",
		Location: "Lyon, France",
		Rooms:    []primitive.ObjectID{},
	}
	insertedHotel, err := hotelStore.InsertHotel(ctx, &hotel)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Hotel: %+v\n", insertedHotel)

	roomA := types.Room{Type: types.SingleRoomType, BasePrice: 99.99}
	roomB := types.Room{Type: types.DeluxeRoomType, BasePrice: 199.99}
	roomC := types.Room{Type: types.SeaviewRoomType, BasePrice: 139.99}
	rooms := []types.Room{roomA, roomB, roomC}

	for _, r := range rooms {
		r.HotelID = insertedHotel.ID
		insertedRoom, err := roomStore.InsertRoom(ctx, &r)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Room: %+v\n", insertedRoom)
	}
}
