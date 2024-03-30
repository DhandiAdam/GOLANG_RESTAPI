package main

import (
	"fmt"

	"github.com/DhandiAdam/GOLANG_RESTAPI/controllers/productcontroller"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/api/products", productcontroller.Index)
	r.GET("/api/product/:id", productcontroller.Show)
	r.POST("/api/product", productcontroller.Create)
	r.PUT("/api/product/:id", productcontroller.Update)
	r.DELETE("/api/product", productcontroller.Delete)

	address := ":8080"
	go func() {
		if err := r.Run(address); err != nil {
			panic(fmt.Sprintf("Failed to run server: %v", err))
		}
	}()
	fmt.Printf("Server is running on %s\n", address)
	select {}
}
