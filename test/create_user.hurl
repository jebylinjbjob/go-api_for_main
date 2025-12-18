# 用戶 API 測試 - 創建用戶
# 測試 POST /api/v1/users 端點

# 測試成功創建用戶
POST http://localhost:8080/api/v1/users
Content-Type: application/json
{
  "name": "測試用戶",
  "email": "test@example.com",
  "password": "test12345"
}

HTTP 201
[Captures]
user_id: jsonpath "$.data.id"

[Asserts]
jsonpath "$.data.name" == "測試用戶"
jsonpath "$.data.email" == "test@example.com"
jsonpath "$.data.id" exists
jsonpath "$.links" isCollection


# 測試創建用戶 - 無效的 JSON
POST http://localhost:8080/api/v1/users
Content-Type: application/json
{"name": "test", "email": invalid}

HTTP 400
[Asserts]
jsonpath "$.error" exists


# 測試創建用戶 - 空用戶名
POST http://localhost:8080/api/v1/users
Content-Type: application/json
{
  "name": "",
  "email": "test@example.com",
  "password": "test123"
}

HTTP 400
[Asserts]
jsonpath "$.error" exists


# 測試創建用戶 - 無效郵箱
POST http://localhost:8080/api/v1/users
Content-Type: application/json
{
  "name": "測試用戶",
  "email": "invalid-email",
  "password": "test123"
}

HTTP 400
[Asserts]
jsonpath "$.error" exists


# 測試創建用戶 - 密碼太短
POST http://localhost:8080/api/v1/users
Content-Type: application/json
{
  "name": "測試用戶",
  "email": "test@example.com",
  "password": "123"
}

HTTP 400
[Asserts]
jsonpath "$.error" exists
