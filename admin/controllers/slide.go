package controllers

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"github.com/dsaouda/fiap-que-isso/admin/models"
	"net/http"
)

func GetAllSlides(c *gin.Context) {
	var results []models.Slide

	db := c.MustGet("db").(*mgo.Database)
	db.C("slide").Find(nil).All(&results)

	c.JSON(http.StatusOK, results)
}
