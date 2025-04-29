package main

import (
	"fmt"
	"gomail/docs"
	db "gomail/models"
	"gomail/queue"
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

	// bl, mux, err := queue.Connect()
	// if err != nil {
	// 	log.Fatalf("Error al iniciar Backlite: %v", err)
	// }
	// // Iniciar procesamiento de tareas en segundo plano
	// go bl.Start(context.Background())
	// // Exponer panel de administración de Backlite
	// r.GET("/panel/*any", gin.WrapH(http.StripPrefix("/panel", mux)))
	// r.GET("/", func(c *gin.Context) {
	// 	c.Redirect(http.StatusFound, "/panel")
	// })
	// r.GET("/succeeded", func(c *gin.Context) {
	// 	c.Redirect(http.StatusFound, "/panel/succeeded")
	// })
	// r.GET("/failed", func(c *gin.Context) {
	// 	c.Redirect(http.StatusFound, "/panel/failed")
	// })
	// r.GET("/upcoming", func(c *gin.Context) {
	// 	c.Redirect(http.StatusFound, "/panel/upcoming")
	// })

	// Conexión al sistema de colas gogite
	// queue :=
	queue.Connect2()

	// Rutas principales de la aplicación
	routes.Launch(r)
	// Ejecutar el servidor en el puerto 8080
	if err := r.Run(":8081"); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}
