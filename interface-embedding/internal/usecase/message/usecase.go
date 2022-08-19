package message

import (
	"github.com/MeteorSis/test-golang/interface-embedding/internal/usecase"
)

type Usecase struct {
	usecase.MessageUsecase

	content string
}

func New(content string) *Usecase {
	return &Usecase{content: content}
}

/* func (u *Usecase) GetMessage(ctx context.Context, cID string, mID int64) (*domain.Message, error) {
	return &domain.Message{ChannelID: cID, ID: mID, Content: u.content}, nil
} */
