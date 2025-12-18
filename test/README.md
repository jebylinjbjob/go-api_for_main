# Hurl 測試指南

## 安裝 Hurl

### Windows

```powershell
# 使用 winget
winget install hurl

# 或使用 Chocolatey
choco install hurl
```

### Linux/WSL

```bash
# Ubuntu/Debian
sudo apt-get update
sudo apt-get install hurl

# 或使用 cargo
cargo install hurl
```

### macOS

```bash
brew install hurl
```

## 執行測試

### 執行單個測試檔案

```bash
# 執行獲取用戶測試
hurl test/get_users.hurl

# 執行創建用戶測試
hurl test/create_user.hurl

# 執行更新用戶測試
hurl test/update_user.hurl

# 執行刪除用戶測試
hurl test/delete_user.hurl

# 執行完整測試套件
hurl test/all_tests.hurl
```

### 執行所有測試

```bash
# 執行 test 目錄下所有 .hurl 檔案
hurl test/*.hurl

# 詳細輸出
hurl --verbose test/*.hurl

# 生成測試報告
hurl --test --report-html test-report test/*.hurl
```

### 常用選項

```bash
# 顯示詳細資訊
hurl --verbose test/all_tests.hurl

# 只顯示錯誤
hurl --error-format long test/*.hurl

# 設定變數
hurl --variable base_url=http://localhost:8080 test/all_tests.hurl

# 生成 JSON 報告
hurl --test --report-json test-report.json test/*.hurl

# 生成 HTML 報告
hurl --test --report-html test-report test/*.hurl
```

## 測試檔案說明

- `get_users.hurl` - 測試獲取所有用戶
- `get_user.hurl` - 測試獲取特定用戶
- `create_user.hurl` - 測試創建用戶及驗證
- `update_user.hurl` - 測試更新用戶
- `delete_user.hurl` - 測試刪除用戶
- `all_tests.hurl` - 完整測試流程

## 在 CI/CD 中使用

### GitHub Actions

```yaml
- name: Run Hurl Tests
  run: |
    hurl --test --report-html test-report test/*.hurl
```

### Docker Compose

可以在 docker-compose.yaml 中添加測試服務：

```yaml
services:
  api-test:
    image: ghcr.io/orange-opensource/hurl:latest
    depends_on:
      - api
    volumes:
      - ./test:/test
    command: hurl --test /test/*.hurl
```

## 注意事項

1. 確保 API 服務正在運行（`http://localhost:8080`）
2. 測試會創建、修改和刪除真實數據
3. 建議在測試環境中執行，不要在生產環境中執行
4. 某些測試之間可能有依賴關係，建議按順序執行

## 替代 Go 測試

原本的 Go 測試檔案：

- `test/user_test.go`

現在已被以下 Hurl 測試檔案取代：

- `test/get_users.hurl`
- `test/get_user.hurl`
- `test/create_user.hurl`
- `test/update_user.hurl`
- `test/delete_user.hurl`
- `test/all_tests.hurl`

Hurl 測試的優勢：

- ✅ 不需要編寫程式碼
- ✅ 可以測試實際運行的 API
- ✅ 易於閱讀和維護
- ✅ 支援生成測試報告
- ✅ 可以在 CI/CD 中輕鬆整合
