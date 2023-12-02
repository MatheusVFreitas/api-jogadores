package models

type Jogador struct {
	Nome    string `json: "nome"`
	Time    string `json: "time"`
	Posicao string `json: "posicao"`
	Numero  int    `json: "numero"`
	Pais    string `json: "pais"`
}

var Jogadores []Jogador
