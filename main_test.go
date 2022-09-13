package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/LucasMelo59/upvoter-go/controllers"
	"github.com/LucasMelo59/upvoter-go/database"
	"github.com/LucasMelo59/upvoter-go/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var ID int

func SetupDasRotasDeTeste() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	rotas := gin.Default()
	return rotas

}

func CriaMoedaMoc() {
	moeda := models.Moeda{Nome: "teste", Descricao: "hahaha", Preco: 22.90, Votos: 5 }
	database.DB.Create(&moeda)
	ID = int(moeda.ID)
}

func DeletaMoedaMock(){
var moeda models.Moeda
database.DB.Delete(&moeda, ID)

}

func TestListandoTodasAsMoedas(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriaMoedaMoc()
	defer DeletaMoedaMock()
	r := SetupDasRotasDeTeste()
	r.GET("/moedas", controllers.TodasMoedas)
	req, _ := http.NewRequest("GET", "/moedas", nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)
}


func TestBuscarPorNome(t *testing.T){
database.ConectaComBancoDeDados()
CriaMoedaMoc()
defer DeletaMoedaMock()
r := SetupDasRotasDeTeste()
r.GET("moedas/nome/:nome", controllers.BuscaMoedaPorNome)
req, _ := http.NewRequest("GET", "/moedas/nome/teste", nil)
resp := httptest.NewRecorder()
r.ServeHTTP(resp, req)
assert.Equal(t, http.StatusOK, resp.Code)

}


func TestBuscarMoedaPorId(t *testing.T){

database.ConectaComBancoDeDados()
CriaMoedaMoc()
defer DeletaMoedaMock()
r := SetupDasRotasDeTeste()
r.GET("/moedas:id", controllers.RetornaUmaMoeda)
pathDeBusca := "/alunos/" + strconv.Itoa(ID)
req, _ := http.NewRequest("GET", pathDeBusca,nil)
resp := httptest.NewRecorder()
r.ServeHTTP(resp, req)
var moedaMock models.Moeda
json.Unmarshal(resp.Body.Bytes(), &moedaMock)
assert.Equal(t, "teste", moedaMock.Nome )
}


func TestDeleta(t *testing.T){
database.ConectaComBancoDeDados()
CriaMoedaMoc()
defer DeletaMoedaMock()
r := SetupDasRotasDeTeste()
r.DELETE("/delete/moedas:id", controllers.Deleta)
pathDeBusca := "/moedas/" + strconv.Itoa(ID)
req , _ := http.NewRequest("DELETE", pathDeBusca, nil)
resp:= httptest.NewRecorder()
r.ServeHTTP(resp, req)
assert.Equal(t, http.StatusOK, resp.Code)
}


func TestUpdate(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriaMoedaMoc()
	defer DeletaMoedaMock()
	r := SetupDasRotasDeTeste()
	r.PATCH("/editar/moedas:id", controllers.EditarMoeda)
	moeda := models.Moeda{Nome: "teste", Descricao: "hahaha", Preco: 22.00, Votos: 8 }
	valorJson , _ := json.Marshal(moeda)
	pathParaEditar := "/moedas/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("PATCH", pathParaEditar, bytes.NewBuffer(valorJson))
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)
	var moedaMocAtualizada models.Moeda
	json.Unmarshal(resp.Body.Bytes(), &moedaMocAtualizada)
	assert.Equal(t, " 8", moedaMocAtualizada.Votos)
	assert.Equal(t, " 22.00", moedaMocAtualizada.Preco)
}

func TestUpvoter(t *testing.T){

	database.ConectaComBancoDeDados()
	CriaMoedaMoc()
	 
	r := SetupDasRotasDeTeste()
	r.GET("/upvoter/:id", controllers.Upvoter)
	pathDeBusca := "/upvoter" + strconv.Itoa(ID)
	req, _ := http.NewRequest("GET", pathDeBusca,nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)
	var moedaMock models.Moeda
	json.Unmarshal(resp.Body.Bytes(), &moedaMock)
	assert.Equal(t, 6 , moedaMock.Votos )
	DeletaMoedaMock()
}