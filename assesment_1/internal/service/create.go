package service

import (
	"context"

	"assesment_1/internal/model"
)

func (s *service) Create(ctx context.Context, msg model.Message) error {
	s.cache.LPush(msg.FileID, msg)

	return nil
}
