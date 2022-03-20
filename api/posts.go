package api

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
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

type AddPostRequest struct {
	Author  string `json:"author"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type UpdatePostRequest struct {
	Author  string `json:"author"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (r *AddPostRequest) validate() url.Values {
	errors := url.Values{}

	if r.Author == "" {
		errors.Add("author", "The author field is required")
	}

	if r.Title == "" {
		errors.Add("title", "The title field is required")
	}

	if r.Content == "" {
		errors.Add("content", "The content field is required")
	}
	return errors
}

func (r *UpdatePostRequest) validate() url.Values {
	errors := url.Values{}

	if r.Author == "" {
		errors.Add("author", "The author field is required")
	}

	if r.Title == "" {
		errors.Add("title", "The title field is required")
	}

	if r.Content == "" {
		errors.Add("content", "The content field is required")
	}
	return errors
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
	var request AddPostRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if errors := request.validate(); len(errors) > 0 {
		onValidationError(errors, w)
		return
	}

	err = models.AddPost(db.DbConn, request.Author, request.Title, request.Content)
	if err != nil {
		log.Println("Cannot add post", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func onValidationError(errors url.Values, w http.ResponseWriter) {
	err := map[string]interface{}{"validationError": errors}
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(err)
}

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	postId, err := strconv.Atoi(params["postId"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var request UpdatePostRequest

	err = json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if errors := request.validate(); len(errors) > 0 {
		onValidationError(errors, w)
		return
	}

	err = models.UpdatePost(db.DbConn, postId, request.Author, request.Title, request.Content)
	if err != nil {
		log.Println("Cannot update post", err)
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
