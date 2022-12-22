package controllers

import (
	"api-trunfo/database"
	"api-trunfo/models"
	"api-trunfo/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Nao foi possivel vincular o JSON",
		})
		return
	}

	user.Password = services.SHA256Encoder(user.Password)

	err = database.DB.Create(&user).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Nao foi opssivel criar o usuario " + err.Error(),
		})
		return
	}

	c.Status(http.StatusCreated)
}

func GetAllUser(c *gin.Context) {
	var user []models.User
	database.DB.Find(&user)
	c.JSON(http.StatusOK, user)
}
