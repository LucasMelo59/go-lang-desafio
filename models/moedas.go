package models

import "gorm.io/gorm"

type Moeda struct {
	gorm.Model
	Nome      string  `json:"nome"`
	Descricao string  `json:"descricao"`
	Preco     float64 `json:"preco"`
	Votos     int     `json:"votos"`
}

var Moedas []Moeda