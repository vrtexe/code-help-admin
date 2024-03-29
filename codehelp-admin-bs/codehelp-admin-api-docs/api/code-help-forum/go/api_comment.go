/*
 * Forum application
 *
 * Forum application api
 *
 * API version: 0.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package code_help_forum_api

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// CommentAPIController binds http requests to an api service and writes the service results to the http response
type CommentAPIController struct {
	service CommentAPIServicer
	errorHandler ErrorHandler
}

// CommentAPIOption for how the controller is set up.
type CommentAPIOption func(*CommentAPIController)

// WithCommentAPIErrorHandler inject ErrorHandler into controller
func WithCommentAPIErrorHandler(h ErrorHandler) CommentAPIOption {
	return func(c *CommentAPIController) {
		c.errorHandler = h
	}
}

// NewCommentAPIController creates a default api controller
func NewCommentAPIController(s CommentAPIServicer, opts ...CommentAPIOption) Router {
	controller := &CommentAPIController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the CommentAPIController
func (c *CommentAPIController) Routes() Routes {
	return Routes{
		"CommentOnPost": Route{
			strings.ToUpper("Post"),
			"/api/v1/comment",
			c.CommentOnPost,
		},
		"DeleteComment": Route{
			strings.ToUpper("Delete"),
			"/api/v1/comment/{uid}",
			c.DeleteComment,
		},
		"GetCommentReplies": Route{
			strings.ToUpper("Get"),
			"/api/v1/comment/{uid}",
			c.GetCommentReplies,
		},
		"GetCommentsForPost": Route{
			strings.ToUpper("Get"),
			"/api/v1/comment",
			c.GetCommentsForPost,
		},
		"ReplyToComment": Route{
			strings.ToUpper("Post"),
			"/api/v1/comment/{uid}",
			c.ReplyToComment,
		},
		"UpdateComment": Route{
			strings.ToUpper("Put"),
			"/api/v1/comment/{uid}",
			c.UpdateComment,
		},
	}
}

// CommentOnPost - 
func (c *CommentAPIController) CommentOnPost(w http.ResponseWriter, r *http.Request) {
	query, err := parseQuery(r.URL.RawQuery)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	var postParam string
	if query.Has("post") {
		param := query.Get("post")

		postParam = param
	} else {
		c.errorHandler(w, r, &RequiredError{Field: "post"}, nil)
		return
	}
	commentRequestParam := CommentRequest{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&commentRequestParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertCommentRequestRequired(commentRequestParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertCommentRequestConstraints(commentRequestParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.CommentOnPost(r.Context(), postParam, commentRequestParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// DeleteComment - 
func (c *CommentAPIController) DeleteComment(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	uidParam := params["uid"]
	if uidParam == "" {
		c.errorHandler(w, r, &RequiredError{"uid"}, nil)
		return
	}
	result, err := c.service.DeleteComment(r.Context(), uidParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// GetCommentReplies - 
func (c *CommentAPIController) GetCommentReplies(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	uidParam := params["uid"]
	if uidParam == "" {
		c.errorHandler(w, r, &RequiredError{"uid"}, nil)
		return
	}
	result, err := c.service.GetCommentReplies(r.Context(), uidParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// GetCommentsForPost - 
func (c *CommentAPIController) GetCommentsForPost(w http.ResponseWriter, r *http.Request) {
	query, err := parseQuery(r.URL.RawQuery)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	var postParam string
	if query.Has("post") {
		param := query.Get("post")

		postParam = param
	} else {
		c.errorHandler(w, r, &RequiredError{Field: "post"}, nil)
		return
	}
	result, err := c.service.GetCommentsForPost(r.Context(), postParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// ReplyToComment - 
func (c *CommentAPIController) ReplyToComment(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	uidParam := params["uid"]
	if uidParam == "" {
		c.errorHandler(w, r, &RequiredError{"uid"}, nil)
		return
	}
	commentRequestParam := CommentRequest{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&commentRequestParam); err != nil && !errors.Is(err, io.EOF) {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertCommentRequestRequired(commentRequestParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertCommentRequestConstraints(commentRequestParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.ReplyToComment(r.Context(), uidParam, commentRequestParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// UpdateComment - 
func (c *CommentAPIController) UpdateComment(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	uidParam := params["uid"]
	if uidParam == "" {
		c.errorHandler(w, r, &RequiredError{"uid"}, nil)
		return
	}
	commentRequestParam := CommentRequest{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&commentRequestParam); err != nil && !errors.Is(err, io.EOF) {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertCommentRequestRequired(commentRequestParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertCommentRequestConstraints(commentRequestParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.UpdateComment(r.Context(), uidParam, commentRequestParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}
