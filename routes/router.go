package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	c"gogo/controllers"
)

func RouteRegister(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "☺︎ welcome to golang app!")
	})


	_v1 := r.Group("/v1")
	{
		_v1.GET("books", c.Index)
		_v1.GET("update", c.Update)
		_v1.GET("delete", c.Delete)
		_v1.GET("insert", c.Insert)
	}

}