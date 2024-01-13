package handlers

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"

	"code.local/test/config"
	"code.local/test/models"
	"code.local/test/repositories"
	"code.local/test/services"
)

var JWTSecretKey []byte

func init() {
	JWTSecretKey = []byte(os.Getenv(config.JWTSecretKey))
}

func SignAnswers(c *gin.Context) {
	claims := &jwt.StandardClaims{}

	token, err := jwt.ParseWithClaims(c.GetHeader("Authorization"), claims, func(token *jwt.Token) (interface{}, error) {
		return JWTSecretKey, nil
	})

	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
		return
	}

	if err := claims.Valid(); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
		return
	}

	var signReg models.Sign

	if err := c.ShouldBindJSON(&signReg); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	userID := claims.Subject

	ss := services.NewSignatureService(repositories.NewSignatureRepository(config.DB))
	defer ss.Close()

	signature, err := ss.SignAnswers(userID, signReg.Questions, signReg.Answers)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to sign answers"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"test_signature": signature,
	})
}
