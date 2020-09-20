package store

import (
	"context"
	"github.com/vavilen84/class_booking/constants"
)

func GetDefaultDBContext() context.Context {
	parentCtx := context.Background()
	ctx, _ := context.WithTimeout(parentCtx, constants.DefaultStoreTimeout)
	return ctx
}
