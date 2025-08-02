# Go Gin Docker 環境構築マニュアル

このマニュアルでは、Go GinフレームワークとDockerを使用したREST API開発環境の構築手順を詳しく説明します。

## 📋 前提条件

### 必要なソフトウェア

1. **Docker** (バージョン 20.10 以上)
2. **Docker Compose** (バージョン 2.0 以上)
3. **Git**

### Dockerのインストール

#### Windows
1. [Docker Desktop for Windows](https://docs.docker.com/desktop/windows/install/) をダウンロード
2. インストーラーを実行してインストール
3. システムを再起動

#### macOS
1. [Docker Desktop for Mac](https://docs.docker.com/desktop/mac/install/) をダウンロード
2. .dmgファイルをマウントしてApplicationsフォルダにドラッグ
3. Docker Desktopを起動

#### Linux (Ubuntu)
```bash
# Dockerの公式GPGキーを追加
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo gpg --dearmor -o /usr/share/keyrings/docker-archive-keyring.gpg

# Dockerリポジトリを追加
echo "deb [arch=amd64 signed-by=/usr/share/keyrings/docker-archive-keyring.gpg] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable" | sudo tee /etc/apt/sources.list.d/docker.list > /dev/null

# パッケージリストを更新してDockerをインストール
sudo apt update
sudo apt install docker-ce docker-ce-cli containerd.io docker-compose-plugin

# 現在のユーザーをdockerグループに追加
sudo usermod -aG docker $USER
```

## 🚀 環境構築手順

### 1. リポジトリのクローン

```bash
# リポジトリをクローン
git clone https://github.com/fugafuga-2025/TEST.git

# プロジェクトディレクトリに移動
cd TEST
```

### 2. プロジェクト構成の確認

```bash
# ファイル構成を確認
ls -la

# 以下のファイルが存在することを確認
# - main.go (メインアプリケーション)
# - go.mod (Go モジュール定義)
# - Dockerfile (Docker イメージ定義)
# - docker-compose.yml (Docker Compose 設定)
```

### 3. Dockerイメージのビルド

```bash
# Dockerイメージをビルド
docker build -t gin-docker-sample .

# ビルドが成功したことを確認
docker images | grep gin-docker-sample
```

### 4. アプリケーションの起動

#### 方法1: Docker Composeを使用（推奨）

```bash
# バックグラウンドでアプリケーションを起動
docker-compose up -d

# ログを確認
docker-compose logs -f

# アプリケーションの状態を確認
docker-compose ps
```

#### 方法2: Dockerコマンドを直接使用

```bash
# コンテナを起動
docker run -d -p 8080:8080 --name gin-sample gin-docker-sample

# ログを確認
docker logs gin-sample

# コンテナの状態を確認
docker ps
```

### 5. 動作確認

#### ヘルスチェック
```bash
curl http://localhost:8080/health
```

期待される応答:
```json
{
  "status": "ok",
  "message": "サーバーは正常に動作しています"
}
```

#### 全ユーザー取得
```bash
curl http://localhost:8080/users
```

期待される応答:
```json
{
  "status": "success",
  "data": [
    {"id": 1, "name": "田中太郎", "age": 25},
    {"id": 2, "name": "佐藤花子", "age": 30},
    {"id": 3, "name": "鈴木一郎", "age": 28}
  ],
  "count": 3
}
```

#### 特定ユーザー取得
```bash
curl http://localhost:8080/users/1
```

期待される応答:
```json
{
  "status": "success",
  "data": {"id": 1, "name": "田中太郎", "age": 25}
}
```

#### API情報取得
```bash
curl http://localhost:8080/api/info
```

## 🛠 開発環境での作業

### ローカル開発（Goがインストールされている場合）

```bash
# 依存関係をダウンロード
go mod download

# アプリケーションを直接実行
go run main.go

# ブラウザで http://localhost:8080 にアクセス
```

### Goコマンドが使用できない場合のサンプル実行

Goがローカルにインストールされていない場合、Dockerを使用してサンプルプログラムを実行できます：

```bash
# プロジェクトルートディレクトリから実行

# 四則演算サンプル
docker run --rm -v $(pwd):/app -w /app golang:1.21-alpine go run examples/arithmetic-operations/main.go add 10 5

# 出力練習サンプル
docker run --rm -v $(pwd):/app -w /app golang:1.21-alpine go run examples/output-practice/main.go basic

# CLIツールサンプル
docker run --rm -v $(pwd):/app -w /app golang:1.21-alpine go run examples/simple-cli/main.go hello テスト

# JSON処理サンプル
docker run --rm -v $(pwd):/app -w /app golang:1.21-alpine go run examples/json-parser/main.go

# ミドルウェアデモ（ポート8080でサーバーを起動）
docker run --rm -v $(pwd):/app -w /app -p 8080:8080 golang:1.21-alpine go run examples/middleware-demo/main.go
```

### コードの変更とテスト

1. `main.go` を編集
2. Dockerイメージを再ビルド
```bash
docker-compose down
docker-compose up --build
```

### ログの確認

```bash
# リアルタイムでログを表示
docker-compose logs -f gin-app

# 特定の行数のログを表示
docker-compose logs --tail=50 gin-app
```

## 🔧 トラブルシューティング

### よくある問題と解決方法

#### 1. ポート8080が既に使用されている

**エラー**: `bind: address already in use`

**解決方法**:
```bash
# 使用中のプロセスを確認
lsof -i :8080

# プロセスを終了
kill -9 <PID>

# または別のポートを使用
docker run -p 8081:8080 gin-docker-sample
```

#### 2. Dockerイメージのビルドに失敗

**エラー**: `failed to solve: failed to compute cache key`

**解決方法**:
```bash
# Dockerキャッシュをクリア
docker system prune -a

# 再度ビルド
docker build --no-cache -t gin-docker-sample .
```

#### 3. コンテナが起動しない

**確認手順**:
```bash
# コンテナのログを確認
docker logs <container_id>

# コンテナ内でシェルを実行
docker exec -it <container_id> /bin/sh
```

### デバッグ用コマンド

```bash
# 実行中のコンテナ一覧
docker ps

# 全てのコンテナ一覧（停止中も含む）
docker ps -a

# イメージ一覧
docker images

# ネットワーク一覧
docker network ls

# ボリューム一覧
docker volume ls
```

## 🧹 環境のクリーンアップ

### 開発環境の停止

```bash
# Docker Composeで起動した場合
docker-compose down

# コンテナとボリュームも削除
docker-compose down -v

# イメージも削除
docker-compose down --rmi all
```

### 完全なクリーンアップ

```bash
# 停止中のコンテナを削除
docker container prune

# 未使用のイメージを削除
docker image prune

# 未使用のボリュームを削除
docker volume prune

# 未使用のネットワークを削除
docker network prune

# 全ての未使用リソースを削除
docker system prune -a
```

## 📚 次のステップ

1. [API仕様書](api-reference.md) でREST APIの詳細を確認
2. [コード解説](code-explanation.md) でソースコードの詳細を学習
3. `examples/` ディレクトリのサンプルプログラムを実行
4. 独自のエンドポイントを追加してみる

## 💡 ヒント

- 開発中は `docker-compose up` でフォアグラウンド実行すると、ログがリアルタイムで確認できます
- コードを変更した後は `docker-compose up --build` で再ビルドが必要です
- 本番環境では環境変数を使用してセキュリティを強化してください
- Ginのドキュメント: https://gin-gonic.com/docs/
