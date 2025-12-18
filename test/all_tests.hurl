# 用戶 API 完整測試套件
# 執行所有用戶 API 端點的測試

# ==========================================
# 1. 測試獲取所有用戶（初始狀態）
# ==========================================
GET http://api:8080/api/v1/users

HTTP 200
[Asserts]
jsonpath "$.data" isCollection
jsonpath "$.links" isCollection


# ==========================================
# 2. 測試創建用戶
# ==========================================
POST http://api:8080/api/v1/users
Content-Type: application/json
{
  "name": "張三",
  "email": "zhangsan@example.com",
  "password": "password123"
}

HTTP 201
[Captures]
user1_id: jsonpath "$.data.id"

[Asserts]
jsonpath "$.data.name" == "張三"
jsonpath "$.data.email" == "zhangsan@example.com"


# 創建第二個用戶
POST http://api:8080/api/v1/users
Content-Type: application/json
{
  "name": "李四",
  "email": "lisi@example.com",
  "password": "password456"
}

HTTP 201
[Captures]
user2_id: jsonpath "$.data.id"

[Asserts]
jsonpath "$.data.name" == "李四"


# ==========================================
# 3. 測試獲取所有用戶（應該有數據）
# ==========================================
GET http://api:8080/api/v1/users

HTTP 200
[Asserts]
jsonpath "$.data" count >= 2


# ==========================================
# 4. 測試獲取特定用戶
# ==========================================
GET http://api:8080/api/v1/users/{{user1_id}}

HTTP 200
[Asserts]
jsonpath "$.data.id" == "{{user1_id}}"
jsonpath "$.data.name" == "張三"


# ==========================================
# 5. 測試更新用戶
# ==========================================
PUT http://api:8080/api/v1/users/{{user1_id}}
Content-Type: application/json
{
  "name": "張三豐",
  "email": "zhangsanfeng@example.com"
}

HTTP 200
[Asserts]
jsonpath "$.data.name" == "張三豐"
jsonpath "$.data.email" == "zhangsanfeng@example.com"


# ==========================================
# 6. 驗證更新是否成功
# ==========================================
GET http://api:8080/api/v1/users/{{user1_id}}

HTTP 200
[Asserts]
jsonpath "$.data.name" == "張三豐"


# ==========================================
# 7. 測試刪除用戶
# ==========================================
DELETE http://api:8080/api/v1/users/{{user2_id}}

HTTP 200
[Asserts]
jsonpath "$.message" exists


# ==========================================
# 8. 驗證刪除是否成功
# ==========================================
GET http://api:8080/api/v1/users/{{user2_id}}

HTTP 404


# ==========================================
# 9. 清理測試數據
# ==========================================
DELETE http://api:8080/api/v1/users/{{user1_id}}

HTTP 200
