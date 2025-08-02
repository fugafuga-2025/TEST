package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("使用方法: go run main.go <演算子> <数値1> <数値2>")
		fmt.Println("利用可能な演算子:")
		fmt.Println("  add <数値1> <数値2>    - 足し算")
		fmt.Println("  sub <数値1> <数値2>    - 引き算")
		fmt.Println("  mul <数値1> <数値2>    - 掛け算")
		fmt.Println("  div <数値1> <数値2>    - 割り算")
		fmt.Println("  mod <数値1> <数値2>    - 余り")
		fmt.Println("  all <数値1> <数値2>    - 全ての演算を実行")
		fmt.Println("")
		fmt.Println("例:")
		fmt.Println("  go run main.go add 10 5")
		fmt.Println("  go run main.go all 12 4")
		os.Exit(1)
	}

	operation := os.Args[1]

	switch operation {
	case "add":
		performOperation(addition, "足し算", "+")
	case "sub":
		performOperation(subtraction, "引き算", "-")
	case "mul":
		performOperation(multiplication, "掛け算", "*")
	case "div":
		performOperation(division, "割り算", "/")
	case "mod":
		performOperation(modulo, "余り", "%")
	case "all":
		performAllOperations()
	default:
		fmt.Printf("不明な演算子: %s\n", operation)
		os.Exit(1)
	}
}

func performOperation(op func(float64, float64) (float64, error), name, symbol string) {
	if len(os.Args) < 4 {
		fmt.Printf("%sには2つの数値が必要です\n", name)
		fmt.Printf("例: go run main.go %s 10 5\n", os.Args[1])
		return
	}

	num1, num2, err := parseNumbers()
	if err != nil {
		fmt.Println("有効な数値を入力してください")
		return
	}

	result, err := op(num1, num2)
	if err != nil {
		fmt.Printf("エラー: %s\n", err.Error())
		return
	}

	if symbol == "/" || symbol == "%" {
		fmt.Printf("%.2f %s %.2f = %.2f\n", num1, symbol, num2, result)
	} else {
		fmt.Printf("%.0f %s %.0f = %.0f\n", num1, symbol, num2, result)
	}
}

func performAllOperations() {
	if len(os.Args) < 4 {
		fmt.Println("全ての演算には2つの数値が必要です")
		fmt.Println("例: go run main.go all 12 4")
		return
	}

	num1, num2, err := parseNumbers()
	if err != nil {
		fmt.Println("有効な数値を入力してください")
		return
	}

	fmt.Printf("=== 四則演算の結果 (%.0f と %.0f) ===\n", num1, num2)
	
	result, _ := addition(num1, num2)
	fmt.Printf("足し算: %.0f + %.0f = %.0f\n", num1, num2, result)
	
	result, _ = subtraction(num1, num2)
	fmt.Printf("引き算: %.0f - %.0f = %.0f\n", num1, num2, result)
	
	result, _ = multiplication(num1, num2)
	fmt.Printf("掛け算: %.0f * %.0f = %.0f\n", num1, num2, result)
	
	result, err = division(num1, num2)
	if err != nil {
		fmt.Printf("割り算: エラー - %s\n", err.Error())
	} else {
		fmt.Printf("割り算: %.0f / %.0f = %.2f\n", num1, num2, result)
	}
	
	result, err = modulo(num1, num2)
	if err != nil {
		fmt.Printf("余り: エラー - %s\n", err.Error())
	} else {
		fmt.Printf("余り: %.0f %% %.0f = %.0f\n", num1, num2, result)
	}
}

func parseNumbers() (float64, float64, error) {
	num1Str := os.Args[2]
	num2Str := os.Args[3]

	num1, err1 := strconv.ParseFloat(num1Str, 64)
	num2, err2 := strconv.ParseFloat(num2Str, 64)

	if err1 != nil || err2 != nil {
		return 0, 0, fmt.Errorf("数値の解析に失敗しました")
	}

	return num1, num2, nil
}

func addition(a, b float64) (float64, error) {
	return a + b, nil
}

func subtraction(a, b float64) (float64, error) {
	return a - b, nil
}

func multiplication(a, b float64) (float64, error) {
	return a * b, nil
}

func division(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("ゼロで割ることはできません")
	}
	return a / b, nil
}

func modulo(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("ゼロで余りを求めることはできません")
	}
	return float64(int(a) % int(b)), nil
}
