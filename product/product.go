package product

import "github.com/gin-gonic/gin"

type Product struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Price int `json:"price"`
	Description string `json:"description"`
}

func Register(router *gin.RouterGroup) {
	productsRegister(router)
}

func productsRegister(router *gin.RouterGroup) {
	router.GET("/", productsHandler)
}

func productsHandler(c *gin.Context) {
	products := []Product {
		Product {100, "BassTune Headset 2.0", 200, "A headphone with a inbuilt high-quality microphone"},
		Product {101, "Fastlane Toy Car", 100, "A toy car that comes with a free HD camera"},
		Product {101, "ATV Gear Mouse", 75, "A high-quality mouse for office work and gaming"},
	}
	c.JSON(200, gin.H{
		"products": products,
	})
}

