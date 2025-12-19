# 測試 CORS 設定
# 在 test 目錄中建立 CORS 測試檔案

# 測試 CORS 預檢請求 (OPTIONS)
OPTIONS http://api:8080/api/v1/users
Origin: http://localhost:3000
Access-Control-Request-Method: POST
Access-Control-Request-Headers: Content-Type

HTTP 204
[Asserts]
header "Access-Control-Allow-Origin" exists
header "Access-Control-Allow-Methods" contains "POST"
header "Access-Control-Allow-Headers" exists


# 測試 CORS GET 請求
GET http://api:8080/api/v1/users
Origin: http://localhost:3000

HTTP 200
[Asserts]
header "Access-Control-Allow-Origin" exists
jsonpath "$.data" isCollection


# 測試 CORS POST 請求
POST http://api:8080/api/v1/users
Origin: http://localhost:3000
Content-Type: application/json
{
  "name": "CORS 測試用戶",
  "email": "cors@example.com",
  "password": "test123456"
}

HTTP 201
[Captures]
cors_user_id: jsonpath "$.data.id"

[Asserts]
header "Access-Control-Allow-Origin" exists
jsonpath "$.data.name" == "CORS 測試用戶"


# 清理測試數據
DELETE http://api:8080/api/v1/users/{{cors_user_id}}
Origin: http://localhost:3000

HTTP 200
[Asserts]
header "Access-Control-Allow-Origin" exists
