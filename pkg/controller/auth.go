package controller

import (
	"github.com/NekruzRakhimov/unconvicted/logger"
	"github.com/NekruzRakhimov/unconvicted/models"
	"github.com/NekruzRakhimov/unconvicted/pkg/service"
	"github.com/NekruzRakhimov/unconvicted/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SignUp(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		logger.Error.Printf("[%s] Error is: %s\n", utils.FuncName(), err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"reason": "ошибка во время парсинга данных"})
		return
	}

	if err := service.CreateUser(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	token, err := service.GenerateToken(user.Email, user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"access_token": token})
}

func SignIn(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		logger.Error.Printf("[%s] Error is: %s\n", utils.FuncName(), err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"reason": "ошибка во время парсинга данных"})
		return
	}

	token, err := service.GenerateToken(user.Email, user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"access_token": token})
}
