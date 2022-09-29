package controller

import (
	"github.com/NekruzRakhimov/unconvicted/pkg/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetMe(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"reason": err.Error()})
		return
	}

	user, err := service.GetUserByID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}
