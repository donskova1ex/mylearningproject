// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Swagger Potion maker - OpenAPI 3.0
 *
 * Сервис изготовления зелий
 *
 * API version: 1.0.0
 */

package openapi

import (
	"encoding/json"
	"net/http"
	"strings"
	"errors"
	"github.com/gorilla/mux"
	"io"
)

// IngredientAPIController binds http requests to an api service and writes the service results to the http response
type IngredientAPIController struct {
	service IngredientAPIServicer
	errorHandler ErrorHandler
}

// IngredientAPIOption for how the controller is set up.
type IngredientAPIOption func(*IngredientAPIController)

// WithIngredientAPIErrorHandler inject ErrorHandler into controller
func WithIngredientAPIErrorHandler(h ErrorHandler) IngredientAPIOption {
	return func(c *IngredientAPIController) {
		c.errorHandler = h
	}
}

// NewIngredientAPIController creates a default api controller
func NewIngredientAPIController(s IngredientAPIServicer, opts ...IngredientAPIOption) *IngredientAPIController {
	controller := &IngredientAPIController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the IngredientAPIController
func (c *IngredientAPIController) Routes() Routes {
	return Routes{
		"IngredientsByName": Route{
			strings.ToUpper("Get"),
			"/api/v1/ingredient",
			c.IngredientsByName,
		},
		"IngredientsList": Route{
			strings.ToUpper("Get"),
			"/api/v1/ingredients",
			c.IngredientsList,
		},
		"GetIngredientById": Route{
			strings.ToUpper("Get"),
			"/api/v1/ingredient/{id}",
			c.GetIngredientById,
		},
		"UpdateIngredient": Route{
			strings.ToUpper("Put"),
			"/api/v1/ingredient/{id}",
			c.UpdateIngredient,
		},
		"DeleteIngredient": Route{
			strings.ToUpper("Delete"),
			"/api/v1/ingredient/{id}",
			c.DeleteIngredient,
		},
	}
}

// IngredientsByName - Ingredient by name
func (c *IngredientAPIController) IngredientsByName(w http.ResponseWriter, r *http.Request) {
	query, err := parseQuery(r.URL.RawQuery)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	var nameParam string
	if query.Has("name") {
		param := query.Get("name")

		nameParam = param
	} else {
	}
	var idParam string
	if query.Has("id") {
		param := query.Get("id")

		idParam = param
	} else {
	}
	result, err := c.service.IngredientsByName(r.Context(), nameParam, idParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = EncodeJSONResponse(result.Body, &result.Code, w)
}

// IngredientsList - Ingredients list
func (c *IngredientAPIController) IngredientsList(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.IngredientsList(r.Context())
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = EncodeJSONResponse(result.Body, &result.Code, w)
}

// GetIngredientById - Find ingredient by ID
func (c *IngredientAPIController) GetIngredientById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idParam := params["id"]
	if idParam == "" {
		c.errorHandler(w, r, &RequiredError{"id"}, nil)
		return
	}
	result, err := c.service.GetIngredientById(r.Context(), idParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = EncodeJSONResponse(result.Body, &result.Code, w)
}

// UpdateIngredient - Update a ingredient by id
func (c *IngredientAPIController) UpdateIngredient(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idParam := params["id"]
	if idParam == "" {
		c.errorHandler(w, r, &RequiredError{"id"}, nil)
		return
	}
	ingredientParam := Ingredient{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&ingredientParam); err != nil && !errors.Is(err, io.EOF) {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertIngredientRequired(ingredientParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertIngredientConstraints(ingredientParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.UpdateIngredient(r.Context(), idParam, ingredientParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = EncodeJSONResponse(result.Body, &result.Code, w)
}

// DeleteIngredient - Delete ingredient
func (c *IngredientAPIController) DeleteIngredient(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idParam, err := parseNumericParameter[int64](
		params["id"],
		WithRequire[int64](parseInt64),
	)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Param: "id", Err: err}, nil)
		return
	}
	result, err := c.service.DeleteIngredient(r.Context(), idParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = EncodeJSONResponse(result.Body, &result.Code, w)
}