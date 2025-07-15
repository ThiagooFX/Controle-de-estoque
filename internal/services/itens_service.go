package services

import (
	"api/internal/models"
	"api/internal/repository"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type ItemHandler struct{
	Repo repository.ItemRepository
}

func NewItemHandler(repo repository.ItemRepository) *ItemHandler{
	return &ItemHandler{Repo:repo}
}

func (r *ItemHandler) GetItens(c echo.Context) error {
	itens, err := r.Repo.GetAllItems()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, itens)
}

func (r *ItemHandler) AddItem(c echo.Context) error {
	var item models.Item
	if err := c.Bind(&item); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if err := r.Repo.AddItem(&item); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	go func() {
		if err := Addlogs(c.Request(), fmt.Sprintf("Adicionou item ID %d", item.ID)); err != nil {
			log.Printf("Erro ao registrar log: %v", err)
		}
	}()

	return c.JSON(http.StatusCreated, item)
}

func (r *ItemHandler) UpdateItem(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "ID inválido"})
	}

	var item models.Item
	if err := c.Bind(&item); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	item.ID = id

	if err := r.Repo.UpdateItem(&item); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, item)
}

func (r *ItemHandler) DeleteItem(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "ID inválido"})
	}

	if err := r.Repo.DeleteItem(id); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.NoContent(http.StatusNoContent)
}
