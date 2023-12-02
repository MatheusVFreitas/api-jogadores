package controllers

import (
	"net/http"

	"github.com/MatheusVFreitas/go-api-jogadores/database"
	"github.com/MatheusVFreitas/go-api-jogadores/models"
	"github.com/gin-gonic/gin"
)

func BuscaJogadorPorId(c *gin.Context) {
	var jogador models.Jogador
	id := c.Params.ByName("id")
	database.DB.First(&jogador, id)

	if jogador.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"Info": "Jogador não encontrado",
		})
		return
	}
	c.JSON(http.StatusOK, jogador)
}

func BuscaJogadorPorTime(c *gin.Context) {
	var jogador []models.Jogador
	time := c.Params.ByName("time")
	database.DB.Where(&models.Jogador{Time: time}).Find(&jogador)

	
	c.JSON(http.StatusOK, jogador)
}