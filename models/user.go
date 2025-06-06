package user_models

import "go.mongodb.org/mongo-driver/bson/primitive"

// User 模型
// @Description 用戶模型
type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty" example:"507f1f77bcf86cd799439011"`
	Name     string             `bson:"name" json:"name" binding:"required" example:"張三"`
	Email    string             `bson:"email" json:"email" binding:"required,email" example:"zhangsan@example.com"`
	Password string             `bson:"password" json:"-" binding:"required"` // 密碼不會在 JSON 中返回
}

// ErrorResponse 錯誤響應結構
type ErrorResponse struct {
	Error string `json:"error" example:"error message"`
}

// SuccessResponse 成功響應結構
type SuccessResponse struct {
	Message string `json:"message" example:"operation successful"`
}
