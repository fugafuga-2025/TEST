# Go Gin Docker サンプル環境

このリポジトリは、Go言語のGinフレームワークを使用したREST APIのDocker環境構築サンプルです。

## 📋 概要

- **言語**: Go 1.21
- **フレームワーク**: Gin Web Framework
- **コンテナ**: Docker & Docker Compose
- **API**: REST API (GET エンドポイント)

## 🚀 クイックスタート

```bash
# リポジトリをクローン
git clone https://github.com/fugafuga-2025/TEST.git
cd TEST

# Docker環境でアプリケーションを起動
docker-compose up --build

# ブラウザで http://localhost:8080/health にアクセス
```

## 📚 ドキュメント

- [環境構築マニュアル](docs/setup-manual.md) - 詳細なセットアップ手順
- [サンドボックスの解説](docs/sandbox.md) - 簡単なプログラムのサンドボックス
- [API仕様書](docs/api-reference.md) - REST APIの詳細仕様
- [サンプルコード解説](docs/code-explanation.md) - コードの詳細解説

## 🔗 API エンドポイント

| メソッド | エンドポイント | 説明 |
|---------|---------------|------|
| GET | `/health` | ヘルスチェック |
| GET | `/users` | 全ユーザー取得 |
| GET | `/users/:id` | 特定ユーザー取得 |
| GET | `/api/info` | API情報取得 |

## 📁 プロジェクト構成

```
TEST/
├── main.go              # メインアプリケーション
├── go.mod              # Go モジュール定義
├── go.sum              # 依存関係のハッシュ
├── Dockerfile          # Docker イメージ定義
├── docker-compose.yml  # Docker Compose 設定
├── .dockerignore       # Docker ビルド除外ファイル
├── docs/               # ドキュメント
│   ├── setup-manual.md
│   ├── api-reference.md
│   └── code-explanation.md
└── examples/           # サンプルプログラム
    ├── simple-cli/     # CLIツールサンプル
    ├── json-parser/    # JSON処理サンプル
    └── middleware-demo/ # ミドルウェア使用例
    └── test/           # サンドボックス
```

## 🛠 開発環境

- Go 1.21+
- Docker
- Docker Compose

## 🎯 サンプルプログラム

### 1. メインREST API (`main.go`)
- Ginフレームワークを使用したREST API
- ユーザー情報の取得エンドポイント
- ヘルスチェック機能

### 2. CLIツールサンプル (`examples/simple-cli/`)
```bash
cd examples/simple-cli
go run main.go hello 太郎
go run main.go calc 10 20
go run main.go info
```

### 3. JSON処理サンプル (`examples/json-parser/`)
```bash
cd examples/json-parser
go run main.go
```

### 4. ミドルウェアデモ (`examples/middleware-demo/`)
```bash
cd examples/middleware-demo
go run main.go
# 認証が必要なエンドポイントをテスト
curl -H "Authorization: Bearer demo-token" http://localhost:8080/api/v1/users
```

## 📄 ライセンス

MIT License
