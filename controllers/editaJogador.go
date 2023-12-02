package controllers

import (
	"net/http"

	"github.com/MatheusVFreitas/go-api-jogadores/database"
	"github.com/MatheusVFreitas/go-api-jogadores/models"
	"github.com/gin-gonic/gin"
)

func EditaJogador(c *gin.Context) {
	var jogador models.Jogador
	id := c.Params.ByName("id")
	database.DB.First(&jogador, id)

	if err := c.ShouldBindJSON(&jogador); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro":err.Error(),
		})
		return
	}
	database.DB.Model(&jogador).UpdateColumns(jogador)
	c.JSON(http.StatusOK, jogador)
}