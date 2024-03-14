package api

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/iam-vl/hr3/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type HotelHandler struct {
	roomStore  db.RoomStore
	hotelStore db.HotelStore
}

type HotelQueryParams struct {
	Rooms  bool
	Rating int
}

func NewHotelHandler(hs db.HotelStore, rs db.RoomStore) *HotelHandler {
	return &HotelHandler{
		roomStore:  rs,
		hotelStore: hs,
	}
}

func (h *HotelHandler) HandleGetRooms(c *fiber.Ctx) error {
	id := c.Params("id")
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	filter := bson.M{"hotelID": oid}
	rooms, err := h.roomStore.GetRoomsByHotelId(c.Context(), filter)
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

	hotels, err := h.hotelStore.GetHotels(c.Context(), nil)
	if err != nil {
		return err
	}

	fmt.Println("Hotels:", hotels)
	return c.JSON(hotels)

}
