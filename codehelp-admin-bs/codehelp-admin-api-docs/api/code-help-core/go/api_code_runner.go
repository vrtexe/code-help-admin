/*
 * Coding helper spec
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package codehelp_core_api

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// CodeRunnerAPIController binds http requests to an api service and writes the service results to the http response
type CodeRunnerAPIController struct {
	service CodeRunnerAPIServicer
	errorHandler ErrorHandler
}

// CodeRunnerAPIOption for how the controller is set up.
type CodeRunnerAPIOption func(*CodeRunnerAPIController)

// WithCodeRunnerAPIErrorHandler inject ErrorHandler into controller
func WithCodeRunnerAPIErrorHandler(h ErrorHandler) CodeRunnerAPIOption {
	return func(c *CodeRunnerAPIController) {
		c.errorHandler = h
	}
}

// NewCodeRunnerAPIController creates a default api controller
func NewCodeRunnerAPIController(s CodeRunnerAPIServicer, opts ...CodeRunnerAPIOption) Router {
	controller := &CodeRunnerAPIController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the CodeRunnerAPIController
func (c *CodeRunnerAPIController) Routes() Routes {
	return Routes{
		"RunCode": Route{
			strings.ToUpper("Post"),
			"/api/run-code",
			c.RunCode,
		},
	}
}

// RunCode - 
func (c *CodeRunnerAPIController) RunCode(w http.ResponseWriter, r *http.Request) {
	runCodeRequestParam := RunCodeRequest{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&runCodeRequestParam); err != nil && !errors.Is(err, io.EOF) {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertRunCodeRequestRequired(runCodeRequestParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertRunCodeRequestConstraints(runCodeRequestParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.RunCode(r.Context(), runCodeRequestParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}
