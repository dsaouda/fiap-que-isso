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

	router.Use(middlewares.Mgo)

	router.POST("/login", cors, controllers.PostLogin)
	router.OPTIONS("/login", cors)
	router.Run("localhost:8089")

}

func cors(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "Content-Type, access-control-allow-origin, access-control-allow-headers, AnonymousToken")
}