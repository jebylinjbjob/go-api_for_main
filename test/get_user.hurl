# 用戶 API 測試 - 獲取特定用戶
# 測試 GET /api/v1/users/:id 端點

# 先創建一個用戶以便後續測試
POST http://localhost:8080/api/v1/users
Content-Type: application/json
{
  "name": "特定用戶",
  "email": "specific@example.com",
  "password": "password123"
}

HTTP 201
[Captures]
created_user_id: jsonpath "$.data.id"


# 測試獲取特定用戶（成功）
GET http://localhost:8080/api/v1/users/{{created_user_id}}

HTTP 200
[Asserts]
jsonpath "$.data.id" == "{{created_user_id}}"
jsonpath "$.data.name" == "特定用戶"
jsonpath "$.data.email" == "specific@example.com"
jsonpath "$.links" isCollection


# 測試獲取不存在的用戶
GET http://localhost:8080/api/v1/users/000000000000000000000000

HTTP 404
[Asserts]
jsonpath "$.error" exists


# 測試使用無效的用戶 ID
GET http://localhost:8080/api/v1/users/invalid-id

HTTP 400
[Asserts]
jsonpath "$.error" exists
