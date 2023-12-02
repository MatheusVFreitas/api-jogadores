package controllers

import (
	"net/http"

	"github.com/MatheusVFreitas/go-api-jogadores/database"
	"github.com/MatheusVFreitas/go-api-jogadores/models"
	"github.com/gin-gonic/gin"
)

func ExcluiJogador(c *gin.Context) {
	var jogador models.Jogador
	id := c.Params.ByName("id")
	database.DB.Delete(&jogador, id)
	c.JSON(http.StatusOK, gin.H{
		"Info":"Jogador deletado com sucesso.",
	})
}
