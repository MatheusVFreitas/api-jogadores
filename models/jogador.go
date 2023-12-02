package models

import "gorm.io/gorm"

type Jogador struct {
	gorm.Model
	Nome    string `json: "nome"`
	Time    string `json: "time"`
	Posicao string `json: "posicao"`
	Numero  int    `json: "numero"`
	Pais    string `json: "pais"`
}

var Jogadores []Jogador
