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

// CategoryAPIController binds http requests to an api service and writes the service results to the http response
type CategoryAPIController struct {
	service CategoryAPIServicer
	errorHandler ErrorHandler
}

// CategoryAPIOption for how the controller is set up.
type CategoryAPIOption func(*CategoryAPIController)

// WithCategoryAPIErrorHandler inject ErrorHandler into controller
func WithCategoryAPIErrorHandler(h ErrorHandler) CategoryAPIOption {
	return func(c *CategoryAPIController) {
		c.errorHandler = h
	}
}

// NewCategoryAPIController creates a default api controller
func NewCategoryAPIController(s CategoryAPIServicer, opts ...CategoryAPIOption) Router {
	controller := &CategoryAPIController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the CategoryAPIController
func (c *CategoryAPIController) Routes() Routes {
	return Routes{
		"CreateCategory": Route{
			strings.ToUpper("Post"),
			"/api/v1/category",
			c.CreateCategory,
		},
		"GetAllCategories": Route{
			strings.ToUpper("Get"),
			"/api/v1/category",
			c.GetAllCategories,
		},
	}
}

// CreateCategory - 
func (c *CategoryAPIController) CreateCategory(w http.ResponseWriter, r *http.Request) {
	categoryCreateParam := CategoryCreate{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&categoryCreateParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertCategoryCreateRequired(categoryCreateParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertCategoryCreateConstraints(categoryCreateParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.CreateCategory(r.Context(), categoryCreateParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// GetAllCategories - 
func (c *CategoryAPIController) GetAllCategories(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.GetAllCategories(r.Context())
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}