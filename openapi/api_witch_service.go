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
	"context"
	"net/http"
	"errors"
)

// WitchAPIService is a service that implements the logic for the WitchAPIServicer
// This service should implement the business logic for every endpoint for the WitchAPI API.
// Include any external packages or services that will be required by this service.
type WitchAPIService struct {
}

// NewWitchAPIService creates a default api service
func NewWitchAPIService() *WitchAPIService {
	return &WitchAPIService{}
}

// WitchesList - witches list
func (s *WitchAPIService) WitchesList(ctx context.Context) (ImplResponse, error) {
	// TODO - update WitchesList with the required logic for this service method.
	// Add api_witch_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, []Witch{}) or use other options such as http.Ok ...
	// return Response(200, []Witch{}), nil

	// TODO: Uncomment the next line to return response Response(204, {}) or use other options such as http.Ok ...
	// return Response(204, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("WitchesList method not implemented")
}

// GetWitch - Find witch by parametrs
func (s *WitchAPIService) GetWitch(ctx context.Context, id string, name string) (ImplResponse, error) {
	// TODO - update GetWitch with the required logic for this service method.
	// Add api_witch_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, Witch{}) or use other options such as http.Ok ...
	// return Response(200, Witch{}), nil

	// TODO: Uncomment the next line to return response Response(400, {}) or use other options such as http.Ok ...
	// return Response(400, nil),nil

	// TODO: Uncomment the next line to return response Response(404, {}) or use other options such as http.Ok ...
	// return Response(404, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetWitch method not implemented")
}

// GetWitchById - Find witch by ID
func (s *WitchAPIService) GetWitchById(ctx context.Context, id string) (ImplResponse, error) {
	// TODO - update GetWitchById with the required logic for this service method.
	// Add api_witch_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, Witch{}) or use other options such as http.Ok ...
	// return Response(200, Witch{}), nil

	// TODO: Uncomment the next line to return response Response(400, {}) or use other options such as http.Ok ...
	// return Response(400, nil),nil

	// TODO: Uncomment the next line to return response Response(404, {}) or use other options such as http.Ok ...
	// return Response(404, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetWitchById method not implemented")
}

// UpdateWitchWithForm - Update a witch in the store with form data
func (s *WitchAPIService) UpdateWitchWithForm(ctx context.Context, id string) (ImplResponse, error) {
	// TODO - update UpdateWitchWithForm with the required logic for this service method.
	// Add api_witch_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, Witch{}) or use other options such as http.Ok ...
	// return Response(200, Witch{}), nil

	// TODO: Uncomment the next line to return response Response(400, {}) or use other options such as http.Ok ...
	// return Response(400, nil),nil

	// TODO: Uncomment the next line to return response Response(404, {}) or use other options such as http.Ok ...
	// return Response(404, nil),nil

	// TODO: Uncomment the next line to return response Response(422, {}) or use other options such as http.Ok ...
	// return Response(422, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("UpdateWitchWithForm method not implemented")
}

// DeleteWitch - Deletes a witch
func (s *WitchAPIService) DeleteWitch(ctx context.Context, id int64) (ImplResponse, error) {
	// TODO - update DeleteWitch with the required logic for this service method.
	// Add api_witch_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(400, {}) or use other options such as http.Ok ...
	// return Response(400, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("DeleteWitch method not implemented")
}