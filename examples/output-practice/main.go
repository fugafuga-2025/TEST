package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("使用方法: go run main.go <練習タイプ>")
		fmt.Println("利用可能な練習タイプ:")
		fmt.Println("  basic     - 基本的な出力練習")
		fmt.Println("  format    - フォーマット出力練習")
		fmt.Println("  variables - 変数と型の出力練習")
		fmt.Println("  advanced  - 応用的な出力練習")
		fmt.Println("  all       - 全ての練習を実行")
		fmt.Println("")
		fmt.Println("例:")
		fmt.Println("  go run main.go basic")
		fmt.Println("  go run main.go all")
		os.Exit(1)
	}

	practiceType := os.Args[1]

	switch practiceType {
	case "basic":
		basicOutput()
	case "format":
		formatOutput()
	case "variables":
		variablesOutput()
	case "advanced":
		advancedOutput()
	case "all":
		allPractices()
	default:
		fmt.Printf("不明な練習タイプ: %s\n", practiceType)
		os.Exit(1)
	}
}

func basicOutput() {
	fmt.Println("=== 基本的な出力練習 ===")
	
	fmt.Println("こんにちは、Go言語！")
	fmt.Println("これは基本的な出力練習です。")
	
	fmt.Println("1行目")
	fmt.Println("2行目")
	fmt.Println("3行目")
	
	fmt.Print("改行なし1 ")
	fmt.Print("改行なし2 ")
	fmt.Println("最後に改行")
	
	fmt.Println("日本語の文字列: こんにちは、世界！")
	fmt.Println("ひらがな: あいうえお")
	fmt.Println("カタカナ: アイウエオ")
	fmt.Println("漢字: 日本語プログラミング")
}

func formatOutput() {
	fmt.Println("\n=== フォーマット出力練習 ===")
	
	name := "田中太郎"
	age := 25
	fmt.Printf("名前: %s, 年齢: %d歳\n", name, age)
	
	integer := 42
	float := 3.14159
	fmt.Printf("整数: %d\n", integer)
	fmt.Printf("浮動小数点: %f\n", float)
	fmt.Printf("小数点2桁: %.2f\n", float)
	fmt.Printf("小数点0桁: %.0f\n", float)
	
	text := "Go言語"
	fmt.Printf("文字列: %s\n", text)
	fmt.Printf("文字列（幅10、右寄せ）: %10s\n", text)
	fmt.Printf("文字列（幅10、左寄せ）: %-10s|\n", text)
	
	number := 255
	fmt.Printf("10進数: %d\n", number)
	fmt.Printf("16進数: %x\n", number)
	fmt.Printf("16進数（大文字）: %X\n", number)
	fmt.Printf("8進数: %o\n", number)
	
	isTrue := true
	isFalse := false
	fmt.Printf("真: %t\n", isTrue)
	fmt.Printf("偽: %t\n", isFalse)
}

func variablesOutput() {
	fmt.Println("\n=== 変数と型の出力練習 ===")
	
	var intVar int = 100
	var floatVar float64 = 3.14
	var stringVar string = "Hello, Go!"
	var boolVar bool = true
	
	fmt.Printf("int型: %d (型: %T)\n", intVar, intVar)
	fmt.Printf("float64型: %.2f (型: %T)\n", floatVar, floatVar)
	fmt.Printf("string型: %s (型: %T)\n", stringVar, stringVar)
	fmt.Printf("bool型: %t (型: %T)\n", boolVar, boolVar)
	
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Printf("スライス: %v\n", numbers)
	fmt.Printf("スライスの長さ: %d\n", len(numbers))
	
	person := map[string]interface{}{
		"名前": "佐藤花子",
		"年齢": 30,
		"職業": "エンジニア",
	}
	fmt.Printf("マップ: %v\n", person)
	
	x := 42
	ptr := &x
	fmt.Printf("変数x: %d\n", x)
	fmt.Printf("ポインタptr: %p\n", ptr)
	fmt.Printf("ポインタが指す値: %d\n", *ptr)
	
	type Student struct {
		Name  string
		Grade int
		Score float64
	}
	
	student := Student{
		Name:  "山田一郎",
		Grade: 3,
		Score: 85.5,
	}
	fmt.Printf("構造体: %+v\n", student)
	fmt.Printf("学生名: %s, 学年: %d年, 点数: %.1f点\n", student.Name, student.Grade, student.Score)
}

func advancedOutput() {
	fmt.Println("\n=== 応用的な出力練習 ===")
	
	now := time.Now()
	fmt.Printf("現在時刻: %s\n", now.Format("2006-01-02 15:04:05"))
	fmt.Printf("日付のみ: %s\n", now.Format("2006年01月02日"))
	fmt.Printf("時刻のみ: %s\n", now.Format("15時04分05秒"))
	
	fmt.Println("\n--- 成績表 ---")
	fmt.Printf("%-10s | %4s | %6s\n", "名前", "学年", "点数")
	fmt.Println("-----------|------|--------")
	fmt.Printf("%-10s | %4d | %6.1f\n", "田中太郎", 1, 85.5)
	fmt.Printf("%-10s | %4d | %6.1f\n", "佐藤花子", 2, 92.0)
	fmt.Printf("%-10s | %4d | %6.1f\n", "鈴木一郎", 3, 78.5)
	
	fmt.Println("\n--- 進捗表示 ---")
	for i := 0; i <= 100; i += 20 {
		bar := ""
		for j := 0; j < i/5; j++ {
			bar += "█"
		}
		for j := i/5; j < 20; j++ {
			bar += "░"
		}
		fmt.Printf("進捗: [%s] %3d%%\n", bar, i)
	}
	
	fmt.Println("\n--- カラー出力 ---")
	fmt.Printf("\033[31m赤色のテキスト\033[0m\n")
	fmt.Printf("\033[32m緑色のテキスト\033[0m\n")
	fmt.Printf("\033[34m青色のテキスト\033[0m\n")
	fmt.Printf("\033[33m黄色のテキスト\033[0m\n")
	fmt.Printf("\033[1m太字のテキスト\033[0m\n")
	
	fmt.Println("\n--- 複雑なフォーマット ---")
	price := 1234.56
	fmt.Printf("価格: ¥%,.2f\n", price)
	
	fmt.Printf("複数の値: %d, %s, %.2f, %t\n", 42, "テスト", 3.14, true)
}

func allPractices() {
	fmt.Println("=== 全ての出力練習を実行 ===")
	
	basicOutput()
	formatOutput()
	variablesOutput()
	advancedOutput()
	
	fmt.Println("\n=== 練習完了 ===")
	fmt.Println("お疲れ様でした！Go言語の出力機能をマスターしましたね。")
}
