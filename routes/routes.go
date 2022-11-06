package routes

import (
	"fmt"
	"github.com/NekruzRakhimov/unconvicted/docs"
	"github.com/NekruzRakhimov/unconvicted/pkg/controller"
	"github.com/NekruzRakhimov/unconvicted/utils"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"os"
)

func RunAllRoutes() {
	//gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	r.Use(controller.CORSMiddleware())

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
	r.GET("/", controller.PingPong)
	docs.SwaggerInfo.BasePath = "/"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.POST("/sign-up", controller.SignUp)
	r.POST("/sign-in", controller.SignIn)
	r.POST("/reference", controller.CreateReference)
	//r.POST("/image", controller.SaveImage)
	r.GET("/image", controller.GetImage)

	api := r.Group("/api", controller.UserIdentity)
	api.GET("/profile", controller.GetMe)
	api.GET("/reference", controller.GetMyReferences)
}
