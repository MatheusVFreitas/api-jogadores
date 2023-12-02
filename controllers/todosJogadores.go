package controllers

import (
	"net/http"

	"github.com/MatheusVFreitas/go-api-jogadores/database"
	"github.com/MatheusVFreitas/go-api-jogadores/models"
	"github.com/gin-gonic/gin"
)

func TodosJogadores(c *gin.Context) {
	var jogador []models.Jogador
	database.DB.Find(&jogador)
	c.JSON(http.StatusOK, jogador)
}
