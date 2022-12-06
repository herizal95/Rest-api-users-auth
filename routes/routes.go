package routes

import (
	"golang-api/controllers"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func SetRoutes(db *gorm.DB) *gin.Engine {
	routes := gin.Default()
	routes.Use(func(ctx *gin.Context) {
		ctx.Set("db", db)
	})

	routes.GET("/users", controllers.GetUsers)
	routes.POST("/users", controllers.CreateUsers)

	return routes
}
