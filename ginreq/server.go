package ginreq

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	HOST string = "localhost:8888"
	CERT string = ""
	KEY	 string = ""
)

func StartServer() {
	r := gin.Default()

	groupv1 := r.Group("/v1")
	setRoute(groupv1)

	groupt2 := r.Group("/test")
	testRoute(groupt2)

	//r.RunTLS(HOST, CERT, KEY)
	r.Run(HOST)
}

func setRoute(router *gin.RouterGroup) {

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "success",
			"data": "pong",
			"code": 0,
		})

		return
	})

	router.POST("/say", func(c *gin.Context) {
		param := &struct {
			name 	string
			age 	int
		}{}

		err := c.BindJSON(param)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "failed", "data": "", "code": 1})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "success",
			"data": fmt.Sprintf("name: %s, age: %d", param.name, param.age),
			"code": 0,
		})

		return
	})
}

func testRoute(r *gin.RouterGroup) {

	r.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "success",
			"data": "testing",
			"code": 0,
		})
		return
	})
}

