package controllers

import (
	"fmt"
	user_models "go-api_for_main/models"

	"github.com/gin-gonic/gin"
)

// getBaseURL 獲取請求的基本 URL
func getBaseURL(c *gin.Context) string {
	scheme := "http"
	if c.Request.TLS != nil {
		scheme = "https"
	}
	return fmt.Sprintf("%s://%s", scheme, c.Request.Host)
}

// RespondWithUserHATEOAS 回傳單個使用者的 HATEOAS 響應
func RespondWithUserHATEOAS(c *gin.Context, statusCode int, user user_models.User) {
	baseURL := getBaseURL(c)

	response := user_models.UserResponse{
		Data:  user,
		Links: user_models.GenerateUserLinks(baseURL, user.ID.Hex()),
	}

	c.JSON(statusCode, response)
}

// RespondWithUsersHATEOAS 回傳多個使用者的 HATEOAS 響應
func RespondWithUsersHATEOAS(c *gin.Context, statusCode int, users []user_models.User, page int, size int, total int) {
	baseURL := getBaseURL(c)

	response := user_models.UsersCollectionResponse{
		Data:  users,
		Links: user_models.GenerateUsersCollectionLinks(baseURL, page, size, total),
		Page:  page,
		Size:  size,
		Total: total,
	}

	c.JSON(statusCode, response)
}

// RespondWithAPIError 回傳 API 錯誤響應
func RespondWithAPIError(c *gin.Context, statusCode int, errMessage string) {
	response := user_models.APIResponse{
		Status:  statusCode,
		Message: "操作失敗",
		Error:   errMessage,
	}

	c.JSON(statusCode, response)
}

// RespondWithAPISuccess 回傳 API 成功響應
func RespondWithAPISuccess(c *gin.Context, statusCode int, message string, data interface{}, links []user_models.HATEOASLink) {
	response := user_models.APIResponse{
		Status:  statusCode,
		Message: message,
		Data:    data,
		Links:   links,
	}

	c.JSON(statusCode, response)
}
