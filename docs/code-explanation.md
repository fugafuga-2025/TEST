# コード解説

このドキュメントでは、Go Gin Docker サンプルアプリケーションのソースコードを詳しく解説します。

## 📁 プロジェクト構成

```
TEST/
├── main.go              # メインアプリケーション
├── go.mod              # Go モジュール定義
├── Dockerfile          # Docker イメージ定義
├── docker-compose.yml  # Docker Compose 設定
├── .dockerignore       # Docker ビルド除外ファイル
├── docs/               # ドキュメント
└── examples/           # サンプルプログラム
```

## 🔍 main.go 詳細解説

### パッケージとインポート

```go
package main

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
)
```

- `package main`: 実行可能なプログラムのメインパッケージ
- `net/http`: HTTPステータスコード定数を使用
- `strconv`: 文字列と数値の変換に使用
- `github.com/gin-gonic/gin`: Gin Webフレームワーク

### データ構造定義

```go
type User struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
    Age  int    `json:"age"`
}
```

**解説:**
- `User`構造体: ユーザー情報を表現
- `json:"id"`: JSONシリアライゼーション時のフィールド名を指定
- 構造体タグにより、GoのフィールドとJSONキーのマッピングを定義

### サンプルデータ

```go
var users = []User{
    {ID: 1, Name: "田中太郎", Age: 25},
    {ID: 2, Name: "佐藤花子", Age: 30},
    {ID: 3, Name: "鈴木一郎", Age: 28},
}
```

**解説:**
- グローバル変数として定義されたユーザーデータ
- 実際のアプリケーションではデータベースを使用
- 日本語の名前を使用してローカライゼーションの例を示す

### メイン関数

```go
func main() {
    r := gin.Default()
    
    // ルート定義
    r.GET("/health", func(c *gin.Context) { ... })
    r.GET("/users", getUsers)
    r.GET("/users/:id", getUserByID)
    r.GET("/api/info", func(c *gin.Context) { ... })
    
    r.Run(":8080")
}
```

**解説:**
- `gin.Default()`: デフォルトのミドルウェア（ログ、リカバリ）付きのGinエンジンを作成
- `r.GET()`: GETメソッドのルートを定義
- `:id`: パスパラメータの定義
- `r.Run(":8080")`: ポート8080でサーバーを起動

### ヘルスチェックエンドポイント

```go
r.GET("/health", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "status":  "ok",
        "message": "サーバーは正常に動作しています",
    })
})
```

**解説:**
- 無名関数（クロージャ）を使用したインラインハンドラ
- `c.JSON()`: JSON形式でレスポンスを返す
- `gin.H`: `map[string]interface{}`のエイリアス
- `http.StatusOK`: HTTPステータス200を表す定数

### 全ユーザー取得関数

```go
func getUsers(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "status": "success",
        "data":   users,
        "count":  len(users),
    })
}
```

**解説:**
- 独立した関数として定義
- `len(users)`: スライスの長さを取得
- 構造化されたレスポンス形式（status, data, count）

### 特定ユーザー取得関数

```go
func getUserByID(c *gin.Context) {
    idParam := c.Param("id")
    id, err := strconv.Atoi(idParam)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "status":  "error",
            "message": "無効なユーザーIDです",
        })
        return
    }

    for _, user := range users {
        if user.ID == id {
            c.JSON(http.StatusOK, gin.H{
                "status": "success",
                "data":   user,
            })
            return
        }
    }

    c.JSON(http.StatusNotFound, gin.H{
        "status":  "error",
        "message": "ユーザーが見つかりません",
    })
}
```

**解説:**
- `c.Param("id")`: URLパスパラメータを取得
- `strconv.Atoi()`: 文字列を整数に変換
- エラーハンドリング: 無効なIDの場合は400エラー
- `range`: スライスをイテレート
- 早期リターン: 条件に合致した時点で処理を終了
- 404エラー: ユーザーが見つからない場合

## 🐳 Dockerfile 解説

### マルチステージビルド

```dockerfile
FROM golang:1.21-alpine AS builder
```

**解説:**
- `golang:1.21-alpine`: Go 1.21がインストールされたAlpine Linuxベースイメージ
- `AS builder`: ビルドステージに名前を付ける
- Alpine Linuxは軽量で本番環境に適している

### 作業ディレクトリ設定

```dockerfile
WORKDIR /app
```

**解説:**
- コンテナ内の作業ディレクトリを`/app`に設定
- 以降のコマンドはこのディレクトリで実行される

### 依存関係の管理

```dockerfile
COPY go.mod ./
RUN go mod download
```

**解説:**
- `go.mod`のみを先にコピー
- `go mod download`: 依存関係をダウンロード
- Dockerのレイヤーキャッシュを効率的に活用
- `go.sum`は`go mod download`時に自動生成される

