package controllers

import (
	"context"
	"errors"
	"net/http"

	user_models "go-api_for_main/models/user_models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection
var ErrMongoDBNotConnected = errors.New("MongoDB is not connected")

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
		c.JSON(http.StatusServiceUnavailable, user_models.ErrorResponse{Error: "Database service is currently unavailable"})
		return
	}

	var users []user_models.User
	cursor, err := userCollection.Find(context.Background(), bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, user_models.ErrorResponse{Error: err.Error()})
		return
	}
	defer cursor.Close(context.Background())

	if err = cursor.All(context.Background(), &users); err != nil {
		c.JSON(http.StatusInternalServerError, user_models.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

func GetUsers_test(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"users_name":    "test users",
		"user_id":       "1234567890",
		"user_email":    "test@example.com",
		"user_password": "1234567890",
	})
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
		c.JSON(http.StatusServiceUnavailable, user_models.ErrorResponse{Error: "Database service is currently unavailable"})
		return
	}

	var user user_models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, user_models.ErrorResponse{Error: err.Error()})
		return
	}

	result, err := userCollection.InsertOne(context.Background(), user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, user_models.ErrorResponse{Error: err.Error()})
		return
	}

	user.ID = result.InsertedID.(primitive.ObjectID)
	c.JSON(http.StatusCreated, user)
}

func CreateUser_test(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"users_name":    "test users",
		"user_id":       "1234567890",
		"user_email":    "test@example.com",
		"user_password": "1234567890",
		"status":        "success",
	})
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
		c.JSON(http.StatusServiceUnavailable, user_models.ErrorResponse{Error: "Database service is currently unavailable"})
		return
	}

	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, user_models.ErrorResponse{Error: "Invalid ID"})
		return
	}

	var user user_models.User
	err = userCollection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, user_models.ErrorResponse{Error: "User not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, user_models.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

func GetUser_test(c *gin.Context) {

	user_id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"users_name":    "test users",
		"user_id":       user_id,
		"user_email":    "test@example.com",
		"user_password": "1234567890",
		"status":        "success",
		"is_find":       true,
	})
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
		c.JSON(http.StatusServiceUnavailable, user_models.ErrorResponse{Error: "Database service is currently unavailable"})
		return
	}

	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, user_models.ErrorResponse{Error: "Invalid ID"})
		return
	}

	var user user_models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, user_models.ErrorResponse{Error: err.Error()})
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
		c.JSON(http.StatusInternalServerError, user_models.ErrorResponse{Error: err.Error()})
		return
	}

	if result.MatchedCount == 0 {
		c.JSON(http.StatusNotFound, user_models.ErrorResponse{Error: "User not found"})
		return
	}

	c.JSON(http.StatusOK, user_models.SuccessResponse{Message: "User updated successfully"})
}

func UpdateUser_test(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"users_name":    "test users",
		"user_id":       "1234567890",
		"user_email":    "test@example.com",
		"user_password": "1234567890",
		"status":        "success",
		"is_update":     true,
	})
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
		c.JSON(http.StatusServiceUnavailable, user_models.ErrorResponse{Error: "Database service is currently unavailable"})
		return
	}

	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, user_models.ErrorResponse{Error: "Invalid ID"})
		return
	}

	result, err := userCollection.DeleteOne(context.Background(), bson.M{"_id": id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, user_models.ErrorResponse{Error: err.Error()})
		return
	}

	if result.DeletedCount == 0 {
		c.JSON(http.StatusNotFound, user_models.ErrorResponse{Error: "User not found"})
		return
	}

	c.JSON(http.StatusOK, user_models.SuccessResponse{Message: "User deleted successfully"})
}

func DeleteUser_test(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"users_name":    "test users",
		"user_id":       "1234567890",
		"user_email":    "test@example.com",
		"user_password": "1234567890",
		"status":        "success",
		"is_delete":     true,
	})
}
