package controllers

import (
	"api-trunfo/database"
	"api-trunfo/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllVeiculos(c *gin.Context) {
	var v []models.Veiculos
	database.DB.Find(&v)
	c.JSON(http.StatusOK, v)
}

func CreateVeiculos(c *gin.Context) {
	var v models.Veiculos
	if err := c.ShouldBindJSON(&v); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	database.DB.Create(&v)
	c.JSON(http.StatusOK, v)
}

func GetIDVeiculos(c *gin.Context) {
	var v models.Veiculos
	id := c.Params.ByName("id")
	database.DB.First(&v, id)

	if v.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"not found": "Veiculo nao encontrado",
		})
		return
	}

	c.JSON(http.StatusOK, v)
}

func DeletaVeiculo(c *gin.Context) {
	var v []models.Veiculos
	id := c.Params.ByName("id")

	database.DB.Delete(&v, id)

	c.JSON(http.StatusOK, gin.H{
		"message": "Trunfo deletado com sucesso",
	})
}

func EditaVeiculo(c *gin.Context) {
	var v models.Veiculos
	id := c.Params.ByName("id")

	database.DB.First(&v, id)

	if err := c.ShouldBindJSON(&v); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	database.DB.Model(&v).UpdateColumns(v)
	c.JSON(http.StatusOK, v)
}
