package routes

import (
	"github.com/LucasMelo59/upvoter-go/controllers"
	"github.com/gin-gonic/gin"
)

func HandleResquest() {
	r := gin.Default()
	r.GET("/api/moedas", controllers.TodasMoedas)
	r.GET("/moedas:id", controllers.RetornaUmaMoeda)
	r.POST("/moedas/criar", controllers.CriaUmaNovaMoeda)
	r.DELETE("/delete/moedas:id", controllers.Deleta)
	r.PATCH("/editar/moedas:id", controllers.EditarMoeda)
	r.GET("moedas/nome/:nome", controllers.BuscaMoedaPorNome)
	r.GET("/upvoter/:id", controllers.Upvoter)
	r.Run(":3000")
}
