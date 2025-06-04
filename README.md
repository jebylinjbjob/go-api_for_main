# 🌈 可愛的 Go API 專案 ʕ •ᴥ•ʔ

歡迎來到我們充滿愛的 Go API 小天地！這是一個使用 Gin 框架和 MongoDB 的超級可愛 RESTful API 專案 ✨
讓我們一起來管理用戶數據，創造一個溫馨快樂的後端服務吧！(◕‿◕✿)

## 🎨 特色功能

- 🌟 漂亮的 RESTful API 設計
- 🍃 MongoDB 數據小倉庫
- 📚 可愛的 Swagger API 文檔
- 🎯 溫柔的錯誤處理機制
- 🧪 貼心的單元測試
- 📝 細心的日誌記錄

## 🎮 開始冒險前的準備

請確保你有這些小夥伴：
- 🚀 Go 1.21+ 的魔法工具
- 🗄️ MongoDB 的神奇數據庫
- 🐙 Git 的版本控制寶貝

## 🌱 種下專案的種子

1. 首先，把專案帶回家：
```bash
git clone [your-repository-url]
cd go-api_for_main
```

2. 召喚所需的小精靈：
```bash
go mod download
```

3. 確保 MongoDB 小精靈已經醒來了（他喜歡待在 27017 端口）

4. 讓專案綻放光芒：
```bash
go run main.go
```

## 🎯 API 小精靈們

### 👥 用戶管理小隊
- `GET /api/v1/users` - 召喚所有用戶 ✨
- `POST /api/v1/users` - 創造新的小夥伴 🎉
- `GET /api/v1/users/:id` - 尋找特定的朋友 🔍
- `PUT /api/v1/users/:id` - 幫朋友換新衣服 👕
- `DELETE /api/v1/users/:id` - 說再見（揮手） 👋

### 🎪 系統小天地
- `GET /ping` - 戳戳看我們是否還醒著 👉
- `GET /swagger/*any` - 翻閱我們的魔法書 📖

## 📚 魔法使用說明書

想看看更多魔法嗎？運行服務後訪問：
```
http://localhost:8080/swagger/index.html
```

## 🧪 測試小實驗室

測試所有的魔法：
```bash
go test ./... -v
```

測試特定的咒語：
```bash
go test ./test -v
```

## 🏰 專案城堡結構

```
.
├── controllers/     # 🎮 控制中心
├── models/         # 📝 數據模型小屋
├── routes/         # 🛣️ 路線圖
├── docs/          # 📚 魔法書庫
├── test/          # 🧪 實驗室
├── main.go        # 🎯 主要入口
└── README.md      # 📖 使用說明書
```

## 🎨 錯誤處理小幫手

我們用不同的表情來表達不同的情況：

- 200: ✅ 成功啦！
- 400: ❌ 哎呀，請求有點問題
- 404: 🔍 找不到想要的東西
- 500: 😱 服務器打了個噴嚏
- 503: 🏥 數據庫小精靈在休息

## 🎛️ 環境設定小工具

- `MONGODB_URI`: MongoDB 小精靈的家（預設是 mongodb://localhost:27017）
- `PORT`: 服務開門的地方（預設是 8080 號門）

## 🚧 正在建設中的新設施

- 🔐 用戶認證和授權系統
- 🚦 請求速率限制器
- 💾 快取記憶空間
- ✨ 更多數據驗證魔法
- 📤 文件上傳傳送門

## 🌟 加入我們的冒險

想要一起創造魔法嗎？

1. 🍴 Fork 一個自己的分支
2. 🌱 創造新的特性
3. ✨ 提交你的魔法
4. 🚀 推送到你的分支
5. 🎉 發起 Pull Request

## 📜 神奇許可證

MIT License （充滿愛的開源許可證 💝）

---
用 ❤️ 製作，希望你也能感受到這份溫暖！ (｡♥‿♥｡) 