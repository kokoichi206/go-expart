/*
 * Pet store schema
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"github.com/gin-gonic/gin"
)

type PetAPI struct {
	// Post /v3/pet
	// Add a new pet to the store
	AddPet gin.HandlerFunc
	// Delete /v3/pet/:petId
	// Deletes a pet
	DeletePet gin.HandlerFunc
	// Get /v3/pet/:petId
	// Find pet by ID
	GetPetById gin.HandlerFunc
	// Post /v3/pet/:petId
	// Updates a pet in the store
	UpdatePet gin.HandlerFunc
}
