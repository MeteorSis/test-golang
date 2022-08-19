package usecase

import (
	"context"

	"github.com/MeteorSis/test-golang/interface-embedding/internal/domain"
)

type MessageUsecase interface {
	GetMessage(ctx context.Context, cID string, mID int64) (*domain.Message, error)
}
