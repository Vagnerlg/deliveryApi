package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	provider "github.com/vagnerlg/deliveryApi/internal"
	httpApi "github.com/vagnerlg/deliveryApi/internal/http"
)

func main() {
	route := gin.Default()
	godotenv.Load()

	repo := provider.Repository()
	http := httpApi.New(repo)

	rProduct := route.Group("/product")
	{
		rProduct.POST("", http.Create)
		rProduct.GET("", http.List)
		rProduct.GET("/:id", http.Get)
	}

	route.Run(":8088")
}
