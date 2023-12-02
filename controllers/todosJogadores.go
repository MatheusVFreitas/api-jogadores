package controllers

import (
	"net/http"

	"github.com/MatheusVFreitas/go-api-jogadores/models"
	"github.com/gin-gonic/gin"
)

func TodosJogadores(c *gin.Context) {
	c.JSON(http.StatusOK, models.Jogadores)
}
