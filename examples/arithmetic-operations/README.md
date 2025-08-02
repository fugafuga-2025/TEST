# 四則演算サンプル

このディレクトリには、Go言語での基本的な四則演算（足し算、引き算、掛け算、割り算、余り）を学ぶためのサンプルプログラムが含まれています。

## 📋 概要

数値の入力処理、エラーハンドリング、基本的な算術演算の実装方法を学ぶことができます。

## 🚀 実行方法

### ローカル環境（Goがインストールされている場合）

```bash
# このディレクトリに移動
cd examples/arithmetic-operations

# 使用方法を表示
go run main.go

# 足し算
go run main.go add 10 5

# 引き算
go run main.go sub 20 8

# 掛け算
go run main.go mul 6 7

# 割り算
go run main.go div 15 3

# 余り
go run main.go mod 17 5

# 全ての演算を実行
go run main.go all 12 4
```

### Docker環境（Goコマンドが使えない場合）

```bash
# プロジェクトルートディレクトリから実行

# 使用方法を表示
docker run --rm -v $(pwd):/app -w /app golang:1.21-alpine go run examples/arithmetic-operations/main.go

# 足し算
docker run --rm -v $(pwd):/app -w /app golang:1.21-alpine go run examples/arithmetic-operations/main.go add 10 5

# 引き算
docker run --rm -v $(pwd):/app -w /app golang:1.21-alpine go run examples/arithmetic-operations/main.go sub 20 8

# 掛け算
docker run --rm -v $(pwd):/app -w /app golang:1.21-alpine go run examples/arithmetic-operations/main.go mul 6 7

# 割り算
docker run --rm -v $(pwd):/app -w /app golang:1.21-alpine go run examples/arithmetic-operations/main.go div 15 3

# 余り
docker run --rm -v $(pwd):/app -w /app golang:1.21-alpine go run examples/arithmetic-operations/main.go mod 17 5

# 全ての演算を実行
docker run --rm -v $(pwd):/app -w /app golang:1.21-alpine go run examples/arithmetic-operations/main.go all 12 4
```

## 📚 学習ポイント

1. **基本的な算術演算**: 四則演算（+, -, *, /, %）の実装
2. **数値の型変換**: `strconv.ParseFloat()`を使用した文字列から数値への変換
3. **エラーハンドリング**: ゼロ除算などの例外処理
4. **関数の設計**: 各演算を独立した関数として実装
5. **コマンドライン引数**: 演算子と数値の受け取り
6. **フォーマット出力**: `fmt.Printf()`での数値の表示制御

## 🔍 サンプル実行結果

### 個別演算の例
```bash
$ go run main.go add 15 25
15 + 25 = 40

$ go run main.go div 20 3
20.00 / 3.00 = 6.67

$ go run main.go mod 17 5
17 % 5 = 2
```

### 全演算の例
```bash
$ go run main.go all 12 4
=== 四則演算の結果 (12 と 4) ===
足し算: 12 + 4 = 16
引き算: 12 - 4 = 8
掛け算: 12 * 4 = 48
割り算: 12 / 4 = 3.00
余り: 12 % 4 = 0
```

### エラーハンドリングの例
```bash
$ go run main.go div 10 0
エラー: ゼロで割ることはできません

$ go run main.go add abc 5
有効な数値を入力してください
```

## 💡 実用的な応用例

- 計算機アプリケーションの基礎
- 数値処理プログラムの作成
- ユーザー入力の検証とエラー処理
- 関数型プログラミングの基礎

## 🔧 カスタマイズのヒント

- 小数点以下の桁数を変更: `%.2f` の数値を調整
- 新しい演算の追加: 累乗、平方根など
- 対話型モードの実装: `bufio.Scanner`を使用
- 計算履歴の保存: ファイル出力機能の追加
