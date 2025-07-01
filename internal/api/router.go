package api

import (
	"api/internal/services"
	"log"

	"github.com/labstack/echo"
)

func InitRouters() {

	e := echo.New()
	e.GET("/itens", services.GetItens)
	e.POST("/itens", services.AddItem)
	e.PUT("/itens/:id", services.UpdateItem)
	e.DELETE("/itens/:id", services.DeleteItem)

	log.Println("Servidor rodando na porta 2000")
	e.Start(":2000")
}
