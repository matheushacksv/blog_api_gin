package main

import (
	"api/internal/config"
	"api/internal/database"
	"api/internal/post"
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg, erro := config.Load()
	if erro != nil {
		log.Fatal(erro)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pool, erro := database.Connect(ctx, cfg.DatabaseURL)
	if erro != nil {
		log.Fatal(erro)
	}
	defer pool.Close()

	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		erro := pool.Ping(c.Request.Context())
		if erro != nil {
			c.JSON(http.StatusServiceUnavailable, gin.H{
				"status": "down",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	postHandler := post.NewPostHandler(pool)

	router.POST("/posts", postHandler.Create)
	router.GET("/posts", postHandler.Search)
	router.GET("/posts/:ID", postHandler.Get)
	router.PUT("/posts/:ID", postHandler.Update)
	router.DELETE("/posts/:ID", postHandler.Delete)

	if erro := router.Run(
		fmt.Sprintf(":%s", cfg.Port),
	); erro != nil {
		log.Fatal("API error")
	}
}
