package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("使用方法: go run main.go <コマンド> [引数...]")
		fmt.Println("利用可能なコマンド:")
		fmt.Println("  hello <名前>     - 挨拶を表示")
		fmt.Println("  calc <数値1> <数値2> - 2つの数値を足し算")
		fmt.Println("  info             - システム情報を表示")
		os.Exit(1)
	}

	command := os.Args[1]

	switch command {
	case "hello":
		handleHello()
	case "calc":
		handleCalc()
	case "info":
		handleInfo()
	default:
		fmt.Printf("不明なコマンド: %s\n", command)
		os.Exit(1)
	}
}

func handleHello() {
	if len(os.Args) < 3 {
		fmt.Println("名前を指定してください")
		fmt.Println("例: go run main.go hello 太郎")
		return
	}

	name := os.Args[2]
	fmt.Printf("こんにちは、%sさん！\n", name)
	fmt.Println("Go言語のCLIツールサンプルへようこそ！")
}

func handleCalc() {
	if len(os.Args) < 4 {
		fmt.Println("2つの数値を指定してください")
		fmt.Println("例: go run main.go calc 10 20")
		return
	}

	num1Str := os.Args[2]
	num2Str := os.Args[3]

	num1, err1 := strconv.Atoi(num1Str)
	num2, err2 := strconv.Atoi(num2Str)

	if err1 != nil || err2 != nil {
		fmt.Println("有効な数値を入力してください")
		return
	}

	result := num1 + num2
	fmt.Printf("%d + %d = %d\n", num1, num2, result)
}

func handleInfo() {
	fmt.Println("=== システム情報 ===")
	fmt.Printf("Go バージョン: %s\n", "1.21")
	fmt.Printf("OS: %s\n", "Linux")
	fmt.Printf("アーキテクチャ: %s\n", "amd64")
	fmt.Println("このプログラムはGo言語で作成されたCLIツールのサンプルです。")
}
