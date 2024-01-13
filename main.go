package main

import (
	"github.com/gin-gonic/gin"

	"code.local/test/handlers"
)

func main() {
	router := gin.Default()

	router.POST("/sign", handlers.SignAnswers)
	router.POST("/verify", handlers.VerifySignature)

	router.Run(":8080")
}
