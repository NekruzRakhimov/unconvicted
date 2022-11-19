package service

import (
	"errors"
	"github.com/NekruzRakhimov/unconvicted/models"
	"github.com/NekruzRakhimov/unconvicted/pkg/repository"
	"github.com/NekruzRakhimov/unconvicted/utils"
)

func GetAllAdmins() (a []models.User, err error) {
	return repository.GetAllAdmins()
}

func CreateNewAdmin(admin models.User) (err error) {
	u, err := repository.GetUserByEmail(admin.Email)
	if err != nil {
		return err
	}

	if u.ID != 0 {
		return errors.New("пользователь с этим email'ом уже существует")
	}

	admin.Password = utils.GenerateHash(admin.Password)
	return repository.CreateNewAdmin(admin)
}

func DeleteAdmin(id int) error {
	return repository.DeleteAdmin(id)
}
