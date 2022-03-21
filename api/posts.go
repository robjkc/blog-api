package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/robjkc/blog-api/db"
	"github.com/robjkc/blog-api/models"
)

type GetPostsResponse struct {
	ID         int       `json:"id"`
	Author     string    `json:"author"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	CreateDate time.Time `json:"create_date"`
}

type GetPostResponse struct {
	ID         int               `json:"id"`
	Author     string            `json:"author"`
	Title      string            `json:"title"`
	Content    string            `json:"content"`
	CreateDate time.Time         `json:"create_date"`
	Comments   []CommentResponse `json:"comments"`
}

type CommentResponse struct {
	ID              int       `json:"id"`
	PostID          int       `json:"post_id"`
	TopLevel        int       `json:"top_level"`
	ParentCommentID int       `json:"parent_comment_id"`
	Author          string    `json:"author"`
	Content         string    `json:"content"`
	CreateDate      time.Time `json:"create_date"`
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

func GetPosts(w http.ResponseWriter, r *http.Request) {

	posts, err := models.GetPosts(db.DbConn)
	if err != nil {
		log.Println("Cannot get posts", err)
		http.Error(w, "Request failed!", http.StatusNoContent)
	}

	responses := []GetPostsResponse{}

	for _, post := range posts {
		responses = append(responses, GetPostsResponse{ID: post.ID, Author: post.Author, Title: post.Title,
			Content: post.Content, CreateDate: post.CreateDate})
	}

	json, err := json.Marshal(responses)
	if err != nil {
		http.Error(w, "Request failed!", http.StatusNoContent)
	}
	w.Write(json)
}

func GetPost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	postId, err := strconv.Atoi(params["postId"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	post, err := models.GetPost(db.DbConn, postId)
	if err != nil {
		log.Println("Cannot get post", err)
		http.Error(w, "Request failed!", http.StatusNoContent)
	}

	comments, err := models.GetComments(db.DbConn, postId)
	if err != nil {
		log.Println("Cannot get comments", err)
		http.Error(w, "Request failed!", http.StatusNoContent)
	}

	commentResponses := []CommentResponse{}

	for _, comment := range comments {
		commentResponses = append(commentResponses, CommentResponse{ID: comment.ID, PostID: comment.PostID,
			TopLevel: comment.TopLevel, ParentCommentID: comment.ParentCommentID, Author: comment.Author,
			Content: comment.Content, CreateDate: comment.CreateDate,
		})
	}

	response := GetPostResponse{ID: post.ID, Author: post.Author, Title: post.Title, Content: post.Content,
		CreateDate: post.CreateDate}
	response.Comments = commentResponses

	json, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Request failed!", http.StatusNoContent)
	}
	w.Write(json)
}

func AddPostWithAuth(w http.ResponseWriter, r *http.Request) {
	authToken := r.Header.Get("Authorization")
	authArr := strings.Split(authToken, " ")

	if len(authArr) != 2 {
		log.Println("Authentication header is invalid: " + authToken)
		http.Error(w, "Request failed!", http.StatusUnauthorized)
		return
	}

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
