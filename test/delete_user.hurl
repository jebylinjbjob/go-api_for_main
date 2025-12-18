# 用戶 API 測試 - 刪除用戶
# 測試 DELETE /api/v1/users/:id 端點

# 先創建一個用戶以便後續測試
POST http://localhost:8080/api/v1/users
Content-Type: application/json
{
  "name": "待刪除用戶",
  "email": "todelete@example.com",
  "password": "password123"
}

HTTP 201
[Captures]
delete_user_id: jsonpath "$.data.id"


# 測試刪除用戶（成功）
DELETE http://localhost:8080/api/v1/users/{{delete_user_id}}

HTTP 200
[Asserts]
jsonpath "$.message" exists


# 確認用戶已被刪除
GET http://localhost:8080/api/v1/users/{{delete_user_id}}

HTTP 404
[Asserts]
jsonpath "$.error" exists


# 測試刪除不存在的用戶
DELETE http://localhost:8080/api/v1/users/000000000000000000000000

HTTP 404
[Asserts]
jsonpath "$.error" exists


# 測試使用無效的用戶 ID
DELETE http://localhost:8080/api/v1/users/invalid-id

HTTP 400
[Asserts]
jsonpath "$.error" exists
