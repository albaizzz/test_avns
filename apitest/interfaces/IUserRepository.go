package interfaces

import (
	"test_avns/apitest/models"
	"context"
)

type IUserRepository interface {
	ICrud
	GetUserByUsername(ctx context.Context, username string) (user models.User, err error)
	GetUsers(ctx context.Context) (users []models.User, err error)
}
