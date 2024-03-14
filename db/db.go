package db

const (
	DBURI       = "mongodb://localhost:27017"
	DBNAME      = "hr3"
	TEST_DBNAME = "hr3-test"
)

type Store struct {
	User  UserStore
	Hotel HotelStore
	Room  RoomStore
}

// func ToObjectId(id string) (primitive.ObjectID, err) {
// 	oid, err := primitive.ObjectIDFromHex(id)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return oid
// }
