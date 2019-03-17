package interfaces

import "context"

type ICrud interface {
	Add(ctx context.Context, data interface{}) error
	Edit(ctx context.Context, data interface{}) error
	Delete(ctx context.Context, id int) error
}
