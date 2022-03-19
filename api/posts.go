package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/robjkc/blog-api/db"
	"github.com/robjkc/blog-api/models"
)

type PostResponse struct {
	ID      int    `json:"id"`
	Author  string `json:"author"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type PostRequest struct {
	Author  string `json:"author"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func GetPosts(w http.ResponseWriter, r *http.Request) {

	posts, err := models.GetPosts(db.DbConn)
	if err != nil {
		log.Println("Cannot get posts", err)
		http.Error(w, "Request failed!", http.StatusNoContent)
	}

	response := []PostResponse{}

	for _, post := range posts {
		response = append(response, PostResponse{ID: post.ID, Author: post.Author, Title: post.Title, Content: post.Content})
	}

	json, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Request failed!", http.StatusNoContent)
	}
	w.Write(json)
}

func AddPost(w http.ResponseWriter, r *http.Request) {
	var post PostRequest

	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = models.AddPost(db.DbConn, post.Author, post.Title, post.Content)
	if err != nil {
		log.Println("Cannot add post", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	postId, err := strconv.Atoi(params["postId"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = models.DeletePost(db.DbConn, postId)
	if err != nil {
		log.Println("Cannot delete post", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
