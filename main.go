package main

import (
	"github.com/LucasMelo59/upvoter-go/database"
	"github.com/LucasMelo59/upvoter-go/routes"
)

func main() {

	database.ConectaComBancoDeDados()
	routes.HandleResquest()
}
