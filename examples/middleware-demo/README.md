# ミドルウェアデモ

このディレクトリには、Go GinフレームワークでのミドルウェアImplementationの実例が含まれています。

## 📋 概要

認証、ログ出力、CORS対応などの実用的なミドルウェアの実装方法を学ぶことができます。

## 🚀 実行方法

```bash
# このディレクトリに移動
cd examples/middleware-demo

# サーバーを起動
go run main.go
```

## 🔗 エンドポイント

### 認証不要
- `GET /health` - ヘルスチェック
- `GET /public` - 公開エンドポイント

### 認証必要（Authorization: Bearer demo-token）
- `GET /api/v1/users` - ユーザー一覧
- `GET /api/v1/users/:id` - 特定ユーザー
- `GET /api/v1/protected` - 保護されたデータ

## 🧪 テスト方法

### 認証不要のエンドポイント
```bash
curl http://localhost:8080/health
curl http://localhost:8080/public
```

### 認証必要のエンドポイント
```bash
# 認証ヘッダーなし（401エラー）
curl http://localhost:8080/api/v1/users

# 正しい認証ヘッダー付き
curl -H "Authorization: Bearer demo-token" http://localhost:8080/api/v1/users
curl -H "Authorization: Bearer demo-token" http://localhost:8080/api/v1/users/1
curl -H "Authorization: Bearer demo-token" http://localhost:8080/api/v1/protected
```

## 📚 実装されているミドルウェア

### 1. ログミドルウェア (`LoggerMiddleware`)
- リクエストの詳細情報をログ出力
- レスポンス時間の計測
- IPアドレス、HTTPメソッド、ステータスコードの記録

### 2. CORSミドルウェア (`CORSMiddleware`)
- Cross-Origin Resource Sharing対応
- フロントエンドアプリケーションからのアクセスを許可
- プリフライトリクエストの処理

### 3. 認証ミドルウェア (`AuthMiddleware`)
- Bearerトークンによる認証
- 認証失敗時の適切なエラーレスポンス
- ユーザー情報のコンテキストへの設定

## 🔧 学習ポイント

1. **ミドルウェアの作成**: `gin.HandlerFunc`の実装
2. **チェーン処理**: `c.Next()`による次のハンドラーへの制御移譲
3. **早期終了**: `c.Abort()`による処理の中断
4. **コンテキスト操作**: `c.Set()`と`c.Get()`によるデータの受け渡し
5. **グループルーティング**: 特定のルートグループにのみミドルウェアを適用
6. **エラーハンドリング**: 認証失敗時の適切なレスポンス

## 💡 実用的な応用例

- JWT認証の実装
- レート制限の実装
- リクエスト/レスポンスの変換
- セキュリティヘッダーの追加
- アクセスログの記録
