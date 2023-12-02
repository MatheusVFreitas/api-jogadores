package routes

import (
	"github.com/MatheusVFreitas/go-api-jogadores/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/", controllers.Home)
	r.GET("/jogadores", controllers.TodosJogadores)
	r.GET("/jogadores/jogador/:id", controllers.BuscaJogadorPorId)
	r.GET("/jogadores/:time", controllers.BuscaJogadorPorTime)
	r.POST("/jogadores", controllers.CriaJogador)
	r.DELETE("/jogadores/:id", controllers.ExcluiJogador)
	r.PATCH("/jogador/:id", controllers.EditaJogador)
	r.Run()
}
