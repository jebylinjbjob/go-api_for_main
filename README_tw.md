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
- 🔄 符合 HATEOAS 原則的 API 響應，提供更好的探索性

## 🎮 開始冒險前的準備

請確保你有這些小夥伴：
- 🚀 Go 1.21+ 的魔法工具
- 🗄️ MongoDB 的神奇數據庫
- 🐙 Git 的版本控制寶貝

## 🌱 種下專案的種子

1. 首先，把專案帶回家：
```bash
git clone https://github.com/jebylinjbjob/go-api_for_main.git
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

### 🧪 測試環境小天地
在測試時，我們把魔法路徑從 v1 變成 test，讓我們一起來玩耍吧！

以下是測試環境的小精靈們：
- `GET /api/test/users` - 呼喚測試用戶來玩耍 🎈
- `POST /api/test/users` - 邀請新朋友加入遊戲 🎪
- `GET /api/test/users/:id` - 找找看誰來玩了 🔮
- `PUT /api/test/users/:id` - 讓朋友換個造型玩玩 🎭
- `DELETE /api/test/users/:id` - 跟朋友玩捉迷藏（暫時說掰掰）🎪


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
│   ├── response.go  # 🔄 HATEOAS 響應結構
│   └── user.go      # 👤 使用者資料模型
├── routes/         # 🛣️ 路線圖
├── docs/          # 📚 魔法書庫
├── test/          # 🧪 實驗室
├── main.go        # 🎯 主要入口
└── README.md      # 📖 使用說明書
```

## 🔄 HATEOAS API 響應

我們的 API 遵循 HATEOAS（Hypermedia as the Engine of Application State，超媒體作為應用程式狀態引擎）原則，讓客戶端可以透過超媒體連結動態地導航 API：

- 🔗 所有響應都包含相關行為的連結
- 🧭 客戶端可以透過響應中的連結發現可用的操作
- 🔍 客戶端應用無需硬編碼 API 端點
- 🎭 支援 API 的演化，只需對客戶端做最小的更改

HATEOAS 響應範例：
```json
{
  "data": {
    "id": "507f1f77bcf86cd799439011",
    "name": "張三",
    "email": "zhangsan@example.com",
    "age": 30
  },
  "_links": [
    {
      "href": "http://api.example.com/users/507f1f77bcf86cd799439011",
      "rel": "self",
      "method": "GET",
      "title": "取得使用者資訊"
    },
    {
      "href": "http://api.example.com/users/507f1f77bcf86cd799439011",
      "rel": "update",
      "method": "PUT",
      "title": "更新使用者"
    },
    {
      "href": "http://api.example.com/users/507f1f77bcf86cd799439011",
      "rel": "delete",
      "method": "DELETE",
      "title": "刪除使用者"
    }
  ]
}
```

## 🎨 錯誤處理小幫手

我們用不同的表情來表達不同的情況：

- 200: ✅ 成功啦！
- 400: ❌ 哎呀，請求有點問題
- 404: 🔍 找不到想要的東西
- 500: 😱 服務器打了個噴嚏
- 503: 🏥 數據庫小精靈在休息

## 🎛️ 環境設定小工具

### 🗄️ MongoDB 設定
- `MONGODB_URI`: MongoDB 小精靈的家（預設是 mongodb://localhost:27017）
- `MONGODB_DATABASE`: 數據庫的名字（預設是 go_api_db）
- `MONGODB_USERNAME`: MongoDB 小精靈的用戶名（選填）
- `MONGODB_PASSWORD`: MongoDB 小精靈的通關密語（選填）
- `MONGODB_TIMEOUT`: 等待 MongoDB 小精靈回應的時間（預設是 10 秒）

💫 **MongoDB 安全小貼士**:
1. 有兩種方式可以告訴 MongoDB 小精靈你是誰：
   - 使用 `MONGODB_USERNAME` 和 `MONGODB_PASSWORD` 兩個小精靈
   - 或是直接在 `MONGODB_URI` 中寫入，像這樣：`mongodb://username:password@localhost:27017`
2. 建議把這些神秘咒語放在 `.env` 檔案中
3. 千萬不要把密碼放進 Git 倉庫哦
4. 在正式環境中，最好使用專門的秘密保管系統

### 🖥️ 服務器設定
- `PORT`: 服務開門的地方（預設是 8080 號門）
- `GIN_MODE`: 服務運行模式（預設是 debug 模式）

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


[![image](https://github.com/jebylinjbjob/go-api_for_main/blob/main/ICON.jpeg))](https://github.com/jebylinjbjob/go-api_for_main/blob/main/ICON.jpeg)
