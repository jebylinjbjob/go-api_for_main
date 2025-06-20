package controllers

import (
	"context"
	"errors"
	"net/http"
	"time"

	user_models "go-api_for_main/models"

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
// @Success 200 {object} user_models.UsersCollectionResponse
// @Failure 500 {object} user_models.APIResponse
// @Router /users [get]
func GetUsers(c *gin.Context) {
	if err := checkMongoDBConnection(); err != nil {
		RespondWithAPIError(c, http.StatusServiceUnavailable, "Database service is currently unavailable")
		return
	}

	// 獲取分頁參數
	page := 1
	size := 10

	var users []user_models.User
	cursor, err := userCollection.Find(context.Background(), bson.M{})
	if err != nil {
		RespondWithAPIError(c, http.StatusInternalServerError, err.Error())
		return
	}
	defer cursor.Close(context.Background())

	if err = cursor.All(context.Background(), &users); err != nil {
		RespondWithAPIError(c, http.StatusInternalServerError, err.Error())
		return
	}

	// 計算總記錄數
	total := len(users)

	RespondWithUsersHATEOAS(c, http.StatusOK, users, page, size, total)
}

func GetUsers_test(c *gin.Context) {

	var users []user_models.User

	users = append(users, user_models.User{
		ID:        primitive.NewObjectID(),
		Name:      "test users",
		Email:     "test@example.com",
		Password:  "1234567890",
		Sex:       "男",
		Age:       20,
		Phone:     "1234567890",
		Address:   "台北市",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	RespondWithUsersHATEOAS(c, http.StatusOK, users, 1, 10, len(users))
}

// CreateUser godoc
// @Summary 創建新用戶
// @Description 創建一個新的用戶
// @Tags users
// @Accept json
// @Produce json
// @Param user body user_models.User true "用戶信息"
// @Success 201 {object} user_models.UserResponse
// @Failure 400 {object} user_models.APIResponse
// @Failure 500 {object} user_models.APIResponse
// @Router /users [post]
func CreateUser(c *gin.Context) {
	if err := checkMongoDBConnection(); err != nil {
		RespondWithAPIError(c, http.StatusServiceUnavailable, "Database service is currently unavailable")
		return
	}

	var user user_models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		RespondWithAPIError(c, http.StatusBadRequest, err.Error())
		return
	}

	result, err := userCollection.InsertOne(context.Background(), user)
	if err != nil {
		RespondWithAPIError(c, http.StatusInternalServerError, err.Error())
		return
	}

	user.ID = result.InsertedID.(primitive.ObjectID)
	RespondWithUserHATEOAS(c, http.StatusCreated, user)
}

func CreateUser_test(c *gin.Context) {

	var user user_models.User

	user = user_models.User{
		ID:        primitive.NewObjectID(),
		Name:      "test users",
		Email:     "test@example.com",
		Password:  "1234567890",
		Sex:       "男",
		Age:       20,
		Phone:     "1234567890",
		Address:   "台北市",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	RespondWithUserHATEOAS(c, http.StatusCreated, user)
}

// GetUser godoc
// @Summary 獲取特定用戶
// @Description 通過ID獲取特定用戶的信息
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "用戶ID"
// @Success 200 {object} user_models.UserResponse
// @Failure 400 {object} user_models.APIResponse
// @Failure 404 {object} user_models.APIResponse
// @Failure 500 {object} user_models.APIResponse
// @Router /users/{id} [get]
func GetUser(c *gin.Context) {
	if err := checkMongoDBConnection(); err != nil {
		RespondWithAPIError(c, http.StatusServiceUnavailable, "Database service is currently unavailable")
		return
	}

	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		RespondWithAPIError(c, http.StatusBadRequest, "Invalid ID")
		return
	}

	var user user_models.User
	err = userCollection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			RespondWithAPIError(c, http.StatusNotFound, "User not found")
			return
		}
		RespondWithAPIError(c, http.StatusInternalServerError, err.Error())
		return
	}

	RespondWithUserHATEOAS(c, http.StatusOK, user)
}

func GetUser_test(c *gin.Context) {
	id := c.Param("id")

	// 將字符串ID轉換為ObjectID
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		RespondWithAPIError(c, http.StatusBadRequest, "Invalid ID")
		return
	}
	user := user_models.User{
		ID:        objectID,
		Name:      "test users",
		Email:     "test@example.com",
		Password:  "1234567890",
		Sex:       "男",
		Age:       20,
		Phone:     "1234567890",
		Address:   "台北市",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	RespondWithUserHATEOAS(c, http.StatusOK, user)
}

// UpdateUser godoc
// @Summary 更新用戶
// @Description 更新特定用戶的信息
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "用戶ID"
// @Param user body user_models.User true "用戶信息"
// @Success 200 {object} user_models.APIResponse
// @Failure 400 {object} user_models.APIResponse
// @Failure 404 {object} user_models.APIResponse
// @Failure 500 {object} user_models.APIResponse
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
	id := c.Param("id")

	// 將字符串ID轉換為ObjectID
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, user_models.ErrorResponse{Error: "Invalid ID"})
		return
	}

	var updateUser user_models.User
	if err := c.ShouldBindJSON(&updateUser); err != nil {
		c.JSON(http.StatusBadRequest, user_models.ErrorResponse{Error: err.Error()})
		return
	}

	// 模擬更新後的用戶數據
	user := user_models.User{
		ID:        objectID,
		Name:      updateUser.Name,
		Email:     updateUser.Email,
		Password:  "1234567890",
		Sex:       "男",
		Age:       20,
		Phone:     "1234567890",
		Address:   "台北市",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	c.JSON(http.StatusOK, user)
}

// DeleteUser godoc
// @Summary 刪除用戶
// @Description 刪除特定用戶
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "用戶ID"
// @Success 200 {object} user_models.APIResponse
// @Failure 400 {object} user_models.APIResponse
// @Failure 404 {object} user_models.APIResponse
// @Failure 500 {object} user_models.APIResponse
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
	// 在測試環境中，直接返回刪除成功的訊息
	c.JSON(http.StatusOK, user_models.SuccessResponse{Message: "User deleted successfully"})
}
