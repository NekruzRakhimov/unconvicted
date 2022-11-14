package repository

import (
	"errors"
	"fmt"
	"github.com/NekruzRakhimov/unconvicted/db"
	"github.com/NekruzRakhimov/unconvicted/logger"
	"github.com/NekruzRakhimov/unconvicted/models"
	"github.com/NekruzRakhimov/unconvicted/utils"
)

func CreateReference(r models.Reference) error {
	if err := db.GetDBConn().Table("references").Create(&r).Error; err != nil {
		logger.Error.Printf("[%s] Error is: %s\n", utils.FuncName(), err.Error())
		return errors.New("ошибка во время создания записи")
	}

	return nil
}

func GetMyReferences(email string) (r []models.Reference, err error) {
	fmt.Println(email)
	sqlQuery := "SELECT * FROM \"references\" WHERE email = ?"
	if err = db.GetDBConn().Raw(sqlQuery, email).Scan(&r).Error; err != nil {
		logger.Error.Printf("[%s] Error is: %s\n", utils.FuncName(), err.Error())
		return nil, errors.New("ошибка во время получения данных")
	}

	return r, nil
}

func GetReferenceByID(id int) (r models.Reference, err error) {
	sqlQuery := "SELECT * FROM \"references\" WHERE id = ?"
	if err = db.GetDBConn().Raw(sqlQuery, id).Scan(&r).Error; err != nil {
		logger.Error.Printf("[%s] Error is: %s\n", utils.FuncName(), err.Error())
		return models.Reference{}, errors.New("ошибка во время получения данных")
	}

	return r, nil
}

func ChangeReferenceStatus(id int, status string) error {
	sqlQuery := "UPDATE \"references\" set status = ? WHERE id = ?"
	if err := db.GetDBConn().Exec(sqlQuery, status, id).Error; err != nil {
		logger.Error.Printf("[%s] Error is: %s\n", utils.FuncName(), err.Error())
		return errors.New("ошибка во время получения данных")
	}

	return nil
}
