package controllers

import (
	"api-trunfo/database"
	"api-trunfo/models"
	"api-trunfo/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var login models.Login
	err := c.ShouldBindJSON(&login)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Nao foi enviado o body" + err.Error(),
		})
		return
	}

	var user models.User
	dbError := database.DB.Where("email = ?", login.Email).First(&user).Error
	if dbError != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "falha ao carregar o usuario",
		})
		return
	}

	if user.Password != services.SHA256Encoder(login.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Senha icnorreta",
		})
		return
	}

	token, err := services.NewJWTService().GenerateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Token": token,
	})
}
