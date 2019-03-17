package interfaces

import (
	"test_avns/apitest/models"
	"context"
)

type IAuthRepository interface {
	SaveSession(ctx context.Context, userSession models.UserSession) (err error)
	GetSession(ctx context.Context, session string) (userSession models.UserSession, err error)
	Logout(ctx context.Context, session string) (err error)
}
