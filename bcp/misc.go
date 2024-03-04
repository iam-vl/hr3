package bcp

// func deleteData() {
// 	ctx := context.Background()
// 	coll := client.Database(DBNAME).Collection(USERCOLL)

// 	user := types.User{
// 		FirstName: "Vassily",
// 		LastName:  "La",
// 	}

// 	_, err = coll.InsertOne(ctx, user)
// 	logf(err)

// 	var vl types.User
// 	err = coll.FindOne(ctx, bson.M{}).Decode(&vl)
// 	logf(err)

// 	line()
// 	fmt.Println(vl)
// }
