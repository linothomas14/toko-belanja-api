package main

import (
	"os"

	"github.com/gin-gonic/gin"
	_handler "github.com/linothomas14/toko-belanja-api/app/delivery"
	_repository "github.com/linothomas14/toko-belanja-api/app/repository"
	_usecase "github.com/linothomas14/toko-belanja-api/app/usecase"
	"github.com/linothomas14/toko-belanja-api/config"
)

func main() {
	router := gin.Default()
	config.StartDB()
	db := config.GetDBConnection()

	userRepository := _repository.NewUserRepository(db)
	userUsecase := _usecase.NewUserUsecase(userRepository)

	categoryRepository := _repository.NewCategoryRepository(db)
	categoryUsecase := _usecase.NewCategoryUsecase(categoryRepository)

	productRepository := _repository.NewProductRepository(db)
	productUsecase := _usecase.NewProductUsecase(productRepository, categoryRepository)

	api := router.Group("/")
	_handler.NewUserHandler(api, userUsecase)
	_handler.NewCategoryHandler(api, categoryUsecase)
	_handler.NewProductHandler(api, productUsecase)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	router.Run(":" + port)
}
