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
func GetItens(c echo.Context) error {
	itens, err := repository.GetAllItems(repository.DB)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, itens)
}

func AddItem(c echo.Context) error {
	var item models.Item
	if err := c.Bind(&item); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if err := repository.AddItem(repository.DB, &item); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	go func() {
		if err := Addlogs(c.Request(), fmt.Sprintf("Adicionou item ID %d", item.ID)); err != nil {
			log.Printf("Erro ao registrar log: %v", err)
		}
	}()

	return c.JSON(http.StatusCreated, item)
}

func UpdateItem(c echo.Context) error {
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

	if err := repository.UpdateItem(repository.DB, &item); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, item)
}

func DeleteItem(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "ID inválido"})
	}

	if err := repository.DeleteItem(repository.DB, id); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.NoContent(http.StatusNoContent)
}
