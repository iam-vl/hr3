package api

import (
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/iam-vl/hr3/db"
	"github.com/iam-vl/hr3/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserHandler struct {
	userStore db.UserStore
}

func NewUserHandler(userStore db.UserStore) *UserHandler {
	return &UserHandler{
		userStore: userStore,
	}
}

func (h *UserHandler) HandleDeleteUser(c *fiber.Ctx) error {
	userId := c.Params("id")
	err := h.userStore.DeleteUser(c.Context(), userId)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return c.JSON(map[string]string{"error": "Cannot delete a user with an incorrect id"})
		}
	}
	return c.JSON(map[string]string{"Deleted user": userId})
}

func (h *UserHandler) HandlePutUser(c *fiber.Ctx) error {
	fmt.Println("Inside HandlePutUser()")
	var (
		// values bson.M
		params types.UpdateUserParams
		userId = c.Params("id")
	)
	oid, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return err
	}

	// if err = c.BodyParser(&values); err != nil {
	if err = c.BodyParser(&params); err != nil {
		return err
	}
	filter := bson.M{"_id": oid}
	fmt.Println("Inside HandlePutUser() - updating user")
	if err = h.userStore.UpdateUser(c.Context(), filter, params); err != nil {
		return err
	}
	// fmt.Printf("Context: %+v\n", c.Context())
	// fmt.Printf("Filter: %+v\n", filter)
	// fmt.Printf("Update: %+v\n", update)
	// fmt.Printf("User ID: %v\n", userId)
	// At around 25:00, you get an error like that: {"error": "update document must contain key beginning with '$"}
	return c.JSON(map[string]string{"Updated user": userId})
}

func (h *UserHandler) HandlePostUser(c *fiber.Ctx) error {
	var params types.UserParams
	if err := c.BodyParser(&params); err != nil {
		return err
	}
	errors := params.Validate()
	if len(errors) > 0 {
		return c.JSON(errors)
	}
	user, err := types.NewUserFromParams(params)
	if err != nil {
		return err
	}
	insertedUser, err := h.userStore.InsertUser(c.Context(), user)
	if err != nil {
		return err
	}

	return c.JSON(insertedUser)
}

func (h *UserHandler) HandleGetUser(c *fiber.Ctx) error {
	var (
		id = c.Params("id")
	)

	user, err := h.userStore.GetUserById(c.Context(), id)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return c.JSON(map[string]string{"msg": "Cannot find a user with this ID"})
		}
	}
	return c.JSON(user)
}

func (h *UserHandler) HandleGetUsers(c *fiber.Ctx) error {
	fmt.Println("Inside HandleGetUsers")
	users, err := h.userStore.GetUsers(c.Context())
	if err != nil {
		return err
	}
	return c.JSON(users)
}
