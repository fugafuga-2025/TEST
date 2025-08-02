package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var users = []User{
	{ID: 1, Name: "田中太郎", Age: 25},
	{ID: 2, Name: "佐藤花子", Age: 30},
	{ID: 3, Name: "鈴木一郎", Age: 28},
}

func main() {
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"message": "サーバーは正常に動作しています",
		})
	})

	r.GET("/users", getUsers)

	r.GET("/users/:id", getUserByID)

	r.GET("/api/info", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"api_name":    "Gin Docker Sample API",
			"version":     "1.0.0",
			"description": "Go GinとDockerを使ったサンプルREST API",
			"endpoints": []string{
				"GET /health - ヘルスチェック",
				"GET /users - 全ユーザー取得",
				"GET /users/:id - 特定ユーザー取得",
				"GET /api/info - API情報取得",
			},
		})
	})

	r.Run(":8080")
}

func getUsers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   users,
		"count":  len(users),
	})
}

func getUserByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "無効なユーザーIDです",
		})
		return
	}

	for _, user := range users {
		if user.ID == id {
			c.JSON(http.StatusOK, gin.H{
				"status": "success",
				"data":   user,
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"status":  "error",
		"message": "ユーザーが見つかりません",
	})
}
