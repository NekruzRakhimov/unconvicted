package service

import (
	"github.com/NekruzRakhimov/unconvicted/logger"
	"github.com/NekruzRakhimov/unconvicted/models"
	"github.com/NekruzRakhimov/unconvicted/pkg/repository"
)

func CreateReference(r models.Reference) error {
	return repository.CreateReference(r)
}

func GetMyReference(userID int) (r []models.Reference, err error) {
	u, err := repository.GetUserByID(userID)
	if err != nil {
		return nil, err
	}

	logger.Debug.Println(u)

	r, err = repository.GetMyReferences(u.Email)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func GetReferenceByID(id int) (r models.Reference, err error) {
	return repository.GetReferenceByID(id)
}

func ChangeReferenceStatus(id int, status string) error {
	return repository.ChangeReferenceStatus(id, status)
}
