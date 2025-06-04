package main

import (
	"context"
	"log"
	"time"

	"go-api_for_main/controllers"
	_ "go-api_for_main/docs" // 導入 swagger 文檔
	"go-api_for_main/routes"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// @title Go API with Gin and MongoDB
// @version 1.0
// @description 這是一個使用 Gin 和 MongoDB 的 RESTful API 服務
// @host localhost:8080
// @BasePath /api/v1
// @schemes http
// @produce application/json
// @consume application/json

var client *mongo.Client
var database *mongo.Database

func initMongoDB() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	var err error
	client, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		return err
	}

	// 檢查連接
	err = client.Ping(ctx, nil)
	if err != nil {
		return err
	}

	database = client.Database("go_api_db")
	controllers.SetupUserController(database)

	log.Println("Connected to MongoDB!")
	return nil
}

func main() {
	// 初始化 MongoDB 連接
	err := initMongoDB()
	if err != nil {
		log.Printf("Warning: MongoDB connection failed: %v\n", err)
		log.Println("Starting server without MongoDB connection...")
	} else {
		defer func() {
			if err := client.Disconnect(context.Background()); err != nil {
				log.Printf("Error disconnecting from MongoDB: %v\n", err)
			}
		}()
	}

	// 創建 Gin 路由器
	r := gin.Default()

	// 設置路由
	routes.SetupRouter(r)

	// Swagger 文檔路由
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 添加一個簡單的健康檢查端點
	r.GET("/health", func(c *gin.Context) {
		status := "OK"
		dbStatus := "Connected"
		if client == nil {
			dbStatus = "Disconnected"
		}
		c.JSON(200, gin.H{
			"status":         status,
			"mongodb_status": dbStatus,
		})
	})

	// 啟動服務器
	log.Println("Server starting on :8080...")
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
