package main

import (
	"context"
	"fmt"
	"log"

	"github.com/mongodb/mongo-go-driver/mongo"
	// "gopkg.in/mgo.v2/bson"
	"github.com/mongodb/mongo-go-driver/bson"
)

type Book struct {
	// _id   string
	Isbn  string
	title string
}

// var DB *mgo.Database
// var Books *mgo.Collection
var collection *mongo.Collection

func main() {
	connect()
	findAll()
	// books, err := allBook(Books)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(books)
}

func connect() {
	client, err := mongo.Connect(context.TODO(), "mongodb://localhost:27017")

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	collection = client.Database("bookstore").Collection("books")
	// mdb, err := mgo.Dial("mongodb://localhost/bookstore")
	// if err != nil {
	// 	panic(err)
	// }
	// if err = mdb.Ping(); err != nil {
	// 	panic(err)
	// }
	// DB = mdb.DB("bookstore")
	// Books = DB.C("books")

	// fmt.Println("You connected to your mongo database.")
}

func findAll() {
	bks := []Book{}
	// _, err = col.InsertOne(ctx, bson.M{"_id": "111", "name": "ddd", "age": 50})
	res := collection.FindOne(context.TODO(), bson.D{{}})
	err := res.Decode(&bks)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Found a single document: %+v\n", bks)
}

// func allBook(Books *mgo.Collection) ([]Book, error) {
// 	bks := []Book{}
// 	// err := Books.Find(bson.M{}).All(&bks)
// 	err = Books.FindOne().Decode(&bks)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Printf("Found a single document: %+v\n", result)

// 	if err != nil {
// 		return nil, err
// 	}
// 	return bks, nil
// }
