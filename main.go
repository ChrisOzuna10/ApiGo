package main

import (
	"api/src/core"
	"api/src/infrastructure"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	core.ConnectToDataBase()

	productRepo := infrastructure.NewMySQLRepository()

	router := gin.Default()

	infrastructure.SetupProductRoutes(router, productRepo)

	log.Println("Iniciando el Servidor en el puerto 8080...")

	if err := router.Run(":8080"); err!= nil {
		log.Fatal("Erro al iniciar el servidor", err)
    }

}
