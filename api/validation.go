package api

import (
	"encoding/json"
	"net/http"
	"net/url"
)

func onValidationError(errors url.Values, w http.ResponseWriter) {
	err := map[string]interface{}{"validationError": errors}
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(err)
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

func (r *AddCommentRequest) validate() url.Values {
	errors := url.Values{}

	if r.Author == "" {
		errors.Add("author", "The author field is required")
	}

	if r.Content == "" {
		errors.Add("content", "The content field is required")
	}
	return errors
}
