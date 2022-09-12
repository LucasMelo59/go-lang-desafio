package routes

import (
	"log"
	"net/http"

	"github.com/LucasMelo59/upvoter-go/controllers"
	"github.com/LucasMelo59/upvoter-go/middleware"
	"github.com/gorilla/mux"
)

func HandleResquest() {
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleWare)
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/moedas", controllers.TodasMoedas).Methods("Get")
	r.HandleFunc("/api/moedas/{id}", controllers.RetornaUmaMoeda).Methods("Get")
	r.HandleFunc("/api/moedas", controllers.CriaUmaNovaMoeda).Methods("Post")
	r.HandleFunc("/api/moedas/{id}", controllers.CriaUmaNovaMoeda).Methods("Delete")
	r.HandleFunc("/api/moedas/{id}", controllers.EditarMoeda).Methods("Put")

	log.Fatal(http.ListenAndServe(":8000", r))
}
