package bot

import (
	"context"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"

	"tg.bot/telegram"
)

type Token string

type Service struct {
	Logger log.Logger
	Token  Token
}

func NewService(logger log.Logger, token Token) *Service {
	return &Service{
		Logger: logger,
		Token:  token,
	}
}

func (s *Service) Update(ctx context.Context, updateReq *telegram.Update) error {
	level.Info(s.Logger).Log("msg", "received update request", "username", updateReq.Message.From.UserName)
	return nil
}