package controllers

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"net/http"
)

func GetAllGroups(c *gin.Context) {
	var result []string

	db := c.MustGet("db").(*mgo.Database)
	collection := db.C("slide")

	err := collection.Find(nil).Distinct("group", &result)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "nao foi possivel obter os grupos"})
		return
	}

	c.JSON(http.StatusOK, result)
}
