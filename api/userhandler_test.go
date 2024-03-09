package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/iam-vl/hr3/db"
	"github.com/iam-vl/hr3/types"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	testDbUri = "mongodb://localhost:27017"
	dbName    = "hr3-test"
)

type testDb struct {
	db.UserStore
}

func (tdb *testDb) tearDown(t *testing.T) {
	if err := tdb.UserStore.Drop(context.TODO()); err != nil {
		t.Fatal(err)
	}

}

func setup(t *testing.T) *testDb {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(testDbUri))
	if err != nil {
		log.Fatal(err)
	}
	return &testDb{
		UserStore: db.NewMongoUserStore(client, dbName),
	}

}

func TestPostUser(t *testing.T) {
	tdb := setup(t)
	defer tdb.tearDown(t)

	app := fiber.New()
	uh := NewUserHandler(tdb.UserStore)
	app.Post("/", uh.HandlePostUser)
	params := types.UserParams{
		Email:     "gztrk@eng.com",
		FirstName: "Vas",
		LastName:  "Lap",
		Password:  "hfgdhjgfhjvfc",
	}
	b, _ := json.Marshal(params)
	req := httptest.NewRequest("POST", "/", bytes.NewReader(b))
	req.Header.Add("Content-Type", "application/json")
	resp, err := app.Test(req)
	if err != nil {
		t.Error(err)
	}
	var user types.User
	json.NewDecoder(resp.Body).Decode(&user)
	fmt.Println(user)
	if len(user.ID) == 0 {
		t.Errorf("Cannot confirm setting user ID")
	}
	if len(user.EncryptedPassword) > 0 {
		t.Errorf("Encrypted password must not be included in a JSON response")
	}
	if user.FirstName != params.FirstName {
		t.Errorf("Expected first name %s, but got %s", params.FirstName, user.FirstName)
	}
	if user.LastName != params.LastName {
		t.Errorf("Expected last name %s, but got %s", params.LastName, user.LastName)
	}
	if user.Email != params.Email {
		t.Errorf("Expected email %s, but got %s", params.Email, user.Email)
	}
	// bb, err := io.ReadAll(resp.Body)
	// fmt.Println(string(bb))

}
