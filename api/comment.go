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

type AddCommentRequest struct {
	ParentCommentID int    `json:"parent_comment_id"`
	Author          string `json:"author"`
	Content         string `json:"content"`
}

func AddComment(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	postId, err := strconv.Atoi(params["postId"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var request AddCommentRequest

	err = json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if errors := request.validate(); len(errors) > 0 {
		onValidationError(errors, w)
		return
	}

	err = models.AddComment(db.DbConn, postId, request.ParentCommentID, request.Author, request.Content)
	if err != nil {
		log.Println("Cannot add comment", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
