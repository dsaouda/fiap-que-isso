package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/dsaouda/fiap-que-isso/admin/db"
)

func Mgo(c *gin.Context) {
	s := db.Session.Clone()
	defer s.Close()

	c.Set("db", s.DB(db.Mongo.Database))
	c.Next()
}
