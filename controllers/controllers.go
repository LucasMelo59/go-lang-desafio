package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/LucasMelo59/upvoter-go/database"
	"github.com/LucasMelo59/upvoter-go/models"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "homepage")

}


func TodasMoedas(w http.ResponseWriter, r*http.Request){
	var moeda []models.Moeda
	database.DB.Find(&moeda)
	json.NewEncoder(w).Encode(moeda)

}

func RetornaUmaMoeda(w http.ResponseWriter, r*http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]
	var moeda models.Moeda
	database.DB.First(&moeda, id)
	json.NewEncoder(w).Encode(moeda)
}

func CriaUmaNovaMoeda(w http.ResponseWriter, r*http.Request) {
	var moeda models.Moeda
	json.NewDecoder(r.Body).Decode(&moeda)
	database.DB.Create(&moeda)
	json.NewEncoder(w).Encode(moeda)
}

func Deleta(w http.ResponseWriter, r*http.Request){
	vars := mux.Vars(r)
	id := vars["id"]
	var moeda models.Moeda
	database.DB.Delete(&moeda, id)
	json.NewEncoder(w).Encode(moeda)
}

func EditarMoeda(w http.ResponseWriter, r*http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var moeda models.Moeda
	database.DB.First(&moeda, id)
	json.NewDecoder(r.Body).Decode(&moeda)
	database.DB.Save(moeda)
	json.NewEncoder(w).Encode(moeda)

}

func Upvoter(w http.ResponseWriter, r*http.Request){
	vars := mux.Vars(r)
	id := vars["id"]
	var moeda models.Moeda
	database.DB.First(&moeda, id)
	if moeda.Id == 0 {
		
	}

	}


	



