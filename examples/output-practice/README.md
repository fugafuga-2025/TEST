# 出力練習サンプル

このディレクトリには、Go言語での様々な出力方法を学ぶためのサンプルプログラムが含まれています。

## 📋 概要

基本的な文字列出力から、フォーマット出力、変数の型表示、応用的な出力技術まで、Go言語の出力機能を包括的に学ぶことができます。

## 🚀 実行方法

### ローカル環境（Goがインストールされている場合）

```bash
# このディレクトリに移動
cd examples/output-practice

# 使用方法を表示
go run main.go

# 基本的な出力練習
go run main.go basic

# フォーマット出力練習
go run main.go format

# 変数と型の出力練習
go run main.go variables

# 応用的な出力練習
go run main.go advanced

# 全ての練習を実行
go run main.go all
```

### Docker環境（Goコマンドが使えない場合）

```bash
# プロジェクトルートディレクトリから実行

# 使用方法を表示
docker run --rm -v $(pwd):/app -w /app golang:1.21-alpine go run examples/output-practice/main.go

# 基本的な出力練習
docker run --rm -v $(pwd):/app -w /app golang:1.21-alpine go run examples/output-practice/main.go basic

# フォーマット出力練習
docker run --rm -v $(pwd):/app -w /app golang:1.21-alpine go run examples/output-practice/main.go format

# 変数と型の出力練習
docker run --rm -v $(pwd):/app -w /app golang:1.21-alpine go run examples/output-practice/main.go variables

# 応用的な出力練習
docker run --rm -v $(pwd):/app -w /app golang:1.21-alpine go run examples/output-practice/main.go advanced

# 全ての練習を実行
docker run --rm -v $(pwd):/app -w /app golang:1.21-alpine go run examples/output-practice/main.go all
```

## 📚 学習ポイント

### 1. 基本的な出力 (`basic`)
- `fmt.Println()`: 改行付き出力
- `fmt.Print()`: 改行なし出力
- 日本語文字列の出力
- 複数行の出力方法

### 2. フォーマット出力 (`format`)
- `fmt.Printf()`: フォーマット指定子を使用した出力
- `%d`: 整数の出力
- `%f`, `%.2f`: 浮動小数点数の出力と桁数指定
- `%s`: 文字列の出力
- `%t`: ブール値の出力
- `%x`, `%X`, `%o`: 16進数、8進数の出力
- 文字列の幅指定と配置

### 3. 変数と型の出力 (`variables`)
- `%T`: 変数の型を表示
- `%v`: 値のデフォルト表示
- `%+v`: 構造体のフィールド名付き表示
- `%p`: ポインタのアドレス表示
- 配列、スライス、マップの出力
- 構造体の出力

### 4. 応用的な出力 (`advanced`)
- 時刻のフォーマット出力
- 表形式の整列出力
- 進捗バーの表示
- ANSIカラーコードを使用した色付き出力
- 複雑なフォーマット指定

## 🔍 サンプル実行結果

### 基本的な出力の例
```
=== 基本的な出力練習 ===
こんにちは、Go言語！
これは基本的な出力練習です。
1行目
2行目
3行目
改行なし1 改行なし2 最後に改行
日本語の文字列: こんにちは、世界！
```

### フォーマット出力の例
```
=== フォーマット出力練習 ===
名前: 田中太郎, 年齢: 25歳
整数: 42
浮動小数点: 3.141590
小数点2桁: 3.14
10進数: 255
16進数: ff
真: true
```

### 変数と型の出力の例
```
=== 変数と型の出力練習 ===
int型: 100 (型: int)
float64型: 3.14 (型: float64)
string型: Hello, Go! (型: string)
bool型: true (型: bool)
スライス: [1 2 3 4 5]
構造体: {Name:山田一郎 Grade:3 Score:85.5}
```

### 応用的な出力の例
```
=== 応用的な出力練習 ===
現在時刻: 2024-01-15 14:30:25
日付のみ: 2024年01月15日

--- 成績表 ---
名前       | 学年 |   点数
-----------|------|--------
田中太郎   |    1 |   85.5
佐藤花子   |    2 |   92.0

--- 進捗表示 ---
進捗: [████░░░░░░░░░░░░░░░░]  20%
進捗: [████████░░░░░░░░░░░░]  40%
```

## 💡 実用的な応用例

- ログ出力の実装
- デバッグ情報の表示
- レポート生成
- コンソールアプリケーションのUI
- データの可視化
- 進捗表示機能

## 🔧 カスタマイズのヒント

- 独自のフォーマット関数の作成
- ログレベル別の出力
- ファイルへの出力機能追加
- JSON形式での出力
- テンプレートエンジンの活用

## 📖 参考資料

- [Go言語 fmt パッケージ](https://pkg.go.dev/fmt)
- [Go言語 time パッケージ](https://pkg.go.dev/time)
- [Printf フォーマット指定子一覧](https://pkg.go.dev/fmt#hdr-Printing)
