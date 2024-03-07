package api

import (
	"context"
	"log"
	"testing"

	"github.com/iam-vl/hr3/db"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const testDbUri = "mongodb://localhost:27017"

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
		UserStore: db.NewMongoUserStore(client),
	}

}

func TestPostUser(t *testing.T) {
	tdb := setup(t)
	defer tdb.tearDown(t)
}
