package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	_ "modernc.org/sqlite"
)

type Item struct {
	ID        int    `json:"id"`
	Nome      string `json:"nome"`
	Quantidade int   `json:"quantidade"`
}

var db *sql.DB

func initDB() {
	var err error
	db, err = sql.Open("sqlite", "estoque.db")
	if err != nil {
		log.Fatal(err)
	}

	createTable := `CREATE TABLE IF NOT EXISTS itens (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		nome TEXT NOT NULL,
		quantidade INTEGER NOT NULL
	);`
	if _, err := db.Exec(createTable); err != nil {
		log.Fatal(err)
	}
}

func enableCORS(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

func getItens(w http.ResponseWriter, r *http.Request) {
	enableCORS(w)
	rows, err := db.Query("SELECT id, nome, quantidade FROM itens")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var itens []Item
	for rows.Next() {
		var item Item
		if err := rows.Scan(&item.ID, &item.Nome, &item.Quantidade); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		itens = append(itens, item)
	}

	json.NewEncoder(w).Encode(itens)
}

func addItem(w http.ResponseWriter, r *http.Request) {
	enableCORS(w)
	var item Item
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := db.Exec("INSERT INTO itens (nome, quantidade) VALUES (?, ?)", item.Nome, item.Quantidade)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, _ := result.LastInsertId()
	item.ID = int(id)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(item)
}

func updateItem(w http.ResponseWriter, r *http.Request) {
	enableCORS(w)
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var item Item
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = db.Exec("UPDATE itens SET nome=?, quantidade=? WHERE id=?", item.Nome, item.Quantidade, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	item.ID = id
	json.NewEncoder(w).Encode(item)
}

func deleteItem(w http.ResponseWriter, r *http.Request) {
	enableCORS(w)
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	_, err = db.Exec("DELETE FROM itens WHERE id=?", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func main() {
	initDB()
	r := mux.NewRouter()
	r.HandleFunc("/itens", getItens).Methods("GET")
	r.HandleFunc("/itens", addItem).Methods("POST")
	r.HandleFunc("/itens/{id}", updateItem).Methods("PUT")
	r.HandleFunc("/itens/{id}", deleteItem).Methods("DELETE")

	log.Println("Servidor rodando na porta 2000")
	http.ListenAndServe(":2000", r)
}