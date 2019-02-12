package books

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Book struct {
	Isbn  string
	title string
}

func allBook(Books mgo.Collection) ([]Book, error) {
	bks := []Book{}
	err := Books.Find(bson.M{}).All(&bks)
	if err != nil {
		return nil, err
	}
	return bks, nil
}
