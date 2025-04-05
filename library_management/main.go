package main

import (
	"library_management/controllers"
	"library_management/services"
)

func main() {
	libraryService := services.NewLibrary()
	libraryController := controllers.NewLibraryController(libraryService)

	libraryController.Landing()
}
