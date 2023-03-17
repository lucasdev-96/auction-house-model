package routes

import "github.com/gin-gonic/gin"

func Routes(routes *gin.Engine) {
	routes.NoRoute(func(c *gin.Context) {
		c.JSON(404, "Route not found, please try again")
	})
	v1 := routes.Group("/v1/api")
	{
		v1.GET("/", store)
	}
}

func store(c *gin.Context) {
	c.JSON(200, "testando")
}
