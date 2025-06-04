# Go API with Gin and MongoDB

這是一個使用 Gin 框架和 MongoDB 的 RESTful API 專案。提供用戶管理的基本功能，包括創建、讀取、更新和刪除（CRUD）操作。

## 功能特點

- RESTful API 設計
- MongoDB 數據存儲
- Swagger API 文檔
- 完整的錯誤處理
- 單元測試覆蓋
- 日誌記錄

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

### 用戶管理

- `GET /api/v1/users` - 獲取所有用戶
- `POST /api/v1/users` - 創建新用戶
- `GET /api/v1/users/:id` - 獲取特定用戶
- `PUT /api/v1/users/:id` - 更新用戶信息
- `DELETE /api/v1/users/:id` - 刪除用戶

### 系統

- `GET /ping` - 健康檢查端點
- `GET /swagger/*any` - Swagger API 文檔

## API 文檔

專案整合了 Swagger 文檔，運行服務器後訪問：
```
http://localhost:8080/swagger/index.html
```

## 測試

運行所有測試：
```bash
go test ./... -v
```

運行特定測試：
```bash
go test ./test -v
```

## 項目結構

```
.
├── controllers/     # API 控制器
├── models/         # 數據模型
├── routes/         # 路由配置
├── docs/          # Swagger 文檔
├── test/          # 測試文件
├── main.go        # 主程序入口
└── README.md      # 項目文檔
```

## 錯誤處理

API 使用標準的 HTTP 狀態碼：

- 200: 成功
- 400: 請求錯誤
- 404: 資源未找到
- 500: 服務器錯誤
- 503: 服務不可用（數據庫連接失敗等）

## 環境變量

- `MONGODB_URI`: MongoDB 連接字符串（默認：mongodb://localhost:27017）
- `PORT`: 服務器端口（默認：8080）

## 開發中的功能

- 用戶認證和授權
- 請求速率限制
- 緩存層
- 更多數據驗證
- 文件上傳

## 貢獻指南

1. Fork 專案
2. 創建特性分支
3. 提交更改
4. 推送到分支
5. 創建 Pull Request

## 許可證

MIT License 