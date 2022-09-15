package routes

import (
	"github.com/LucasMelo59/upvoter-go/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)





func HandleResquest() {
	r := gin.Default()
	r.Use(cors.Default())
	r.GET("/moedas", controllers.TodasMoedas) 
	r.GET("/moedas/:id", controllers.RetornaUmaMoeda) 
	r.POST("/moedas/criar", controllers.CriaUmaNovaMoeda)
	r.DELETE("/moedas/delete/:id", controllers.Deleta) 
	r.PATCH("/moedas/editar/:id", controllers.EditarMoeda) 
	r.GET("moedas/nome/:nome", controllers.BuscaMoedaPorNome)
	r.GET("moedas/upvoter/:id", controllers.Upvoter)
	r.GET("moedas/downvoter/:id", controllers.Downvoter)
	r.Run()
}

