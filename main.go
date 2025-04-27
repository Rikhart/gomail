package main

import (
	"context"
	"fmt"
	"gomail/docs"
	db "gomail/models"
	"gomail/queue"
	routes "gomail/routes"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("hola")

	// Conexión a la base de datos principal
	db.Connect()

	// Configuración de Swagger
	docs.SwaggerInfo.Title = "Swagger API"
	docs.SwaggerInfo.Description = "API de gomailo"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = ""
	docs.SwaggerInfo.BasePath = "/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	// Inicializar el router de Gin
	r := gin.Default()

	// Conexión al sistema de colas Backlite
	bl, mux, err := queue.Connect()
	if err != nil {
		log.Fatalf("Error al iniciar Backlite: %v", err)
	}

	// Iniciar procesamiento de tareas en segundo plano
	go bl.Start(context.Background())

	// Exponer panel de administración de Backlite
	r.GET("/panel/*any", gin.WrapH(http.StripPrefix("/panel", mux)))

	// Rutas principales de la aplicación
	routes.Launch(r)

	// Ejecutar el servidor en el puerto 8080
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}
