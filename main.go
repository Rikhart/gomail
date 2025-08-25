package main

import (
	"context"
	"fmt"
	"gomail/docs"
	db "gomail/models"
	queue "gomail/queue"
	routes "gomail/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("inicia")
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

	// err = queue.EnqueueNewOrderEmail(bl, "12345", "cliente@correo.com")
	// // Iniciar procesamiento de tareas en segundo plano
	go bl.Start(context.Background())

	err = queue.EnqueueNewOrderEmail(bl, "12345", "cliente@correo.com")

	if err != nil {
		log.Printf("Error encolando tarea: %v", err)
	}

	// // Exponer panel de administración de Backlite
	r.Any("/dashboard", gin.WrapH(mux))
	r.Any("/dashboard/*any", gin.WrapH(mux))

	// Ejemplo: encolar tarea de nuevo pedido

	if err != nil {
		log.Printf("Error encolando tarea: %v", err)
	}

	// Rutas principales de la aplicación
	routes.Launch(r)
	// Ejecutar el servidor en el puerto 8080
	if err := r.Run(":8081"); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}
