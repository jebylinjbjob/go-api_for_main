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

		// 可以添加更多路由組
		// 例如：產品、訂單等
	}

	test := r.Group("/api/test")
	{
		users := test.Group("/users")
		{
			users.GET("/", controllers.GetUsers_test)         // 獲取所有用戶
			users.GET("/:id", controllers.GetUser_test)       // 獲取特定用戶
			users.POST("/", controllers.CreateUser_test)      // 創建用戶
			users.PUT("/:id", controllers.UpdateUser_test)    // 更新用戶
			users.DELETE("/:id", controllers.DeleteUser_test) // 刪除用戶

		}
	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "is alive",
		})
	})

}
