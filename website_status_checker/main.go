package main

import (
	"website_status_checker/controllers"
	"website_status_checker/database"

	"github.com/gin-gonic/gin"
)

//var r *gin.Engine

func main() {
	r := gin.Default()

	database.ConnectDataBase()
	go controllers.Checklink()

	r.GET("/urls/", controllers.GetUrls)
	r.GET("/urls/:id", controllers.GetUrl)
	r.POST("/urls", controllers.CreateUrl)
	r.PATCH("urls/:id", controllers.Updateurl)
	r.DELETE("urls/:id", controllers.Deleteurl)

	r.Run(":8080")
}
