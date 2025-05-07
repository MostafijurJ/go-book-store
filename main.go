package main

import (
	"github.com/gorilla/mux"
	_ "github.com/mostafijurj/go-book-store/docs"
	_ "github.com/mostafijurj/go-book-store/pkg/controllers" // Import controllers
	"github.com/mostafijurj/go-book-store/pkg/routes"
	httpSwagger "github.com/swaggo/http-swagger"
	"log"
	"net/http"
)

// @title Go Book Store API
// @description This is the API documentation for the Go Book Store application.
// @version 1.0
// @host go-book-store-1ewq.onrender.com
// @BasePath /api/v1

// @contact.name API Support
// @contact.url https://www.example.com/support
// @contact.email support@example.com
func main() {
	router := mux.NewRouter()

	// Create a subroutine for versioned API
	api := router.PathPrefix("/api/v1").Subrouter()

	// Register routes on the subroutine
	routes.RegisterBookRoutes(api)

	// Swagger route (still accessible directly)
	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	log.Println("Server started at https://go-book-store-1ewq.onrender.com/")
	log.Println("Swagger docs at https://go-book-store-1ewq.onrender.com/swagger/index.html")
	log.Println("API base path: https://go-book-store-1ewq.onrender.com/api/v1")

	// Start server
	log.Fatal(http.ListenAndServe(":8888", router))
}
