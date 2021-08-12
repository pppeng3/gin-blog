package router

import (
	"Blog/handler"
	"Blog/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CollectRoutes(r *gin.Engine) {
	r.LoadHTMLGlob("temp/*")

	r.POST("/api/user/register", handler.Register)
	r.POST("/api/user/login", handler.Login)

	mustLogin := r.Group("/api", middleware.JWTAuthMiddleware)
	{
		//blog
		mustLogin.POST("/blog/create", handler.CreateBlog)
		mustLogin.POST("/blog/update", handler.UpdateBlog)
		mustLogin.DELETE("/blog/delete", handler.DeleteBlog)
		//comment
		mustLogin.POST("comment/create", handler.CreateComment)
		mustLogin.POST("comment/update", handler.UpdateComment)
		mustLogin.DELETE("comment/delete", handler.DeleteComment)
	}

	r.GET("api/blog/get", handler.GetBlog)
	r.GET("api/comment/get", handler.GetComment)

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "pp",
		})
	})

}
