package controllers

import (
	"net/http"
	"github.com/LucasMelo59/upvoter-go/database"
	"github.com/LucasMelo59/upvoter-go/models"
	"github.com/gin-gonic/gin"

)




func TodasMoedas(c *gin.Context){
	var moedas []models.Moeda
	database.DB.Find(&moedas)
	c.JSON(200, moedas)
}

func RetornaUmaMoeda(c * gin.Context) {
	var moeda models.Moeda
	id := c.Params.ByName("id")
	database.DB.First(&moeda , id)
	if moeda.Id == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Moeda não encontrada"})
			return
	}
	c.JSON(http.StatusOK, moeda)
}

func CriaUmaNovaMoeda(c *gin.Context) {
	var moeda models.Moeda
	if err := c.ShouldBindJSON(&moeda); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error()})
				return 
		}
			database.DB.Create(&moeda)
			c.JSON(http.StatusOK, moeda)
		}
	



func Deleta(c *gin.Context){
	var moeda models.Moeda
	id := c.Params.ByName("id")
	database.DB.First(&moeda , id)
	if moeda.Id == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Moeda não encontrada"})
			return
	}
	database.DB.Delete(&moeda , id)
	c.JSON(http.StatusOK, gin.H{"data": "Moeda Deletada"})
}

func EditarMoeda(c *gin.Context) {
	var moeda models.Moeda
	id := c.Params.ByName("id")
	database.DB.First(&moeda , id)
	if err := c.ShouldBindJSON(&moeda); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error()})
				return 
		}
	database.DB.Model(&moeda).UpdateColumns(moeda)
	c.JSON(http.StatusOK, moeda)
}

func Upvoter(c *gin.Context){
	var moeda models.Moeda
	id := c.Params.ByName("id")
	database.DB.First(&moeda , id)
	if moeda.Id == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Moeda não encontrada"})
			return
	}
	moeda.Votos = moeda.Votos + 1
	database.DB.Model(&moeda).UpdateColumns(moeda)
	
	}

	func Downvoter(c *gin.Context){
		var moeda models.Moeda
		id := c.Params.ByName("id")
		database.DB.First(&moeda , id)
		if moeda.Id == 0 {
			c.JSON(http.StatusNotFound, gin.H{
				"Not found": "Moeda não encontrada"})
				return
		}
		moeda.Votos = moeda.Votos - 1
		database.DB.Model(&moeda).UpdateColumns(moeda)
		
		}


	func BuscaMoedaPorNome(c *gin.Context) {
		var moeda models.Moeda
		name := c.Param("name")
		database.DB.Where(&models.Moeda{Nome: name}).First(&moeda)
		if moeda.Id == 0 {
			c.JSON(http.StatusNotFound, gin.H{
				"Not found": "Moeda não encontrada"})
				return
		}
		c.JSON(http.StatusOK, moeda)
	}

	



