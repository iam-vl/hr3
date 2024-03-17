package api

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/iam-vl/hr3/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type HotelHandler struct {
	store *db.Store
	// roomStore  db.RoomStore
	// hotelStore db.HotelStore
}

// type HotelQueryParams struct {
// 	Rooms  bool
// 	Rating int
// }

func NewHotelHandler(store *db.Store) *HotelHandler {
	return &HotelHandler{
		store: store,
	}
}

func (h *HotelHandler) HandleGetRooms(c *fiber.Ctx) error {
	id := c.Params("id")
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	filter := bson.M{"hotelID": oid}
	rooms, err := h.store.Room.GetRoomsByHotelId(c.Context(), filter)
	// rooms, err := h.Store.room.GetRoomsByHotelId(c.Context(), filter)
	if err != nil {
		return err
	}
	return c.JSON(rooms)
}

func (h *HotelHandler) HandleGetHotels(c *fiber.Ctx) error {
	fmt.Println("Running hotel handler..")
	// var qparams HotelQueryParams
	// if err := c.QueryParser(&qparams); err != nil {
	// 	return err
	// }

	hotels, err := h.store.Hotel.GetHotels(c.Context(), nil)
	if err != nil {
		return err
	}

	fmt.Println("Hotels:", hotels)
	return c.JSON(hotels)

}

func (h *HotelHandler) HandleGetHotel(c *fiber.Ctx) error {
	fmt.Println("Getting one hotel handler..")
	id := c.Params("id")
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	hotel, err := h.store.Hotel.GetHotelById(c.Context(), oid)
	if err != nil {
		return err
	}

	fmt.Println("Found a hotel:", hotel)
	return c.JSON(hotel)

}
