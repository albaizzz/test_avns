package interfaces

import (
	"test_avns/apitest/models"
	"context"
)

type IUserService interface {
	EditUser(ctx context.Context, user models.User) (response models.APIResponse)
	DeleteUser(ctx context.Context, user models.User) (response models.APIResponse)
	AuthUser(ctx context.Context, user models.User) (response models.APIResponse)
	Register(ctx context.Context, user models.User) (response models.APIResponse)
	GetUserByUsername(ctx context.Context, username string) (response models.APIResponse)
	Logout(ctx context.Context, session string) (response models.APIResponse)
}
