package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"code.local/test/config"
	"code.local/test/models"
	"code.local/test/repositories"
	"code.local/test/services"
)

func VerifySignature(c *gin.Context) {
	var verifyReq models.Verify

	if err := c.ShouldBindJSON(&verifyReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request format"})
		return
	}

	ss := services.NewSignatureService(repositories.NewSignatureRepository(config.DB))
	defer ss.Close()

	data, isValid, err := ss.VerifySignature(verifyReq.UserID, verifyReq.Signature)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal signature verification error"})
		return
	}

	if !isValid {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": "invalid signature"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":   "OK",
		"answers":   data.Answers,
		"timestamp": data.Timestamp,
	})
}
