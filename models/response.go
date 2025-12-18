package user_models

import "strconv"

// HATEOASLink 代表 HATEOAS 連結
// @Description HATEOAS 連結結構
type HATEOASLink struct {
	Href   string `json:"href" example:"http://api.example.com/users/1"` // 連結目標 URL
	Rel    string `json:"rel" example:"self"`                            // 關係類型
	Method string `json:"method" example:"GET"`                          // HTTP 方法
	Title  string `json:"title,omitempty" example:"取得使用者資訊"`             // 連結描述
}

// UserResponse 為符合 HATEOAS 的使用者響應結構
// @Description 符合 HATEOAS 的使用者響應結構
type UserResponse struct {
	Data  User          `json:"data"`   // 使用者資料
	Links []HATEOASLink `json:"_links"` // HATEOAS 連結
}

// UsersCollectionResponse 為包含多個使用者的 HATEOAS 響應結構
// @Description 符合 HATEOAS 的多使用者響應結構
type UsersCollectionResponse struct {
	Data  []User        `json:"data"`                // 使用者資料陣列
	Links []HATEOASLink `json:"_links"`              // HATEOAS 連結
	Page  int           `json:"page" example:"1"`    // 頁碼
	Size  int           `json:"size" example:"10"`   // 每頁大小
	Total int           `json:"total" example:"100"` // 總資料數
}

// APIResponse 為通用 API 響應結構
// @Description API 通用響應結構
type APIResponse struct {
	Status  int           `json:"status" example:"200"`           // HTTP 狀態碼
	Message string        `json:"message" example:"操作成功"`         // 響應訊息
	Data    interface{}   `json:"data,omitempty"`                 // 響應資料 (可選)
	Links   []HATEOASLink `json:"_links,omitempty"`               // HATEOAS 連結 (可選)
	Error   string        `json:"error,omitempty" example:"錯誤訊息"` // 錯誤訊息 (可選)
}

// GenerateUserLinks 產生使用者的 HATEOAS 連結
// @Description 產生使用者的 HATEOAS 連結
func GenerateUserLinks(baseURL string, userID string) []HATEOASLink {
	return []HATEOASLink{
		{
			Href:   baseURL + "/users/" + userID,
			Rel:    "self",
			Method: "GET",
			Title:  "取得使用者資訊",
		},
		{
			Href:   baseURL + "/users/" + userID,
			Rel:    "update",
			Method: "PUT",
			Title:  "更新使用者資訊",
		},
		{
			Href:   baseURL + "/users/" + userID,
			Rel:    "delete",
			Method: "DELETE",
			Title:  "刪除使用者",
		},
	}
}

// GenerateUsersCollectionLinks 產生使用者集合的 HATEOAS 連結
// @Description 產生使用者集合的 HATEOAS 連結
func GenerateUsersCollectionLinks(baseURL string, page int, size int, total int) []HATEOASLink {
	links := []HATEOASLink{
		{
			Href:   baseURL + "/users",
			Rel:    "self",
			Method: "GET",
			Title:  "取得使用者列表",
		},
		{
			Href:   baseURL + "/users",
			Rel:    "create",
			Method: "POST",
			Title:  "建立新使用者",
		},
	}

	// 添加分頁連結
	if page > 1 {
		links = append(links, HATEOASLink{
			Href:   baseURL + "/users?page=" + strconv.Itoa(page-1) + "&size=" + strconv.Itoa(size),
			Rel:    "prev",
			Method: "GET",
			Title:  "上一頁使用者",
		})
	}

	totalPages := (total + size - 1) / size
	if page < totalPages {
		links = append(links, HATEOASLink{
			Href:   baseURL + "/users?page=" + strconv.Itoa(page+1) + "&size=" + strconv.Itoa(size),
			Rel:    "next",
			Method: "GET",
			Title:  "下一頁使用者",
		})
	}

	return links
}
