package main

import (
	"github.com/LucasMelo59/upvoter-go/database"
	"github.com/LucasMelo59/upvoter-go/models"
	"github.com/LucasMelo59/upvoter-go/routes"
)

func main() {

	models.Moedas = []models.Moeda{}
	database.ConectaComBancoDeDados()
	routes.HandleResquest()
}
