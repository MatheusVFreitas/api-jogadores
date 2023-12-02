package main

import (
	"github.com/MatheusVFreitas/go-api-jogadores/models"
	"github.com/MatheusVFreitas/go-api-jogadores/routes"
)

func main() {
	models.Jogadores = []models.Jogador{
		{Nome: "Messi", Time: "Barcelona", Posicao: "Atacante", Numero: 10, Pais: "Argentina"},
		{Nome: "Xavi", Time: "Barcelona", Posicao: "Meiocampista", Numero: 6, Pais: "Espanha"},
		{Nome: "Iniesta", Time: "Barcelona", Posicao: "Meiocampista", Numero: 8, Pais: "Espanha"},
	}
	routes.HandleRequests()
}
