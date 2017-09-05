package middlewares

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"github.com/dsaouda/fiap-que-isso/admin/models"
	"gopkg.in/mgo.v2/bson"
	"net/http"
)

func Auth(c *gin.Context) {

	if c.Request.Method == http.MethodOptions {
		c.Next()
		return
	}

	token := c.GetHeader("X-Access-Token")

	db := c.MustGet("db").(*mgo.Database)
	result := models.Login{}
	db.C("login").Find(bson.M{"token": token}).One(&result)

	if (models.Login{}) == result {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Você não tem permissão para acessar esse recurso"})
		c.Abort()
		return
	}

	c.Next()
}
