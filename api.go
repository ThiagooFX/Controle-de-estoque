package main

import (
	"api/internal/api"
	"api/internal/repository"

	_ "modernc.org/sqlite"
)

func main() {
	repository.InitDB()
	api.InitRouters()

}
