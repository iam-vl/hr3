package db

const (
	DBURI  = "mongodb://localhost:27017"
	DBNAME = "hr3"
)

// func ToObjectId(id string) (primitive.ObjectID, err) {
// 	oid, err := primitive.ObjectIDFromHex(id)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return oid
// }
