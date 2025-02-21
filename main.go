package main

import (
	"api/src/core"
	"api/src/musics/dependenciesm"
	"api/src/musics/infrastructurem"
	"api/src/musics/infrastructurem/routesm"
	"api/src/products/dependencies"
	"api/src/products/infrastructure"
	"api/src/products/infrastructure/routes"

	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	core.ConnectToDataBase()

	musicRepo := infrastructurem.NewMySQLRepositoryMusics()
	productRepo := infrastructure.NewMySQLRepositoryProducts()

	productDeps := dependencies.NewProductDependencies(productRepo)
	musicDeps := dependenciesm.NewDependenciesMusic(musicRepo)

	router := gin.Default()
	router.Use(cors.Default())
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:4200"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowCredentials: true,
	}))

	routes.SetupProductRoutes(router, productDeps)
	routesm.SetupMusicRoutes(router, musicDeps)

	log.Println("Iniciando el Servidor en el puerto 8080...")

	if err := router.Run(":8080"); err != nil {
		log.Fatal("Error al iniciar el servidor", err)
	}
}
