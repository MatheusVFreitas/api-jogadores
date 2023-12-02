package controllers

import (
	"net/http"

	"github.com/MatheusVFreitas/go-api-jogadores/database"
	"github.com/MatheusVFreitas/go-api-jogadores/models"
	"github.com/gin-gonic/gin"
)

func CriaJogador(c *gin.Context) {
	var jogador models.Jogador
	if err := c.ShouldBindJSON(&jogador); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Erro": err.Error(),
		})
		return
	}
	database.DB.Create(&jogador)
	c.JSON(http.StatusOK, jogador)
}
