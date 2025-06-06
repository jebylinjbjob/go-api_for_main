package routes

import (
	"go-api_for_main/controllers"

	"github.com/gin-gonic/gin"
)

// SetupRouter 初始化所有路由
func SetupRouter(r *gin.Engine) {
	// API v1 路由組
	v1 := r.Group("/api/v1")
	{
		// 用戶相關路由
		users := v1.Group("/users")
		{
			users.GET("/", controllers.GetUsers)         // 獲取所有用戶
			users.POST("/", controllers.CreateUser)      // 創建用戶
			users.GET("/:id", controllers.GetUser)       // 獲取特定用戶
			users.PUT("/:id", controllers.UpdateUser)    // 更新用戶
			users.DELETE("/:id", controllers.DeleteUser) // 刪除用戶

		}

		test := v1.Group("/test")
		{
			test.GET("/", func(c *gin.Context) {
				c.JSON(200, gin.H{
					"message": "is alive",
				})
			})
		}

		// 可以添加更多路由組
		// 例如：產品、訂單等
	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "is alive",
		})
	})

}
