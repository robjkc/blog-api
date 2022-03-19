package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/robjkc/blog-api/api"
	"github.com/robjkc/blog-api/db"
)

func main() {
	err := db.ConnectDb()
	if err != nil {
		log.Fatal(err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/", handler)
	r.HandleFunc("/posts", api.GetPosts).Methods(http.MethodGet)
	r.HandleFunc("/posts", api.AddPost).Methods(http.MethodPost)
	r.HandleFunc("/post/{postId}", api.UpdatePost).Methods(http.MethodPut)
	r.HandleFunc("/post/{postId}", api.DeletePost).Methods(http.MethodDelete)

	log.Println("Listening for connections on port: 8080")

	http.ListenAndServe(":8080", r)

	/*
		r.HandleFunc("/users/", listUsers).Methods(http.MethodGet)
		r.HandleFunc("/users/", createUser).Methods(http.MethodPost)
		r.HandleFunc("/users/{userId}/", getUser).Methods(http.MethodGet)
		r.HandleFunc("/users/{userId}/", updateUser).Methods(http.MethodPut)
		r.HandleFunc("/users/{userId}/", deleteUser).Methods(Http.MethodDelete)*/
}

func handler(w http.ResponseWriter, r *http.Request) {
	// mux.Vars(r) returns all values captured in the request URL.
	vars := mux.Vars(r)
	// vars is a dictionary whose key-value pairs are variables' name-value pairs.
	fmt.Fprintf(w, "User %s\n", vars["userId"])
}
