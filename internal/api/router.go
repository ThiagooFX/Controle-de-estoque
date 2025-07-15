package api

import (
	"api/internal/repository"
	"api/internal/services"
	"log"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func InitRouters() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.RequestID())
	handler := services.NewItemHandler(repository.NewItemRepository(repository.DB))

	e.GET("/itens", handler.GetItens)
	e.POST("/itens", handler.AddItem)
	e.PUT("/itens/:id", handler.UpdateItem)
	e.DELETE("/itens/:id", handler.DeleteItem)

	log.Println("Servidor rodando na porta 2000")
	e.Logger.Fatal(e.Start(":2000"))
}
