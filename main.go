package main

import (
	"api/src/core"
	"api/src/musics/infrastructure_music"
	"api/src/products/infrastructure"

	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	core.ConnectToDataBase()

	musicRepo := infrastructuremusic.NewMySQLRepositoryMusics()
	productRepo := infrastructure.NewMySQLRepositoryProducts()

	router := gin.Default()

	infrastructure.SetupProductRoutes(router, productRepo)
	infrastructuremusic.SetupMusicRoutes(router, musicRepo)

	log.Println("Iniciando el Servidor en el puerto 8080...")

	if err := router.Run(":8080"); err!= nil {
		log.Fatal("Erro al iniciar el servidor", err)
    }

}
