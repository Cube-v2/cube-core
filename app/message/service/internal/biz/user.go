package biz

import (
	"context"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/go-kratos/kratos/v2/log"
)

type UserRepo interface {
	SendCode(msgs ...*primitive.MessageExt)
	UploadProfileToCos(msgs ...*primitive.MessageExt)
}

type UserUseCase struct {
	repo UserRepo
	log  *log.Helper
}

func NewUserUseCase(repo UserRepo, logger log.Logger) *UserUseCase {
	return &UserUseCase{
		repo: repo,
		log:  log.NewHelper(log.With(logger, "module", "message/biz/userUseCase")),
	}
}

func (r *UserUseCase) SendCode(msgs ...*primitive.MessageExt) {
	r.repo.SendCode(msgs...)
}

func (r *UserUseCase) UploadProfileToCos(msgs ...*primitive.MessageExt) {
	r.repo.UploadProfileToCos(msgs...)
}

func (r *UserUseCase) AvatarReview(ctx context.Context, ar *AvatarReview) error {
	fmt.Println(ar)
	return nil
}

func (r *UserUseCase) ProfileReview(ctx context.Context, tr *TextReview) error {
	if tr.State != "Success" {
		r.log.Info("profile Review failed，%v", tr)
	}
	return nil
}
