package api

import (
	"net/http"

	"github.com/gorilla/mux"

	// Import handlers from other files
	"family-tree-api/family-tree-api/pkg/api/handlers"
)

func Routes() *mux.Router {
	router := mux.NewRouter()

	// Define API routes and handlers

	// Family Tree Routes
	router.HandleFunc("/family-trees", handlers.CreateFamilyTreeHandler).Methods(http.MethodPost)
	// router.HandleFunc("/family-trees/{treeID}", handlers.GetFamilyTreeHandler).Methods(http.MethodGet)

	// Person Routes
	// router.HandleFunc("/family-trees/{treeID}/persons", handlers.AddPersonHandler).Methods(http.MethodPost)
	// router.HandleFunc("/family-trees/{treeID}/persons/{personID}", handlers.GetPersonHandler).Methods(http.MethodGet)

	// Relationship Routes
	// router.HandleFunc("/family-trees/{treeID}/relationships", handlers.FindRelationshipHandler).Methods(http.MethodGet)

	return router
}
