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

	router := mux.NewRouter()
	router.Use(commonMiddleware)

	router.HandleFunc("/", handler)
	router.HandleFunc("/posts", api.GetPosts).Methods(http.MethodGet)
	router.HandleFunc("/posts", api.AddPost).Methods(http.MethodPost)
	router.HandleFunc("/posts/{postId}", api.UpdatePost).Methods(http.MethodPut)
	router.HandleFunc("/posts/{postId}", api.DeletePost).Methods(http.MethodDelete)
	router.HandleFunc("/posts/{postId}", api.GetPost).Methods(http.MethodGet)
	router.HandleFunc("/posts/{postId}/comments", api.AddComment).Methods(http.MethodPost)

	log.Println("Listening for connections on port: 8080")

	http.ListenAndServe(":8080", router)

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

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
