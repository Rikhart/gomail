package routes

import (
	"fmt"
	_ "gomail/docs"
	"net/http"

	models "gomail/models"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// obtenerSensores godoc
// @Summary Listar sensores
// @Description Retorna la lista de sensores disponibles
// @Tags sensores
// @Param message body models.ContactRequest true "Cuerpo del mensaje"
// @Produce json
// @Success 200 {array} string
// @Router /send-mail [post]
func obtenerSensores(c *gin.Context) {
	var body models.ContactRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido"})
		return
	}
	fmt.Println(body.Subject, "asunto")
	c.JSON(200, []string{"ultrasonico", "gas"})
}

func Launch(r *gin.Engine) {

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	group := r.Group("/v1")
	group.POST("/send-mail", obtenerSensores)

}
