package main

import (
	"fmt"
	"os"
)

func main() {
		fmt.Println("Hello, World!")
		os.Exit(0)
}
// 下記コマンドでこのメイン関数を実行出来る
// docker run --rm -v $(pwd):/app -w /app golang:1.21-alpine go run examples/test/main.go
