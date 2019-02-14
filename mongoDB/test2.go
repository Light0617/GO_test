package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	r := httprouter.New()
	uc := controllers.NewUserController(getSession())
	r.Get("/user/:id", uc.GetUser)
	http.ListenAndServe("localhost:8080", r)
}

func getSession() *mgo.Session {
	s, err := mgo.Dail("mongodb://localhost")
	if err != nil {
		panic(err)
	}
	return s
}
