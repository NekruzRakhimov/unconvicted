package repository

import (
	"errors"
	"github.com/NekruzRakhimov/unconvicted/db"
	"github.com/NekruzRakhimov/unconvicted/logger"
	"github.com/NekruzRakhimov/unconvicted/models"
	"github.com/NekruzRakhimov/unconvicted/utils"
)

func GetUserByEmail(email string) (user models.User, err error) {
	sqlQuery := "SELECT * FROM users WHERE email = ?"
	if err := db.GetDBConn().Raw(sqlQuery, email).Scan(&user).Error; err != nil {
		return models.User{}, err
	}

	return
}

func GetUser(email, password string) (user models.User, err error) {
	sqlQuery := "SELECT * FROM users WHERE email = ? AND password = ?"
	if err := db.GetDBConn().Raw(sqlQuery, email, password).Scan(&user).Error; err != nil {
		logger.Error.Printf("[%s] Error is: %s\n", utils.FuncName(), err.Error())
		return models.User{}, errors.New("что-то пошло не так(")
	}

	if user.ID == 0 {
		return models.User{}, errors.New("неправильный логин или пароль")
	}

	return
}

func GetUserByID(id int) (user models.User, err error) {
	sqlQuery := "SELECT * FROM users WHERE id = ?"
	if err := db.GetDBConn().Raw(sqlQuery, id).Scan(&user).Error; err != nil {
		logger.Error.Printf("[%s] Error is: %s\n", utils.FuncName(), err.Error())
		return models.User{}, errors.New("что-то пошло не так(")
	}

	return
}

func CreateUser(user models.User) error {
	if err := db.GetDBConn().Table("users").Create(&user).Error; err != nil {
		logger.Error.Printf("[%s] Error is: %s\n", utils.FuncName(), err.Error())
		return errors.New("что-то пошло не так(")
	}

	return nil
}
