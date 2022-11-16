package repository

import (
	"github.com/NekruzRakhimov/unconvicted/db"
	"github.com/NekruzRakhimov/unconvicted/models"
)

func GetAllAdmins() (a []models.User, err error) {
	sqlQuery := "SELECT * FROM users WHERE is_admin = true"
	if err = db.GetDBConn().Raw(sqlQuery).Scan(&a).Error; err != nil {
		return nil, err
	}

	return
}

func CreateNewAdmin(u models.User) (err error) {
	u.IsSuperAdmin = false
	u.IsAdmin = true
	if err = db.GetDBConn().Table("users").Create(&u).Error; err != nil {
		return err
	}

	return nil
}
