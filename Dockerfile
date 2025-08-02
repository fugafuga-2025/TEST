FROM golang:1.21-alpine AS builder

# 作業ディレクトリを設定
WORKDIR /app

# ソースコードをコピー
COPY . .

# 依存関係をダウンロードしてgo.sumを生成
RUN go mod tidy

# アプリケーションをビルド
RUN go build -o main .

# 実行用の軽量イメージ
FROM alpine:latest

# 必要なパッケージをインストール
RUN apk --no-cache add ca-certificates

# 作業ディレクトリを設定
WORKDIR /root/

# ビルドしたバイナリをコピー
COPY --from=builder /app/main .

# ポート8080を公開
EXPOSE 8080

# アプリケーションを実行
CMD ["./main"]
