package main

//guia
//https://github.com/Coderockr/APIVagaFrontend
//https://github.com/madhums/go-gin-mgo-demo/

import (
	"github.com/dsaouda/fiap-que-isso/admin/db"
	"github.com/gin-gonic/gin"
	"github.com/dsaouda/fiap-que-isso/admin/middlewares"
	"github.com/dsaouda/fiap-que-isso/admin/controllers"
)

func init()  {
	db.Connect()
}

func main() {
	router := gin.Default()

	router.Use(middlewares.Cors)
	router.Use(middlewares.Mgo)


	router.POST("/login", controllers.PostLogin)
	router.OPTIONS("/login")

	authorized := router.Group("/slides")
	authorized.Use(middlewares.Auth)
	{
		authorized.GET("", controllers.GetAllSlides)
		authorized.DELETE("/:id", controllers.DeleteSlide)
		authorized.OPTIONS("")
		authorized.OPTIONS("/:id")
	}



	router.Run("localhost:8089")

}

