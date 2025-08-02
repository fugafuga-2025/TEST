# REST API 仕様書

このドキュメントでは、Go Gin Docker サンプルアプリケーションが提供するREST APIの詳細仕様を説明します。

## 📋 基本情報

- **ベースURL**: `http://localhost:8080`
- **レスポンス形式**: JSON
- **文字エンコーディング**: UTF-8
- **HTTPメソッド**: GET のみ（このサンプルでは）

## 🔗 エンドポイント一覧

| エンドポイント | メソッド | 説明 | 認証 |
|---------------|---------|------|------|
| `/health` | GET | ヘルスチェック | 不要 |
| `/users` | GET | 全ユーザー取得 | 不要 |
| `/users/:id` | GET | 特定ユーザー取得 | 不要 |
| `/api/info` | GET | API情報取得 | 不要 |

## 📖 詳細仕様

### 1. ヘルスチェック

サーバーの稼働状況を確認するためのエンドポイントです。

#### リクエスト
```
GET /health
```

#### レスポンス

**成功時 (200 OK)**
```json
{
  "status": "ok",
  "message": "サーバーは正常に動作しています"
}
```

#### cURLサンプル
```bash
curl -X GET http://localhost:8080/health
```

---

### 2. 全ユーザー取得

システムに登録されている全ユーザーの一覧を取得します。

#### リクエスト
```
GET /users
```

#### レスポンス

**成功時 (200 OK)**
```json
{
  "status": "success",
  "data": [
    {
      "id": 1,
      "name": "田中太郎",
      "age": 25
    },
    {
      "id": 2,
      "name": "佐藤花子",
      "age": 30
    },
    {
      "id": 3,
      "name": "鈴木一郎",
      "age": 28
    }
  ],
  "count": 3
}
```

#### レスポンスフィールド

| フィールド | 型 | 説明 |
|-----------|---|------|
| `status` | string | 処理結果のステータス |
| `data` | array | ユーザー情報の配列 |
| `data[].id` | integer | ユーザーID |
| `data[].name` | string | ユーザー名 |
| `data[].age` | integer | 年齢 |
| `count` | integer | 取得したユーザー数 |

#### cURLサンプル
```bash
curl -X GET http://localhost:8080/users
```

---

### 3. 特定ユーザー取得

指定されたIDのユーザー情報を取得します。

#### リクエスト
```
GET /users/:id
```

#### パスパラメータ

| パラメータ | 型 | 必須 | 説明 |
|-----------|---|------|------|
| `id` | integer | ✓ | 取得したいユーザーのID |

#### レスポンス

**成功時 (200 OK)**
```json
{
  "status": "success",
  "data": {
    "id": 1,
    "name": "田中太郎",
    "age": 25
  }
}
```

**ユーザーが見つからない場合 (404 Not Found)**
```json
{
  "status": "error",
  "message": "ユーザーが見つかりません"
}
```

**無効なIDの場合 (400 Bad Request)**
```json
{
  "status": "error",
  "message": "無効なユーザーIDです"
}
```

#### cURLサンプル
```bash
# 存在するユーザーを取得
curl -X GET http://localhost:8080/users/1

# 存在しないユーザーを取得（404エラー）
curl -X GET http://localhost:8080/users/999

# 無効なIDを指定（400エラー）
curl -X GET http://localhost:8080/users/abc
```

---

### 4. API情報取得

このAPIの基本情報と利用可能なエンドポイント一覧を取得します。

#### リクエスト
```
GET /api/info
```

#### レスポンス

**成功時 (200 OK)**
```json
{
  "api_name": "Gin Docker Sample API",
  "version": "1.0.0",
  "description": "Go GinとDockerを使ったサンプルREST API",
  "endpoints": [
    "GET /health - ヘルスチェック",
    "GET /users - 全ユーザー取得",
    "GET /users/:id - 特定ユーザー取得",
    "GET /api/info - API情報取得"
  ]
}
```

#### レスポンスフィールド

