package db

import (
	"context"

	"github.com/iam-vl/hr3/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	USERCOLL = "users"
)

type UserStore interface {
	GetUserById(context.Context, string) (*types.User, error)
}

type MongoUserStore struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func NewMongoUserStore(cl *mongo.Client) *MongoUserStore {
	return &MongoUserStore{
		client: cl,
		coll:   cl.Database(DBNAME).Collection(USERCOLL),
	}
}

func (s MongoUserStore) GetUserById(ctx context.Context, id string) (*types.User, error) {
	var user types.User
	err := s.coll.FindOne(ctx, bson.M{"_id": ToObjectId(id)}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
