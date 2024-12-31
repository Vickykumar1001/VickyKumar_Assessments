package main

import (
	db "A2_EMS_PROJECT/config"
	"A2_EMS_PROJECT/controller"
	"A2_EMS_PROJECT/middlewares"
	"A2_EMS_PROJECT/repository"
	"A2_EMS_PROJECT/service"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitializeDatabase()
	userRepo := repository.NewUserRepository(db.GetDB())
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)
	productRepo := repository.NewProductRepository(db.GetDB())
	productService := service.NewProductService(productRepo)
	productController := controller.NewProductController(productService)

	r := gin.Default()
	r.Use(middlewares.Logger())
	r.Use(middlewares.RateLimiter(5, 10))

	r.POST("/signup", userController.Signup)
	r.POST("/login", userController.Login)
	protected := r.Group("/api")
	protected.Use(middlewares.JWTAuth())
	protected.POST("/products", productController.CreateProduct)
	protected.GET("/products/:id", productController.GetProduct)
	protected.GET("/products", productController.GetAllProducts)
	protected.PUT("/products/:id", productController.UpdateProduct)
	protected.DELETE("/products/:id", productController.DeleteProduct)

	r.Run(":8085")

}
