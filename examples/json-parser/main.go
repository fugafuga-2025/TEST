package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

type Person struct {
	Name     string   `json:"name"`
	Age      int      `json:"age"`
	Email    string   `json:"email"`
	Hobbies  []string `json:"hobbies"`
	Address  Address  `json:"address"`
	IsActive bool     `json:"is_active"`
}

type Address struct {
	Prefecture string `json:"prefecture"`
	City       string `json:"city"`
	PostalCode string `json:"postal_code"`
}

func main() {
	fmt.Println("=== Go JSON処理サンプル ===")

	jsonExample()
	fmt.Println()

	parseExample()
	fmt.Println()

	createExample()
}

func jsonExample() {
	fmt.Println("1. JSON文字列の解析")

	jsonData := `{
		"name": "田中太郎",
		"age": 30,
		"email": "tanaka@example.com",
		"hobbies": ["読書", "映画鑑賞", "プログラミング"],
		"address": {
			"prefecture": "東京都",
			"city": "渋谷区",
			"postal_code": "150-0001"
		},
		"is_active": true
	}`

	var person Person
	err := json.Unmarshal([]byte(jsonData), &person)
	if err != nil {
		log.Fatal("JSON解析エラー:", err)
	}

	fmt.Printf("名前: %s\n", person.Name)
	fmt.Printf("年齢: %d歳\n", person.Age)
	fmt.Printf("メール: %s\n", person.Email)
	fmt.Printf("住所: %s %s (%s)\n", person.Address.Prefecture, person.Address.City, person.Address.PostalCode)
	fmt.Printf("趣味: %s\n", strings.Join(person.Hobbies, ", "))
	fmt.Printf("アクティブ: %t\n", person.IsActive)
}

func parseExample() {
	fmt.Println("2. 複数のJSONオブジェクトの処理")

	jsonArray := `[
		{
			"name": "佐藤花子",
			"age": 25,
			"email": "sato@example.com",
			"hobbies": ["料理", "旅行"],
			"address": {
				"prefecture": "大阪府",
				"city": "大阪市",
				"postal_code": "530-0001"
			},
			"is_active": true
		},
		{
			"name": "鈴木一郎",
			"age": 35,
			"email": "suzuki@example.com",
			"hobbies": ["スポーツ", "音楽"],
			"address": {
				"prefecture": "愛知県",
				"city": "名古屋市",
				"postal_code": "460-0001"
			},
			"is_active": false
		}
	]`

	var people []Person
	err := json.Unmarshal([]byte(jsonArray), &people)
	if err != nil {
		log.Fatal("JSON配列解析エラー:", err)
	}

	fmt.Printf("総人数: %d人\n", len(people))
	for i, person := range people {
		fmt.Printf("%d. %s (%d歳) - %s\n", i+1, person.Name, person.Age, person.Address.Prefecture)
	}

	activeCount := 0
	for _, person := range people {
		if person.IsActive {
			activeCount++
		}
	}
	fmt.Printf("アクティブユーザー: %d人\n", activeCount)
}

func createExample() {
	fmt.Println("3. Go構造体からJSONの生成")

	newPerson := Person{
		Name:  "山田次郎",
		Age:   28,
		Email: "yamada@example.com",
		Hobbies: []string{"ゲーム", "アニメ", "プログラミング"},
		Address: Address{
			Prefecture: "神奈川県",
			City:       "横浜市",
			PostalCode: "220-0001",
		},
		IsActive: true,
	}

	jsonBytes, err := json.MarshalIndent(newPerson, "", "  ")
	if err != nil {
		log.Fatal("JSON生成エラー:", err)
	}

	fmt.Println("生成されたJSON:")
	fmt.Println(string(jsonBytes))

	compactJSON, err := json.Marshal(newPerson)
	if err != nil {
		log.Fatal("コンパクトJSON生成エラー:", err)
	}

	fmt.Println("\nコンパクト形式:")
	fmt.Println(string(compactJSON))

	fmt.Printf("\nJSONサイズ: %d バイト\n", len(compactJSON))
}
