package controllers

import (
	"crypto/rand"
	"encoding/hex"
	"golang-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// get all
func GetUsers(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var users []models.Users
	db.Find(&users)

	c.JSON(http.StatusOK, gin.H{
		"status":   1,
		"responce": 200,
		"data":     users,
	})
}

// Create new
func CreateUsers(c *gin.Context) {
	var input models.Users
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	b := make([]byte, 8)
	_, err := rand.Read(b)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	passwordEncrypted := hex.EncodeToString(b)

	task := models.Users{Username: input.Username, Password: passwordEncrypted}

	db := c.MustGet("db").(*gorm.DB)
	db.Create(&task)

	c.JSON(http.StatusOK, gin.H{
		"status":   1,
		"response": 200,
		"data":     task,
	})
}
