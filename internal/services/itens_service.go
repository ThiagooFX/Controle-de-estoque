package services

import (
	"api/internal/models"
	"api/internal/repository"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func enableCORS(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

func GetItens(w http.ResponseWriter, r *http.Request) {
	enableCORS(w)
	rows, err := repository.DB.Query("SELECT id, name, quantity FROM itens")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var itens []models.Item
	for rows.Next() {
		var item models.Item
		if err := rows.Scan(&item.ID, &item.Name, &item.Quantity); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		itens = append(itens, item)
	}

	json.NewEncoder(w).Encode(itens)
}

func AddItem(w http.ResponseWriter, r *http.Request) {
	enableCORS(w)

	var item models.Item
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := repository.DB.Exec("INSERT INTO itens (name, quantity) VALUES (?, ?)", item.Name, item.Quantity)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		http.Error(w, "Erro ao obter ID do item", http.StatusInternalServerError)
		return
	}
	item.ID = int(id)

	// Responde ao cliente
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(item)

	// Salva log da ação
	go func() {
		err := Addlogs(r, fmt.Sprintf("Adicionou item ID %d", item.ID))
		if err != nil {
			log.Printf("Erro ao registrar log: %v", err)
		}
	}()
}

func UpdateItem(w http.ResponseWriter, r *http.Request) {
	enableCORS(w)
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var item models.Item
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = repository.DB.Exec("UPDATE itens SET name=?, quantity=? WHERE id=?", item.Name, item.Quantity, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	item.ID = id
	json.NewEncoder(w).Encode(item)
}

func DeleteItem(w http.ResponseWriter, r *http.Request) {
	enableCORS(w)
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	_, err = repository.DB.Exec("DELETE FROM itens WHERE id=?", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
