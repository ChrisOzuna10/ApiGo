package main

import (
	"api/src/core"
	"api/src/musics/infrastructurem"
	"api/src/products/infrastructure"

	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	core.ConnectToDataBase()

	musicRepo := infrastructurem.NewMySQLRepositoryMusics()
	productRepo := infrastructure.NewMySQLRepositoryProducts()

	router := gin.Default()

	infrastructure.SetupProductRoutes(router, productRepo)
	infrastructurem.SetupMusicRoutes(router, musicRepo)

	log.Println("Iniciando el Servidor en el puerto 8080...")

	if err := router.Run(":8080"); err != nil {
		log.Fatal("Error al iniciar el servidor", err)
	}
}
