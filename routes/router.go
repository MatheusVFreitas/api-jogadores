package routes

import (
	"github.com/MatheusVFreitas/go-api-jogadores/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/", controllers.Home)
	r.GET("/jogadores", controllers.TodosJogadores)
	r.Run()
}
