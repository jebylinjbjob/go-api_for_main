package user_models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User 模型
// @Description 用戶模型
type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty" example:"507f1f77bcf86cd799439011"`
	Name      string             `bson:"name" json:"name" binding:"required" example:"張三"`
	Email     string             `bson:"email" json:"email" binding:"required,email" example:"zhangsan@example.com"`
	Password  string             `bson:"password" json:"-" binding:"required"` // 密碼不會在 JSON 中返回
	Sex       string             `bson:"sex" json:"sex" binding:"required" example:"男"`
	Age       int                `bson:"age" json:"age" binding:"required" example:"20"`
	Phone     string             `bson:"phone" json:"phone" binding:"required" example:"1234567890"`
	Address   string             `bson:"address" json:"address" binding:"required" example:"台北市"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at" example:"2021-01-01T00:00:00Z"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at" example:"2021-01-01T00:00:00Z"`
}

// ErrorResponse 錯誤響應結構
type ErrorResponse struct {
	Error string `json:"error" example:"error message"`
}

// SuccessResponse 成功響應結構
type SuccessResponse struct {
	Message string `json:"message" example:"operation successful"`
}
