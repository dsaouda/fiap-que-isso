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
	db := c.MustGet("db").(*mgo.Database)
	collection := db.C("slide")

	group := c.Query("q")
	where := bson.M{}

	if group != "" {
		where = bson.M{"group": group}
	}

	var results []models.Slide
	collection.Find(where).Sort("group", "order").All(&results)


	c.JSON(http.StatusOK, results)
}
