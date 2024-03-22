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

// PostAPIController binds http requests to an api service and writes the service results to the http response
type PostAPIController struct {
	service PostAPIServicer
	errorHandler ErrorHandler
}

// PostAPIOption for how the controller is set up.
type PostAPIOption func(*PostAPIController)

// WithPostAPIErrorHandler inject ErrorHandler into controller
func WithPostAPIErrorHandler(h ErrorHandler) PostAPIOption {
	return func(c *PostAPIController) {
		c.errorHandler = h
	}
}

// NewPostAPIController creates a default api controller
func NewPostAPIController(s PostAPIServicer, opts ...PostAPIOption) Router {
	controller := &PostAPIController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the PostAPIController
func (c *PostAPIController) Routes() Routes {
	return Routes{
		"CreateCommunityPost": Route{
			strings.ToUpper("Post"),
			"/api/v1/post",
			c.CreateCommunityPost,
		},
		"DeletePost": Route{
			strings.ToUpper("Delete"),
			"/api/v1/post/{uid}",
			c.DeletePost,
		},
		"GetPost": Route{
			strings.ToUpper("Get"),
			"/api/v1/post/{uid}",
			c.GetPost,
		},
		"GetPosts": Route{
			strings.ToUpper("Get"),
			"/api/v1/post",
			c.GetPosts,
		},
		"UpdatePost": Route{
			strings.ToUpper("Put"),
			"/api/v1/post/{uid}",
			c.UpdatePost,
		},
	}
}

// CreateCommunityPost - 
func (c *PostAPIController) CreateCommunityPost(w http.ResponseWriter, r *http.Request) {
	query, err := parseQuery(r.URL.RawQuery)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	var communityParam string
	if query.Has("community") {
		param := query.Get("community")

		communityParam = param
	} else {
		c.errorHandler(w, r, &RequiredError{Field: "community"}, nil)
		return
	}
	postRequestParam := PostRequest{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&postRequestParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertPostRequestRequired(postRequestParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertPostRequestConstraints(postRequestParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.CreateCommunityPost(r.Context(), communityParam, postRequestParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// DeletePost - 
func (c *PostAPIController) DeletePost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	uidParam := params["uid"]
	if uidParam == "" {
		c.errorHandler(w, r, &RequiredError{"uid"}, nil)
		return
	}
	result, err := c.service.DeletePost(r.Context(), uidParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// GetPost - 
func (c *PostAPIController) GetPost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	uidParam := params["uid"]
	if uidParam == "" {
		c.errorHandler(w, r, &RequiredError{"uid"}, nil)
		return
	}
	result, err := c.service.GetPost(r.Context(), uidParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// GetPosts - 
func (c *PostAPIController) GetPosts(w http.ResponseWriter, r *http.Request) {
	query, err := parseQuery(r.URL.RawQuery)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	var communityParam string
	if query.Has("community") {
		param := query.Get("community")

		communityParam = param
	} else {
	}
	result, err := c.service.GetPosts(r.Context(), communityParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// UpdatePost - 
func (c *PostAPIController) UpdatePost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	uidParam := params["uid"]
	if uidParam == "" {
		c.errorHandler(w, r, &RequiredError{"uid"}, nil)
		return
	}
	postRequestParam := PostRequest{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&postRequestParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertPostRequestRequired(postRequestParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertPostRequestConstraints(postRequestParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.UpdatePost(r.Context(), uidParam, postRequestParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}