| フィールド | 型 | 説明 |
|-----------|---|------|
| `api_name` | string | API名 |
| `version` | string | APIバージョン |
| `description` | string | APIの説明 |
| `endpoints` | array | 利用可能なエンドポイント一覧 |

#### cURLサンプル
```bash
curl -X GET http://localhost:8080/api/info
```

## 🔍 エラーレスポンス

### 共通エラー形式

全てのエラーレスポンスは以下の形式で返されます：

```json
{
  "status": "error",
  "message": "エラーメッセージ"
}
```

### HTTPステータスコード

| ステータスコード | 説明 | 発生条件 |
|----------------|------|----------|
| 200 | OK | 正常処理 |
| 400 | Bad Request | 無効なリクエストパラメータ |
| 404 | Not Found | 指定されたリソースが見つからない |
| 500 | Internal Server Error | サーバー内部エラー |

## 🧪 テスト用データ

### サンプルユーザーデータ

アプリケーションには以下のテストデータが含まれています：

```json
[
  {
    "id": 1,
    "name": "田中太郎",
    "age": 25
  },
  {
    "id": 2,
    "name": "佐藤花子",
    "age": 30
  },
  {
    "id": 3,
    "name": "鈴木一郎",
    "age": 28
  }
]
```

## 📝 使用例

### JavaScript (Fetch API)

```javascript
// 全ユーザー取得
fetch('http://localhost:8080/users')
  .then(response => response.json())
  .then(data => {
    console.log('ユーザー一覧:', data.data);
    console.log('ユーザー数:', data.count);
  })
  .catch(error => console.error('エラー:', error));

// 特定ユーザー取得
fetch('http://localhost:8080/users/1')
  .then(response => response.json())
  .then(data => {
    if (data.status === 'success') {
      console.log('ユーザー情報:', data.data);
    } else {
      console.error('エラー:', data.message);
    }
  });
```

### Python (requests)

```python
import requests

# 全ユーザー取得
response = requests.get('http://localhost:8080/users')
if response.status_code == 200:
    data = response.json()
    print(f"ユーザー数: {data['count']}")
    for user in data['data']:
        print(f"ID: {user['id']}, 名前: {user['name']}, 年齢: {user['age']}")

# 特定ユーザー取得
user_id = 1
response = requests.get(f'http://localhost:8080/users/{user_id}')
if response.status_code == 200:
    data = response.json()
    user = data['data']
    print(f"ユーザー情報: {user['name']} ({user['age']}歳)")
elif response.status_code == 404:
    print("ユーザーが見つかりません")
```

### Go (net/http)

```go
package main

import (
    "encoding/json"
    "fmt"
    "net/http"
)

type User struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
    Age  int    `json:"age"`
}

type UsersResponse struct {
    Status string `json:"status"`
    Data   []User `json:"data"`
    Count  int    `json:"count"`
}

func main() {
    // 全ユーザー取得
    resp, err := http.Get("http://localhost:8080/users")
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    var usersResp UsersResponse
    json.NewDecoder(resp.Body).Decode(&usersResp)

    fmt.Printf("ユーザー数: %d\n", usersResp.Count)
    for _, user := range usersResp.Data {
        fmt.Printf("ID: %d, 名前: %s, 年齢: %d\n", user.ID, user.Name, user.Age)
    }
}
```

## 🔧 開発者向け情報

### レスポンス時間

- 通常のリクエスト: < 10ms
- ヘルスチェック: < 5ms

### レート制限

現在のサンプルではレート制限は実装されていませんが、本番環境では以下の制限を推奨します：

- 1秒あたり100リクエスト
- 1時間あたり10,000リクエスト

### CORS設定

現在のサンプルではCORSは設定されていません。フロントエンドアプリケーションから利用する場合は、適切なCORS設定を追加してください。

## 📚 関連ドキュメント

- [環境構築マニュアル](setup-manual.md)
- [コード解説](code-explanation.md)
- [Gin公式ドキュメント](https://gin-gonic.com/docs/)
