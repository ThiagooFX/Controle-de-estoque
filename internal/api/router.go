package api

import (
	"api/internal/services"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func InitRouters() {
	r := mux.NewRouter()
	r.HandleFunc("/itens", services.GetItens).Methods("GET")
	r.HandleFunc("/itens", services.AddItem).Methods("POST")
	r.HandleFunc("/itens/{id}", services.UpdateItem).Methods("PUT")
	r.HandleFunc("/itens/{id}", services.DeleteItem).Methods("DELETE")

	log.Println("Servidor rodando na porta 2000")
	http.ListenAndServe(":2000", r)
}
