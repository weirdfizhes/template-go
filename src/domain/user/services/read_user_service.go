package services

import (
	"context"
	"template-go/src/domain/user/models"
	"template-go/src/domain/user/repositories"
	"template-go/src/handlers"
)

func (s *UserService) GetAllUsers(ctx context.Context, paginate handlers.Pagination, search string) (count int64, user []models.GetUserPayload, err error) {
	count, err = repositories.GetCountUser(s.mainDB, search)
	if err != nil {
		return
	}

	user, err = repositories.GetAllUsers(s.mainDB, paginate, search)
	if err != nil {
		return
	}

	return
}

func (s *UserService) GetUser(ctx context.Context, id string) (user models.GetUserPayload, err error) {
	user, err = repositories.GetUserById(s.mainDB, id)
	if err != nil {
		return
	}

	return
}
