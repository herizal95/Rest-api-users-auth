package main

import (
	"golang-api/models"
	"golang-api/routes"
	"golang-api/services"
)

func main() {
	db := services.SetupDB()
	db.AutoMigrate(&models.Users{})

	routes := routes.SetRoutes(db)
	routes.Run()
}
