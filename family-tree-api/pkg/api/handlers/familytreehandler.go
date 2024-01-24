package handlers

import (
	"encoding/json"
	"net/http"
	// "pkg/application"
	// "pkg/domain"
	// "pkg/infrastructure"
)

// CreateFamilyTreeHandler handles creating a new family tree
func CreateFamilyTreeHandler(w http.ResponseWriter, r *http.Request) {
	// Decode request body into a FamilyTree struct
	// var newTree domain.FamilyTree
	// err := json.NewDecoder(r.Body).Decode(&newTree)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	return
	// }

	// // Create a new FamilyTreeService instance
	// familyTreeService := application.NewFamilyTreeService(infrastructure.NewPersonRepository())

	// // Call the service to create the family tree
	// tree, err := familyTreeService.CreateFamilyTree(&newTree)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	// Encode the created tree as JSON and send it as the response
	json.NewEncoder(w).Encode("Oi Familia!")
}
