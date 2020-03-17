package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"web/handlers"
)

func Init() {
	// Creates a default gin router
	r := gin.Default() // Grouping routes
	// group : v1
	// v1 := r.Group("/v1")
	// {
	// 	v1.GET("/hello", handlers.HelloPage)
	// }

	v1 := r.Group("/v1")
	{
		v1.GET("/hello", handlers.HelloPage)
		v1.GET("/hello/:name", func(c *gin.Context) {
			name := c.Param("name")
			c.String(http.StatusOK, "Hello %s", name)
		})
		v1.GET("/welcome", func(c *gin.Context) {
			firstname := c.DefaultQuery("firstname", "Guest")
			lastname := c.Query("lastname")
			c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
		})
	}

	r.LoadHTMLGlob("templates/*")
	v2 := r.Group("/v2")
	{
		v2.GET("/index", func(c *gin.Context) {
			c.HTML(http.StatusOK, "index.html", gin.H{"title": "hello Gin."})
		})
	}
	v2.Use(ValidateToken())

	// 404 NotFound
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"status": 404,
			"error":  "404, page not exists!",
		})
	})

	// r.Use(PrintMiddleware(r))

	r.Run(":8000") // listen and serve on 0.0.0.0:8000
}

// PrintMiddleware is a function for test middleware
func PrintMiddleware(c *gin.Context) {
	fmt.Println("before request")
	c.Next()
}

func ValidateToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.FormValue("token")
		if token == "" {
			c.JSON(401, gin.H{
				"message": "Token required",
			})
			c.Abort()
			return
		}
		if token != "accesstoken" {
			c.JSON(http.StatusOK, gin.H{
				"message": "Invalid Token",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
