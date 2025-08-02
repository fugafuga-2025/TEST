package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

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
	r := gin.New()

	r.Use(LoggerMiddleware())
	r.Use(CORSMiddleware())
	r.Use(gin.Recovery())

	api := r.Group("/api/v1")
	api.Use(AuthMiddleware())
	{
		api.GET("/users", getUsers)
		api.GET("/users/:id", getUserByID)
		api.GET("/protected", getProtectedData)
	}

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":    "ok",
			"message":   "ミドルウェアデモサーバーが動作中",
			"timestamp": time.Now().Format("2006-01-02 15:04:05"),
		})
	})

	r.GET("/public", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "これは認証不要の公開エンドポイントです",
			"data":    "誰でもアクセス可能",
		})
	})

	fmt.Println("ミドルウェアデモサーバーを起動中...")
	fmt.Println("利用可能なエンドポイント:")
	fmt.Println("  GET /health - ヘルスチェック（認証不要）")
	fmt.Println("  GET /public - 公開エンドポイント（認証不要）")
	fmt.Println("  GET /api/v1/users - ユーザー一覧（認証必要）")
	fmt.Println("  GET /api/v1/users/:id - 特定ユーザー（認証必要）")
	fmt.Println("  GET /api/v1/protected - 保護されたデータ（認証必要）")
	fmt.Println("")
	fmt.Println("認証が必要なエンドポイントには以下のヘッダーを追加してください:")
	fmt.Println("  Authorization: Bearer demo-token")
	fmt.Println("")

	log.Fatal(r.Run(":8080"))
}

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		c.Next()

		end := time.Now()
		latency := end.Sub(start)

		clientIP := c.ClientIP()
		method := c.Request.Method
		statusCode := c.Writer.Status()

		if raw != "" {
			path = path + "?" + raw
		}

		fmt.Printf("[GIN] %s | %3d | %13v | %15s | %-7s %s\n",
			end.Format("2006/01/02 - 15:04:05"),
			statusCode,
			latency,
			clientIP,
			method,
			path,
		)
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status":  "error",
				"message": "認証ヘッダーが必要です",
				"hint":    "Authorization: Bearer demo-token を追加してください",
			})
			c.Abort()
			return
		}

		if authHeader != "Bearer demo-token" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status":  "error",
				"message": "無効な認証トークンです",
				"hint":    "正しいトークン: Bearer demo-token",
			})
			c.Abort()
			return
		}

		c.Set("user_id", "demo-user")
		c.Next()
	}
}

func getUsers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   users,
		"count":  len(users),
		"message": "認証済みユーザーのみアクセス可能",
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
				"message": "認証済みユーザーのみアクセス可能",
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"status":  "error",
		"message": "ユーザーが見つかりません",
	})
}

func getProtectedData(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "ユーザー情報の取得に失敗しました",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "これは保護されたデータです",
		"data": gin.H{
			"user_id":     userID,
			"secret_data": "機密情報: Go Ginミドルウェアの実装例",
			"timestamp":   time.Now().Format("2006-01-02 15:04:05"),
			"permissions": []string{"read", "write", "admin"},
		},
	})
}
