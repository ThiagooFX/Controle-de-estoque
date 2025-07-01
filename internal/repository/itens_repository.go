package repository

import (
	"api/internal/models"
	"database/sql"
)

func GetAllItems(db *sql.DB) ([]models.Item, error) {
	rows, err := db.Query("SELECT id, name, quantity FROM itens")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var itens []models.Item
	for rows.Next() {
		var item models.Item
		if err := rows.Scan(&item.ID, &item.Name, &item.Quantity); err != nil {
			return nil, err
		}
		itens = append(itens, item)
	}
	return itens, nil
}

func AddItem(db *sql.DB, item *models.Item) error {
	result, err := db.Exec("INSERT INTO itens (name, quantity) VALUES (?, ?)", item.Name, item.Quantity)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	item.ID = int(id)
	return nil
}

func UpdateItem(db *sql.DB, item *models.Item) error {
	_, err := db.Exec("UPDATE itens SET name=?, quantity=? WHERE id=?", item.Name, item.Quantity, item.ID)
	return err
}

func DeleteItem(db *sql.DB, id int) error {
	_, err := db.Exec("DELETE FROM itens WHERE id=?", id)
	return err
}
