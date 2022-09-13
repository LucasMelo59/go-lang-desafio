package routes

import (
	"github.com/LucasMelo59/upvoter-go/controllers"
	"github.com/gin-gonic/gin"
)

func HandleResquest() {
	r := gin.Default()
	r.GET("/moedas", controllers.TodasMoedas)
	r.GET("/moedas:id", controllers.RetornaUmaMoeda)
	r.POST("/moedas", controllers.CriaUmaNovaMoeda)
	r.DELETE("/moedas:id", controllers.Deleta)
	r.PATCH("/moedas:id", controllers.EditarMoeda)
	r.GET("moedas/nome/:nome", controllers.BuscaMoedaPorNome)
	r.GET("/upvoter/:id", controllers.Upvoter)
	r.Run()
}
