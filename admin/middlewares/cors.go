package middlewares

import "github.com/gin-gonic/gin"

func Cors(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "x-access-token, Content-Type, access-control-allow-origin, access-control-allow-headers, AnonymousToken")
}
