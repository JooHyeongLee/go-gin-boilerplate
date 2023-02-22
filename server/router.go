package server

import (
	"github.com/gin-gonic/gin"
	"go-gin-boilerplate/controllers"
	"go-gin-boilerplate/middlewares"
	"html/template"
	"log"
	"net/http"
	"time"
)

var html = template.Must(template.New("https").Parse(`
<html>
<head>
  <title>Https Test</title>
  <script src="/assets/app.js"></script>
</head>
<body>
  <h1 style="color:red;">Welcome, Ginner!</h1>
</body>
</html>
`))

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(middlewares.Logger())
	router.Use(gin.Recovery())

	health := new(controllers.HealthController)
	router.Static("/assets", "./assets")
	router.SetHTMLTemplate(html)

	router.GET("/", func(c *gin.Context) {
		if pusher := c.Writer.Pusher(); pusher != nil {
			if err := pusher.Push("/assets/app.js", nil); err != nil {
				log.Printf("Failed to push: %v", err)
			}
		}
		c.HTML(200, "https", gin.H{
			"status": "success",
		})
	})
	router.GET("/health", health.Status)
	router.GET("long_async", func(c *gin.Context) {
		cCp := c.Copy()
		go func() {
			time.Sleep(5 * time.Second)
			log.Println("Done! in path " + cCp.Request.URL.Path)
		}()
		c.String(http.StatusOK, "long_async!")
	})

	router.GET("/long_sync", func(c *gin.Context) {
		// simulate a long task with time.Sleep(). 5 seconds
		time.Sleep(5 * time.Second)

		// since we are NOT using a goroutine, we do not have to copy the context
		log.Println("Done! in path " + c.Request.URL.Path)
		c.String(http.StatusOK, "long_sync!")
	})

	v1 := router.Group("v1")
	{
		userGroup := v1.Group("member")
		{
			user := new(controllers.MemberController)
			userGroup.GET("/:id", user.Retrieve)
		}
	}
	return router

}
