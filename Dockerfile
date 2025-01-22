# 第一段階: ビルド環境
FROM golang:1.23 as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY ./ /app/
WORKDIR /app/cmd/bear_webapi

# ARM64向けに静的バイナリをビルド
RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o main .
RUN ls -al main         # デバッグ: ビルド確認
RUN chmod +x main       # 実行権限を付与

# 第二段階: 実行環境
FROM alpine:latest

WORKDIR /app

# glibcのインストール（必要に応じて）
RUN apk add --no-cache libc6-compat

# ビルドしたバイナリだけをコピー
COPY --from=builder /app/cmd/bear_webapi/main /app/main

EXPOSE 8080

CMD ["/app/main"]
