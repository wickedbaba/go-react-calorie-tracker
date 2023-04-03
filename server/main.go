package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/cors"
	"github.com/gin-gonic/gin"
	"github.com/wickedbaba/go-react-calorie-tracker/server/routes"
)

var pl = fmt.Println
var pf = fmt.Printf

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(cors.Default())

	router.POST("/entry/create", routes.AddEntry)
	router.GET("/entries", routes.GetEntries)
	router.GET("/entry/:id/", routes.GetEntryById)
	router.GET("/ingredient/:ingredient", routes.GetEntriesByIngredient)

	router.PUT("/entry/update/:id", routes.UpdateEntry)
	router.PUT("/ingredient/update/:id", routes.UpdateIngredient)
	router.DELETE("/entry/delete/:id", routes.DeleteEntry)
	router.Run(":" + port)

}
