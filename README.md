# Go API with Gin and MongoDB

這是一個使用 Gin 框架和 MongoDB 的 RESTful API 專案。

## 前置需求

- Go 1.21 或更高版本
- MongoDB 服務器
- Git

## 安裝步驟

1. 克隆專案：
```bash
git clone [your-repository-url]
cd go-api_for_main
```

2. 安裝依賴：
```bash
go mod download
```

3. 確保 MongoDB 服務器正在運行（預設端口：27017）

4. 運行專案：
```bash
go run main.go
```

## API 端點

- `GET /ping`: 健康檢查端點

## 開發中的功能

- 用戶認證
- 數據 CRUD 操作
- 錯誤處理中間件
- 日誌記錄 