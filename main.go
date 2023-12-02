package main

import (
	"github.com/MatheusVFreitas/go-api-jogadores/database"
	"github.com/MatheusVFreitas/go-api-jogadores/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequests()
}
