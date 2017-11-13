package main

import (
	"log"
	"os"
	"strconv"

	"github.com/SetyaK/BL-Onboarding3-Go-package"
	"github.com/SetyaK/BL-Onboarding3-Go-package/database"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/subosito/gotenv"
)

func main() {
	gotenv.Load()
	os.Setenv("DATABASE_ADAPTER", "mysql")
	r := gin.Default()

	// Initialize database session
	sess, err := database.NewSession()
	if err != nil {
		log.Fatal(err)
	}

	// Initialize repository
	pr := ministore.ProductRepository{Session: sess}

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	// Get all products
	r.GET("/product", func(c *gin.Context) {
		products, _, err := pr.GetAll()
		if err != nil {
			c.JSON(500, err)
		}
		c.JSON(200, products)
	})

	// Get product by id
	r.GET("/product/:id", func(c *gin.Context) {
		productID, err := strconv.ParseInt(c.Params.ByName("id"), 10, 64)
		if err != nil {
			c.JSON(400, err)
		} else {
			product, err := pr.GetByID(productID)
			if err != nil {
				c.JSON(500, err)
			}
			c.JSON(200, product)
		}
	})

	r.POST("/product", func(c *gin.Context) {
		productName := c.PostForm("name")
		productDescription := c.PostForm("description")
		productStock, err := strconv.Atoi(c.PostForm("initial_stock"))
		if err != nil {
			c.JSON(500, err)
			return
		}
		productID, err := pr.Add(productName, productDescription, productStock)
		if err != nil {
			c.JSON(500, err)
			return
		}

		product, err := pr.GetByID(productID)
		if err != nil {
			c.JSON(500, err)
			return
		}
		c.JSON(200, product)
	})

	r.POST("/product/:id", func(c *gin.Context) {
		productID, err := strconv.ParseInt(c.Params.ByName("id"), 10, 64)
		if err != nil {
			c.JSON(400, err)
			return
		}
		product, err := pr.GetByID(productID)
		if err != nil {
			c.JSON(500, err)
			return
		} else if product.ProductID < 1 {
			c.JSON(400, "Product ID does not exist")
			return
		}
		product.Name = c.PostForm("name")
		product.Description = c.PostForm("description")

		result, err := pr.Update(&product)
		if err != nil {
			c.JSON(500, err)
			return
		}

		c.JSON(200, gin.H{"Updated": result})
	})

	// Delete product by id
	r.DELETE("/product/:id", func(c *gin.Context) {
		productID, err := strconv.ParseInt(c.Params.ByName("id"), 10, 64)
		if err != nil {
			c.JSON(400, err)
			return
		}

		result, err := pr.Delete(productID)
		if err != nil {
			c.JSON(400, err)
			return
		}
		c.JSON(200, gin.H{"Deleted": result})
	})

	// Listen and Server in 0.0.0.0:4567
	r.Run(":4567")
}
