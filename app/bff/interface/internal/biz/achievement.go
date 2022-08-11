package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type AchievementRepo interface {
	GetAchievementList(ctx context.Context, uuids []string) ([]*Achievement, error)
	GetUserAchievement(ctx context.Context, uuid string) (*Achievement, error)
}

type AchievementUseCase struct {
	repo AchievementRepo
	re   Recovery
	log  *log.Helper
}

func NewAchievementUseCase(repo AchievementRepo, re Recovery, logger log.Logger) *AchievementUseCase {
	return &AchievementUseCase{
		repo: repo,
		re:   re,
		log:  log.NewHelper(log.With(logger, "module", "bff/biz/AchievementUseCase")),
	}
}

func (r *AchievementUseCase) GetAchievementList(ctx context.Context, uuids []string) ([]*Achievement, error) {
	return r.repo.GetAchievementList(ctx, uuids)
}

func (r *AchievementUseCase) GetUserAchievement(ctx context.Context, uuid string) (*Achievement, error) {
	return r.repo.GetUserAchievement(ctx, uuid)
}
