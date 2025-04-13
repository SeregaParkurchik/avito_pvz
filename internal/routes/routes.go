package routes

import (
	"avitopvz/internal/api"
	"avitopvz/internal/auth"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRoutes(userHandler *api.UserHandler) *gin.Engine {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router.POST("/register", userHandler.Register)
	router.POST("/login", userHandler.Login)

	router.POST("/pvz", auth.AuthMiddleware(auth.RoleModerator), userHandler.CreatePVZ)

	router.POST("/receptions", auth.AuthMiddleware(auth.RoleEmployee), userHandler.CreateReceptions)
	router.POST("/products", auth.AuthMiddleware(auth.RoleEmployee), userHandler.AddProduct)
	router.POST("/pvz/:pvzId/close_last_reception", auth.AuthMiddleware(auth.RoleEmployee), userHandler.CloseLastReception)
	router.POST("/pvz/:pvzId/delete_last_product", auth.AuthMiddleware(auth.RoleEmployee), userHandler.DeleteLastProduct)

	router.GET("/pvz", auth.AuthMiddleware(auth.RoleModerator, auth.RoleEmployee), userHandler.GetAllPVZ)

	return router
}
