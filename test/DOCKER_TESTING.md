# Docker 容器測試指南 (WSL)

## 前置需求

確保在 WSL 中已安裝 Docker：

```bash
# 檢查 Docker 是否已安裝
docker --version
docker-compose --version

# 如果未安裝，請安裝 Docker Desktop for Windows
# 並在設定中啟用 WSL 2 整合
```

## 啟動服務並執行測試

### 方式 1：啟動服務 + 執行測試

```bash
# 進入專案目錄
cd /mnt/c/side_proj/go-api_for_main

# 啟動 API 和 MongoDB 服務
docker-compose up -d

# 等待服務啟動完成後，執行測試
docker-compose run --rm api-test

# 或者一次性啟動所有服務並執行測試
docker-compose --profile test up --abort-on-container-exit
```

### 方式 2：使用 Docker Compose 配置執行測試

```bash
# 啟動所有服務(包括測試)
docker-compose --profile test up

# 在背景執行
docker-compose --profile test up -d

# 查看測試日誌
docker-compose logs api-test
```

### 方式 3：手動執行特定測試

```bash
# 啟動服務
docker-compose up -d

# 執行所有測試
docker-compose run --rm api-test hurl --test --color /test/*.hurl

# 執行特定測試
docker-compose run --rm api-test hurl --test --verbose /test/create_user.hurl

# 產生 HTML 報告
docker-compose run --rm api-test hurl --test --report-html /test/report /test/all_tests.hurl
```

## 查看測試結果

### 查看測試日誌

```bash
# 即時查看測試日誌
docker-compose logs -f api-test

# 查看所有服務日誌
docker-compose logs -f
```

### 查看測試報告

測試報告會產生在 `test/report` 目錄中(如果執行了 HTML 報告產生)

```bash
# 在 WSL 中開啟報告(使用 Windows 預設瀏覽器)
explorer.exe test/report/index.html

# 或使用 wslview (如果已安裝 wslu 套件)
wslview test/report/index.html
```

## 停止和清理

```bash
# 停止所有服務
docker-compose down

# 停止並刪除資料卷(清理所有資料)
docker-compose down -v

# 重新建構並執行測試
docker-compose build
docker-compose --profile test up --abort-on-container-exit
```

## WSL 特定注意事項

### 檔案路徑

在 WSL 中，Windows 檔案系統掛載在 `/mnt/` 下：

```bash
# Windows 路徑：C:\side_proj\go-api_for_main
# WSL 路徑：/mnt/c/side_proj/go-api_for_main
```

### 效能最佳化

建議將專案放在 WSL 檔案系統中以獲得更好的效能：

```bash
# 複製專案到 WSL 家目錄
cp -r /mnt/c/side_proj/go-api_for_main ~/go-api_for_main
cd ~/go-api_for_main
```

### Docker Desktop 設定

確保 Docker Desktop 已啟用 WSL 2 整合：

1. 開啟 Docker Desktop
2. Settings → Resources → WSL Integration
3. 啟用您的 WSL 發行版

## CI/CD 整合

### GitHub Actions 範例

```yaml
name: API Tests

on: [push, pull_request]

jobs:
  test:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4

      - name: Run tests in Docker
        run: |
          docker-compose --profile test up --abort-on-container-exit

      - name: Check test results
        run: |
          docker-compose logs api-test
```

## 測試檔案說明

在容器中，測試使用 `http://api:8080` 作為基礎 URL(容器內部網路)

- 本機測試：`http://localhost:8080`
- 容器測試：`http://api:8080`

`all_tests.hurl` 已配置為使用容器網路位址。

## 常見問題

### 測試失敗

1. 確認 API 服務健康檢查已通過：

   ```bash
   docker-compose ps
   ```

2. 查看 API 服務日誌：

   ```bash
   docker-compose logs api
   ```

3. 查看 MongoDB 日誌：
   ```bash
   docker-compose logs mongodb
   ```

### 重新執行測試

```bash
# 不需要重啟所有服務，只需重新執行測試容器
docker-compose run --rm api-test
```

### 調試測試

```bash
# 進入測試容器進行互動式調試
docker-compose run --rm --entrypoint sh api-test

# 在容器內手動執行測試
hurl --verbose /test/all_tests.hurl
```

### WSL Docker 無法連線

如果遇到 Docker daemon 連線問題：

```bash
# 確認 Docker 服務執行中
docker ps

# 重啟 WSL (在 PowerShell 中執行)
wsl --shutdown
wsl

# 或重啟 Docker Desktop
```

### 權限問題

```bash
# 如果遇到權限問題，確保當前使用者在 docker 群組中
sudo usermod -aG docker $USER

# 登出後重新登入 WSL
```

## 快速指令參考

```bash
# 啟動所有服務
docker-compose up -d

# 執行測試
docker-compose run --rm api-test

# 查看日誌
docker-compose logs -f

# 停止服務
docker-compose down

# 完整清理
docker-compose down -v

# 重新建構
docker-compose build --no-cache
```

## 優勢

✅ 隔離的測試環境
✅ 不需要本機安裝 Hurl
✅ 確保測試與真實部署環境一致
✅ 易於在 CI/CD 中整合
✅ 可重複執行的測試環境
✅ 跨平台一致性(Windows/WSL/Linux)
