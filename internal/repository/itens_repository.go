package repository

import (
	"api/internal/models"
	"database/sql"
)

type ItemRepository interface {
	GetAllItems() ([]models.Item, error)
	AddItem(item *models.Item) error
	UpdateItem(item *models.Item) error
	DeleteItem(id int) error
}

type itemRepository struct {
	db *sql.DB
}

func NewItemRepository(db *sql.DB) ItemRepository {
	return &itemRepository{db: db}
}

func (r *itemRepository) GetAllItems() ([]models.Item, error) {
	rows, err := r.db.Query("SELECT id, name, quantity FROM itens")
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

func (r *itemRepository) AddItem(item *models.Item) error {
	result, err := r.db.Exec("INSERT INTO itens (name, quantity) VALUES (?, ?)", item.Name, item.Quantity)
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

func (r *itemRepository) UpdateItem(item *models.Item) error {
	_, err := r.db.Exec("UPDATE itens SET name=?, quantity=? WHERE id=?", item.Name, item.Quantity, item.ID)
	return err
}

func (r *itemRepository) DeleteItem(id int) error {
	_, err := r.db.Exec("DELETE FROM itens WHERE id=?", id)
	return err
}
