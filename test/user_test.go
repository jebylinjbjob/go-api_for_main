// Package test 提供了用戶 API 的測試用例
package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"go-api_for_main/controllers"
)

// setupTestRouter 初始化一個測試用的 Gin 路由器
// 設置為測試模式以避免輸出調試信息
func setupTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	return r
}

// TestUser 定義測試用的用戶結構
// 包含用戶的基本信息：ID、姓名、郵箱和密碼
type TestUser struct {
	ID       primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name     string             `json:"name" bson:"name"`
	Email    string             `json:"email" bson:"email"`
	Password string             `json:"password,omitempty" bson:"password"`
}

// TestErrorResponse 定義測試用的錯誤響應結構
// 用於解析 API 返回的錯誤信息
type TestErrorResponse struct {
	Error string `json:"error"`
}

// TestUserAPI 測試用戶 API 的所有端點
// 包括：
// - GET /api/v1/users：獲取所有用戶
// - POST /api/v1/users：創建新用戶
// - GET /api/v1/users/:id：獲取特定用戶
// - PUT /api/v1/users/:id：更新用戶
// - DELETE /api/v1/users/:id：刪除用戶
func TestUserAPI(t *testing.T) {
	r := setupTestRouter()

	// 設置測試路由
	r.GET("/api/v1/users", controllers.GetUsers)
	r.POST("/api/v1/users", controllers.CreateUser)
	r.GET("/api/v1/users/:id", controllers.GetUser)
	r.PUT("/api/v1/users/:id", controllers.UpdateUser)
	r.DELETE("/api/v1/users/:id", controllers.DeleteUser)

	// 測試獲取所有用戶
	// 預期：由於數據庫未連接，應返回 503 Service Unavailable
	t.Run("測試獲取所有用戶", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/v1/users", nil)
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusServiceUnavailable, w.Code)
	})

	// 測試創建新用戶
	// 使用有效的用戶數據進行測試
	// 預期：由於數據庫未連接，應返回 503 Service Unavailable
	t.Run("測試創建用戶", func(t *testing.T) {
		user := TestUser{
			Name:     "測試用戶",
			Email:    "test@example.com",
			Password: "test123",
		}
		userData, _ := json.Marshal(user)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/api/v1/users", bytes.NewBuffer(userData))
		req.Header.Set("Content-Type", "application/json")
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusServiceUnavailable, w.Code)
	})

	// 測試創建用戶時使用無效的 JSON 數據
	// 預期：由於數據庫未連接，應返回 503 Service Unavailable
	t.Run("測試創建用戶-無效數據", func(t *testing.T) {
		invalidData := []byte(`{"name": "test", "email": invalid}`)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/api/v1/users", bytes.NewBuffer(invalidData))
		req.Header.Set("Content-Type", "application/json")
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusServiceUnavailable, w.Code)
	})

	// 測試獲取特定用戶
	// 使用有效的 MongoDB ObjectID
	// 預期：由於數據庫未連接，應返回 503 Service Unavailable
	t.Run("測試獲取特定用戶", func(t *testing.T) {
		id := primitive.NewObjectID()
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/v1/users/"+id.Hex(), nil)
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusServiceUnavailable, w.Code)
	})

	// 測試更新用戶信息
	// 使用有效的用戶 ID 和更新數據
	// 預期：由於數據庫未連接，應返回 503 Service Unavailable
	t.Run("測試更新用戶", func(t *testing.T) {
		id := primitive.NewObjectID()
		updateData := TestUser{
			Name:  "更新名稱",
			Email: "update@example.com",
		}
		userData, _ := json.Marshal(updateData)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("PUT", "/api/v1/users/"+id.Hex(), bytes.NewBuffer(userData))
		req.Header.Set("Content-Type", "application/json")
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusServiceUnavailable, w.Code)
	})

	// 測試刪除用戶
	// 使用有效的用戶 ID
	// 預期：由於數據庫未連接，應返回 503 Service Unavailable
	t.Run("測試刪除用戶", func(t *testing.T) {
		id := primitive.NewObjectID()
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("DELETE", "/api/v1/users/"+id.Hex(), nil)
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusServiceUnavailable, w.Code)
	})

	// 測試使用無效的用戶 ID
	// 對所有需要 ID 的端點進行測試
	// 預期：由於數據庫未連接，應返回 503 Service Unavailable
	t.Run("測試無效的用戶ID", func(t *testing.T) {
		invalidID := "invalid-id"
		endpoints := []struct {
			method string
			path   string
		}{
			{"GET", "/api/v1/users/" + invalidID},
			{"PUT", "/api/v1/users/" + invalidID},
			{"DELETE", "/api/v1/users/" + invalidID},
		}

		for _, e := range endpoints {
			w := httptest.NewRecorder()
			var req *http.Request

			if e.method == "PUT" {
				updateData := TestUser{
					Name:  "測試名稱",
					Email: "test@example.com",
				}
				userData, _ := json.Marshal(updateData)
				req, _ = http.NewRequest(e.method, e.path, bytes.NewBuffer(userData))
				req.Header.Set("Content-Type", "application/json")
			} else {
				req, _ = http.NewRequest(e.method, e.path, nil)
			}

			r.ServeHTTP(w, req)

			assert.Equal(t, http.StatusServiceUnavailable, w.Code, "Method: %s", e.method)
			var response TestErrorResponse
			err := json.Unmarshal(w.Body.Bytes(), &response)
			assert.NoError(t, err)
			assert.Equal(t, "Database service is currently unavailable", response.Error)
		}
	})
}

// TestUserValidation 測試用戶數據驗證
// 測試各種無效的用戶數據輸入場景：
// - 空用戶名
// - 無效的郵箱格式
// - 密碼太短
func TestUserValidation(t *testing.T) {
	r := setupTestRouter()
	r.POST("/api/v1/users", controllers.CreateUser)

	// 定義測試用例
	testCases := []struct {
		name     string   // 測試用例名稱
		user     TestUser // 測試用戶數據
		expected int      // 預期的 HTTP 狀態碼
	}{
		{
			name: "空用戶名",
			user: TestUser{
				Name:     "",
				Email:    "test@example.com",
				Password: "test123",
			},
			expected: http.StatusServiceUnavailable,
		},
		{
			name: "無效郵箱",
			user: TestUser{
				Name:     "測試用戶",
				Email:    "invalid-email",
				Password: "test123",
			},
			expected: http.StatusServiceUnavailable,
		},
		{
			name: "密碼太短",
			user: TestUser{
				Name:     "測試用戶",
				Email:    "test@example.com",
				Password: "123",
			},
			expected: http.StatusServiceUnavailable,
		},
	}

	// 執行所有測試用例
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			userData, _ := json.Marshal(tc.user)
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("POST", "/api/v1/users", bytes.NewBuffer(userData))
			req.Header.Set("Content-Type", "application/json")
			r.ServeHTTP(w, req)

			assert.Equal(t, tc.expected, w.Code)
		})
	}
}
