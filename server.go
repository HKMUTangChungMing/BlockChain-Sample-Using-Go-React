package main

import (
	"myproject/blockchain" // 引入 blockchain package
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 初始化區塊鏈
	blockchain.InitBlockchain()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},  // 允許 React 的來源
		AllowMethods:     []string{"GET", "POST", "OPTIONS"}, // 允許的 HTTP 方法
		AllowHeaders:     []string{"Origin", "Content-Type"}, // 允許的 Header
		AllowCredentials: true,
	}))

	r.GET("/blocks", func(c *gin.Context) {
		c.JSON(http.StatusOK, blockchain.Blockchain)
	})

	r.POST("/deposit", func(c *gin.Context) {
		type Request struct {
			Data string `json:"data"`
		}

		var req Request
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		lastBlock := blockchain.Blockchain[len(blockchain.Blockchain)-1]
		newBlock := blockchain.GenerateBlock(lastBlock, req.Data)
		blockchain.Blockchain = append(blockchain.Blockchain, newBlock)

		c.JSON(http.StatusOK, newBlock)
	})

	r.Run(":8080")
}
