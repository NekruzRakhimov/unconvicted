package routes

import (
	"fmt"
	"github.com/NekruzRakhimov/unconvicted/docs"
	"github.com/NekruzRakhimov/unconvicted/models"
	"github.com/NekruzRakhimov/unconvicted/utils"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"os"
)

func RunAllRoutes() {
	//gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	// Статус код 500, при любых panic()
	r.Use(gin.Recovery())

	// Запуск роутов
	initAllRoutes(r)

	// Запуск сервера
	port, exists := os.LookupEnv("PORT")
	if !exists {
		port = utils.AppSettings.AppParams.PortRun
	}
	_ = r.Run(fmt.Sprintf(":%s", port))
}

func initAllRoutes(r *gin.Engine) {
	r.GET("/ping", PingPong)
	docs.SwaggerInfo.BasePath = "/"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

}

//PingPong unconvicted godoc
// @Summary unconvicted
// @Description Роут для проверки работы сервера
// @Accept  json
// @Produce  json
// @Tags url
// @Success 200 {object} models.PingPong
// @Failure 400,404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /ping [get]
func PingPong(c *gin.Context) {
	c.JSON(http.StatusOK, models.PingPong{Pong: "pong"})
}
