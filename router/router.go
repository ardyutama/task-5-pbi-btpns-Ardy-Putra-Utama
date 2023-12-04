package router

import (
	"rakamin-golang/controllers"
	"rakamin-golang/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	v1 := r.Group("/api/v1")

	userController := controllers.NewUserController()
	photoController := controllers.NewPhotoController()

	// User Endpoints
	v1.POST("/users/register", userController.Register)
	v1.POST("/users/login", userController.Login)

	v1.Use(middlewares.AuthMiddleware()) // Middleware for authentication

	v1.PUT("/users/:userId", userController.UpdateUser)
	v1.DELETE("/users/:userId", userController.DeleteUser)

	// Photo Endpoints
	v1.POST("/photos", photoController.UploadPhoto)
	v1.GET("/photos", photoController.GetPhotos)
	v1.PUT("/photos/:photoId", photoController.UpdatePhoto)
	v1.DELETE("/photos/:photoId", photoController.DeletePhoto)
}
