package middleware

import (
	"net/http"

	"azno-space.com/azno/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("token")

	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "missed token.! unauthorized"})
		return
	}

	userId, err := utils.VerifyToken(token)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "invalid token items"})
		return
	}
	context.Set("userId", userId)

}
