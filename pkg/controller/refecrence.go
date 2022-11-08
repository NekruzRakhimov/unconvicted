package controller

import (
	"encoding/json"
	"github.com/NekruzRakhimov/unconvicted/logger"
	"github.com/NekruzRakhimov/unconvicted/models"
	"github.com/NekruzRakhimov/unconvicted/pkg/service"
	"github.com/NekruzRakhimov/unconvicted/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateReference(c *gin.Context) {
	var (
		err       error
		reference models.Reference
	)
	str, ok := c.GetPostForm("json")
	if !ok {
		logger.Error.Printf("[%s] Error is: %s\n", utils.FuncName(), "json field not found")
		c.JSON(http.StatusBadRequest, gin.H{"reason": "json field not found"})
		return
	}

	if err = json.Unmarshal([]byte(str), &reference); err != nil {
		logger.Error.Printf("[%s] Error is: %s\n", utils.FuncName(), err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"reason": "json field not found"})
		return
	}

	reference.PassportFront, err = SaveImage(c, "passport_front")
	if err != nil {
		logger.Error.Printf("[%s] Error is: %s\n", utils.FuncName(), err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"reason": "json field not found"})
		return
	}

	reference.PassportBack, err = SaveImage(c, "passport_back")
	if err != nil {
		logger.Error.Printf("[%s] Error is: %s\n", utils.FuncName(), err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"reason": "passport_back field not found"})
		return
	}

	reference.PassportSelfie, err = SaveImage(c, "passport_selfie")
	if err != nil {
		logger.Error.Printf("[%s] Error is: %s\n", utils.FuncName(), err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"reason": "passport_selfie field not found"})
		return
	}

	reference.PassportBack, err = SaveImage(c, "payment_receipt")
	if err != nil {
		logger.Error.Printf("[%s] Error is: %s\n", utils.FuncName(), err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"reason": "payment_receipt field not found"})
		return
	}

	if err = service.CreateReference(reference); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"reason": "Запись успешно создана"})
}

func GetMyReferences(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"reason": err.Error()})
		return
	}

	logger.Debug.Println(userID)

	r, err := service.GetMyReference(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	c.JSON(http.StatusOK, r)
}
