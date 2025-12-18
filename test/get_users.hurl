# 用戶 API 測試 - 獲取所有用戶
# 測試 GET /api/v1/users 端點

GET http://localhost:8080/api/v1/users

HTTP 200
[Asserts]
jsonpath "$.data" isCollection
jsonpath "$.links" isCollection
