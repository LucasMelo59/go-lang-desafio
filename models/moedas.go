package models

import (
	
)

type Moeda struct {
	Id int  `json:"id"`
	Nome      string  `json:"nome"`
	Descricao string  `json:"descricao"`
	Preco     float64 `json:"preco"`
	Votos     int     `json:"votos"`
	Src string  `json:"src"`
	}
