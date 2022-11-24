package controller

import (
	"github.com/NekruzRakhimov/unconvicted/models"
	"github.com/NekruzRakhimov/unconvicted/pkg/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetAllAdmins(c *gin.Context) {
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

	if !user.IsSuperAdmin {
		c.JSON(http.StatusForbidden, gin.H{"reason": "у вас нет необходимых прав"})
		return
	}

	a, err := service.GetAllAdmins()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	c.JSON(http.StatusOK, a)
}

func CreateAdmin(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"reason": err.Error()})
		return
	}

	sAdmin, err := service.GetUserByID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	if !sAdmin.IsSuperAdmin {
		c.JSON(http.StatusForbidden, gin.H{"reason": "у вас нет необходимых прав"})
		return
	}

	var u models.User
	if err = c.BindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
		return
	}
	u.Password = "admin"

	if err = service.CreateNewAdmin(u); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"reason": "новый модератор успешно создан"})
}

func DeleteAdmin(c *gin.Context) {
	_, err := getUserID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"reason": err.Error()})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"reason": "id not found"})
		return
	}

	if err = service.DeleteAdmin(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"reason": "пользователь успешно удален"})
}

func GetAdminsActivity(c *gin.Context) {
	a := []models.AdminActivity{
		models.AdminActivity{
			ID:          1,
			FullName:    "Testov Test",
			Description: "Изменил статус заявки",
			CreatedAt:   "2022-11-15 17:45:11.080194 +00:00",
		},
		models.AdminActivity{
			ID:          2,
			FullName:    "Testov Test",
			Description: "Изменил статус заявки",
			CreatedAt:   "2022-11-15 17:45:11.080194 +00:00",
		},
		models.AdminActivity{
			ID:          3,
			FullName:    "Testov Test",
			Description: "Изменил статус заявки",
			CreatedAt:   "2022-11-15 17:45:11.080194 +00:00",
		},
	}

	c.JSON(http.StatusOK, a)
}
