package controllers

import (
	"context"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection
var ErrMongoDBNotConnected = errors.New("MongoDB is not connected")

// ErrorResponse 錯誤響應結構
type ErrorResponse struct {
	Error string `json:"error" example:"error message"`
}

// SuccessResponse 成功響應結構
type SuccessResponse struct {
	Message string `json:"message" example:"operation successful"`
}

// User 模型
// @Description 用戶模型
type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty" example:"507f1f77bcf86cd799439011"`
	Name     string             `bson:"name" json:"name" binding:"required" example:"張三"`
	Email    string             `bson:"email" json:"email" binding:"required,email" example:"zhangsan@example.com"`
	Password string             `bson:"password" json:"-" binding:"required"` // 密碼不會在 JSON 中返回
}

// SetupUserController 初始化用戶控制器
func SetupUserController(db *mongo.Database) {
	if db != nil {
		userCollection = db.Collection("users")
	}
}

func checkMongoDBConnection() error {
	if userCollection == nil {
		return ErrMongoDBNotConnected
	}
	return nil
}

// GetUsers godoc
// @Summary 獲取所有用戶
// @Description 獲取系統中的所有用戶列表
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {array} User
// @Failure 500 {object} ErrorResponse
// @Router /users [get]
func GetUsers(c *gin.Context) {
	if err := checkMongoDBConnection(); err != nil {
		c.JSON(http.StatusServiceUnavailable, ErrorResponse{Error: "Database service is currently unavailable"})
		return
	}

	var users []User
	cursor, err := userCollection.Find(context.Background(), bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}
	defer cursor.Close(context.Background())

	if err = cursor.All(context.Background(), &users); err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

// CreateUser godoc
// @Summary 創建新用戶
// @Description 創建一個新的用戶
// @Tags users
// @Accept json
// @Produce json
// @Param user body User true "用戶信息"
// @Success 201 {object} User
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /users [post]
func CreateUser(c *gin.Context) {
	if err := checkMongoDBConnection(); err != nil {
		c.JSON(http.StatusServiceUnavailable, ErrorResponse{Error: "Database service is currently unavailable"})
		return
	}

	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	result, err := userCollection.InsertOne(context.Background(), user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	user.ID = result.InsertedID.(primitive.ObjectID)
	c.JSON(http.StatusCreated, user)
}

// GetUser godoc
// @Summary 獲取特定用戶
// @Description 通過ID獲取特定用戶的信息
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "用戶ID"
// @Success 200 {object} User
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /users/{id} [get]
func GetUser(c *gin.Context) {
	if err := checkMongoDBConnection(); err != nil {
		c.JSON(http.StatusServiceUnavailable, ErrorResponse{Error: "Database service is currently unavailable"})
		return
	}

	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid ID"})
		return
	}

	var user User
	err = userCollection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, ErrorResponse{Error: "User not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

// UpdateUser godoc
// @Summary 更新用戶
// @Description 更新特定用戶的信息
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "用戶ID"
// @Param user body User true "用戶信息"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /users/{id} [put]
func UpdateUser(c *gin.Context) {
	if err := checkMongoDBConnection(); err != nil {
		c.JSON(http.StatusServiceUnavailable, ErrorResponse{Error: "Database service is currently unavailable"})
		return
	}

	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid ID"})
		return
	}

	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	update := bson.M{
		"$set": bson.M{
			"name":  user.Name,
			"email": user.Email,
		},
	}

	result, err := userCollection.UpdateOne(context.Background(), bson.M{"_id": id}, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	if result.MatchedCount == 0 {
		c.JSON(http.StatusNotFound, ErrorResponse{Error: "User not found"})
		return
	}

	c.JSON(http.StatusOK, SuccessResponse{Message: "User updated successfully"})
}

// DeleteUser godoc
// @Summary 刪除用戶
// @Description 刪除特定用戶
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "用戶ID"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /users/{id} [delete]
func DeleteUser(c *gin.Context) {
	if err := checkMongoDBConnection(); err != nil {
		c.JSON(http.StatusServiceUnavailable, ErrorResponse{Error: "Database service is currently unavailable"})
		return
	}

	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid ID"})
		return
	}

	result, err := userCollection.DeleteOne(context.Background(), bson.M{"_id": id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	if result.DeletedCount == 0 {
		c.JSON(http.StatusNotFound, ErrorResponse{Error: "User not found"})
		return
	}

	c.JSON(http.StatusOK, SuccessResponse{Message: "User deleted successfully"})
}
