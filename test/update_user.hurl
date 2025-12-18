# 用戶 API 測試 - 更新用戶
# 測試 PUT /api/v1/users/:id 端點

# 先創建一個用戶以便後續測試
POST http://localhost:8080/api/v1/users
Content-Type: application/json
{
  "name": "原始用戶",
  "email": "original@example.com",
  "password": "password123"
}

HTTP 201
[Captures]
update_user_id: jsonpath "$.data.id"


# 測試更新用戶（成功）
PUT http://localhost:8080/api/v1/users/{{update_user_id}}
Content-Type: application/json
{
  "name": "更新名稱",
  "email": "updated@example.com"
}

HTTP 200
[Asserts]
jsonpath "$.data.id" == "{{update_user_id}}"
jsonpath "$.data.name" == "更新名稱"
jsonpath "$.data.email" == "updated@example.com"
jsonpath "$.links" isCollection


# 測試更新不存在的用戶
PUT http://localhost:8080/api/v1/users/000000000000000000000000
Content-Type: application/json
{
  "name": "更新名稱",
  "email": "update@example.com"
}

HTTP 404
[Asserts]
jsonpath "$.error" exists


# 測試使用無效的用戶 ID
PUT http://localhost:8080/api/v1/users/invalid-id
Content-Type: application/json
{
  "name": "測試名稱",
  "email": "test@example.com"
}

HTTP 400
[Asserts]
jsonpath "$.error" exists


# 測試更新用戶 - 無效數據
PUT http://localhost:8080/api/v1/users/{{update_user_id}}
Content-Type: application/json
{
  "name": "",
  "email": "invalid-email"
}

HTTP 400
[Asserts]
jsonpath "$.error" exists
