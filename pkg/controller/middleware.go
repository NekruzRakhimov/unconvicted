package controller

import (
	"errors"
	"github.com/NekruzRakhimov/unconvicted/logger"
	"github.com/NekruzRakhimov/unconvicted/pkg/service"
	"github.com/NekruzRakhimov/unconvicted/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userID"
)

func UserIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		logger.Error.Printf("[%s] Error is: %s\n", utils.FuncName(), "empty auth header")
		c.JSON(http.StatusUnauthorized, gin.H{"reason": "empty auth header"})
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		logger.Error.Printf("[%s] Error is: %s\n", utils.FuncName(), "empty auth header")
		c.JSON(http.StatusUnauthorized, gin.H{"reason": "invalid auth header"})
		return
	}

	userId, err := service.ParseToken(headerParts[1])
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"reason": err.Error()})
		return
	}

	c.Set(userCtx, userId)
}

func getUserID(c *gin.Context) (int, error) {
	id, ok := c.Get(userCtx)
	if !ok {
		logger.Error.Printf("[%s] Error is: %s\n", utils.FuncName(), "user id not found")
		return 0, errors.New("user id not found")
	}

	idInt, ok := id.(int)
	if !ok {
		logger.Error.Printf("[%s] Error is: %s\n", utils.FuncName(), "user id is of invalid type")
		return 0, errors.New("user id is of invalid type")
	}
	return idInt, nil
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET,PUT,POST,DELETE,PATCH,OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, auth")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		}

		c.Next()
	}
}
