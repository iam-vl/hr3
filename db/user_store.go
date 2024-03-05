package db

import (
	"context"
	"fmt"

	"github.com/iam-vl/hr3/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	USERCOLL = "users"
)

type UserStore interface {
	GetUserById(context.Context, string) (*types.User, error)
	GetUsers(context.Context) ([]*types.User, error)
	InsertUser(context.Context, *types.User) (*types.User, error)
	DeleteUser(context.Context, string) error
	UpdateUser(ctx context.Context, filter bson.M, update bson.M) error
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

func (s *MongoUserStore) DeleteUser(ctx context.Context, userId string) error {
	oid, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return err
	}
	_, err = s.coll.DeleteOne(ctx, bson.M{"_id": oid})
	if err != nil {
		return err
	}
	// if res.DeletedCount == 0 {}
	return nil
}

func (s *MongoUserStore) UpdateUser(ctx context.Context, filter bson.M, update bson.M) error {
	fmt.Println("Inside UpdateUser()")
	_, err := s.coll.UpdateOne(ctx, filter, update)
	fmt.Printf("Ctx: %+v\n", ctx)
	if err != nil {
		return err
	}
	fmt.Println("Quitting UpdateUser()")

	return nil
}

func (s *MongoUserStore) InsertUser(ctx context.Context, user *types.User) (*types.User, error) {
	res, err := s.coll.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	user.ID = res.InsertedID.(primitive.ObjectID)
	return user, nil
}

func (s MongoUserStore) GetUsers(ctx context.Context) ([]*types.User, error) {
	// fmt.Println("inside get users")
	cursor, err := s.coll.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	var users []*types.User
	if err := cursor.All(ctx, &users); err != nil {
		fmt.Println("error getting cursor")
		return []*types.User{}, err
	}

	return users, nil
}

func (s MongoUserStore) GetUserById(ctx context.Context, id string) (*types.User, error) {
	// validate the id
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var user types.User
	err = s.coll.FindOne(ctx, bson.M{"_id": oid}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