### ソースコードのコピーとビルド

```dockerfile
COPY . .
RUN go build -o main .
```

**解説:**
- 全ソースコードをコンテナにコピー
- `go build -o main .`: 実行ファイル`main`を生成
- 静的リンクされたバイナリが作成される

### 実行用イメージ

```dockerfile
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/main .
EXPOSE 8080
CMD ["./main"]
```

**解説:**
- `alpine:latest`: 軽量な実行用ベースイメージ
- `ca-certificates`: HTTPS通信に必要な証明書
- `COPY --from=builder`: ビルドステージからバイナリをコピー
- `EXPOSE 8080`: ポート8080を公開（ドキュメント目的）
- `CMD`: コンテナ起動時に実行するコマンド

## 🔧 docker-compose.yml 解説

```yaml
version: '3.8'

services:
  gin-app:
    build: .
    ports:
      - "8080:8080"
    environment:
      - GIN_MODE=release
    restart: unless-stopped
    container_name: gin-docker-sample
```

**解説:**
- `version: '3.8'`: Docker Composeファイルのバージョン
- `build: .`: 現在のディレクトリのDockerfileを使用してビルド
- `ports`: ホストの8080ポートをコンテナの8080ポートにマッピング
- `GIN_MODE=release`: Ginを本番モードで実行
- `restart: unless-stopped`: コンテナが停止した場合に自動再起動
- `container_name`: コンテナに固定名を設定

## 📝 .dockerignore 解説

```
*.md
.git
.gitignore
Dockerfile
docker-compose.yml
.dockerignore
```

**解説:**
- Dockerビルド時にコンテキストから除外するファイルを指定
- ビルド時間の短縮とイメージサイズの削減
- 不要なファイルがコンテナに含まれることを防ぐ

## 🔧 go.mod 解説

```go
module gin-docker-sample

go 1.21

require (
    github.com/gin-gonic/gin v1.9.1
)
```

**解説:**
- `module gin-docker-sample`: モジュール名の定義
- `go 1.21`: 使用するGoのバージョン
- `require`: 依存関係の指定
- セマンティックバージョニングを使用

## 🎯 設計パターンと最適化

### RESTful API設計

- **統一されたレスポンス形式**: 全エンドポイントで一貫したJSON構造
- **適切なHTTPステータスコード**: 200, 400, 404の使い分け
- **エラーハンドリング**: 分かりやすい日本語エラーメッセージ

### パフォーマンス最適化

- **マルチステージビルド**: 最終イメージサイズの最小化
- **Dockerレイヤーキャッシュ**: 依存関係の効率的なキャッシュ
- **静的バイナリ**: 実行時依存関係の排除

### セキュリティ考慮事項

- **非rootユーザー**: 本番環境では非rootユーザーでの実行を推奨
- **最小権限の原則**: 必要最小限のパッケージのみインストール
- **入力検証**: パラメータの適切な検証

## 🚀 拡張のヒント

### データベース統合

```go
// データベース接続の例
import (
    "gorm.io/gorm"
    "gorm.io/driver/postgres"
)

func initDB() *gorm.DB {
    dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable"
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }
    return db
}
```

### ミドルウェアの追加

```go
// CORS ミドルウェアの例
func CORSMiddleware() gin.HandlerFunc {
    return gin.HandlerFunc(func(c *gin.Context) {
        c.Header("Access-Control-Allow-Origin", "*")
        c.Header("Access-Control-Allow-Credentials", "true")
        c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
        c.Header("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }

        c.Next()
    })
}

// 使用方法
r.Use(CORSMiddleware())
```

### ログ設定

```go
// カスタムログ設定の例
import "github.com/gin-gonic/gin"

func setupLogger() {
    gin.SetMode(gin.ReleaseMode)
    gin.DefaultWriter = io.MultiWriter(os.Stdout, logFile)
}
```

## 📚 参考資料

- [Gin公式ドキュメント](https://gin-gonic.com/docs/)
- [Go公式ドキュメント](https://golang.org/doc/)
- [Docker公式ドキュメント](https://docs.docker.com/)
- [Docker Compose公式ドキュメント](https://docs.docker.com/compose/)

## 💡 学習のポイント

1. **Ginフレームワークの基本**: ルーティング、ミドルウェア、レスポンス処理
2. **Goの構造体とJSON**: データ構造の定義とシリアライゼーション
3. **Dockerマルチステージビルド**: 効率的なコンテナイメージの作成
4. **RESTful API設計**: 一貫性のあるAPI設計原則
5. **エラーハンドリング**: 適切なエラー処理とユーザーフレンドリーなメッセージ
