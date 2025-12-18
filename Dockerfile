# 建構階段
FROM golang:1.23-alpine AS builder

# 安裝必要的建構工具
RUN apk add --no-cache git

# 設定工作目錄
WORKDIR /app

# 複製 go mod 檔案
COPY go.mod go.sum ./

# 下載相依套件
RUN go mod download

# 複製原始程式碼
COPY . .

# 建構應用程式
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# 執行階段
FROM alpine:latest

# 安裝 CA 憑證（用於 HTTPS 請求）
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# 從建構階段複製編譯好的二進位檔案
COPY --from=builder /app/main .

# 暴露應用程式埠
EXPOSE 8080

# 執行應用程式
CMD ["./main"]
