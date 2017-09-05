package controllers

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"github.com/dsaouda/fiap-que-isso/admin/models"
	"net/http"
	"gopkg.in/mgo.v2/bson"
)

func DeleteSlide(c *gin.Context) {
	id := c.Param("id")
	db := c.MustGet("db").(*mgo.Database)
	db.C("slide").RemoveId(bson.ObjectIdHex(id))

	c.JSON(http.StatusOK, "ok")
}

func GetAllSlides(c *gin.Context) {
	var results []models.Slide

	group := c.Query("q")

	db := c.MustGet("db").(*mgo.Database)
	collection := db.C("slide")

	if group != "" {
		collection.Find(bson.M{"group": group}).All(&results)
	} else {
		collection.Find(nil).Sort("group", "name", "value").All(&results)
	}

	c.JSON(http.StatusOK, results)
}
