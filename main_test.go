package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/LucasMelo59/upvoter-go/controllers"
	"github.com/LucasMelo59/upvoter-go/database"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetupDasRotasDeTeste() *gin.Engine {
	rotas := gin.Default()
	return rotas

}

func TestListandoTodasAsMoedas(t *testing.T) {
	database.ConectaComBancoDeDados()
	r := SetupDasRotasDeTeste()
	r.GET("/moedas", controllers.TodasMoedas)
	req, _ := http.NewRequest("GET", "/moedas", nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)
}
