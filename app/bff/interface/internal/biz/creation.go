package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type ArticleRepo interface {
	CreateArticleDraft(ctx context.Context, uuid string) (int32, error)
	GetLastArticleDraft(ctx context.Context, uuid string) (*ArticleDraft, error)
	ArticleDraftMark(ctx context.Context, uuid string, id int32) error
	GetArticleDraftList(ctx context.Context, uuid string) ([]*ArticleDraft, error)
	SendArticle(ctx context.Context, uuid string, id int32) error
}

type ArticleUseCase struct {
	repo ArticleRepo
	log  *log.Helper
}

func NewArticleUseCase(repo ArticleRepo, logger log.Logger) *ArticleUseCase {
	return &ArticleUseCase{
		repo: repo,
		log:  log.NewHelper(log.With(logger, "module", "bff/biz/ArticleUseCase")),
	}
}

func (r *ArticleUseCase) GetLastArticleDraft(ctx context.Context) (*ArticleDraft, error) {
	uuid := ctx.Value("uuid").(string)
	return r.repo.GetLastArticleDraft(ctx, uuid)
}

func (r *ArticleUseCase) CreateArticleDraft(ctx context.Context) (int32, error) {
	uuid := ctx.Value("uuid").(string)
	return r.repo.CreateArticleDraft(ctx, uuid)
}

func (r *ArticleUseCase) ArticleDraftMark(ctx context.Context, id int32) error {
	uuid := ctx.Value("uuid").(string)
	return r.repo.ArticleDraftMark(ctx, uuid, id)
}

func (r *ArticleUseCase) GetArticleDraftList(ctx context.Context) ([]*ArticleDraft, error) {
	uuid := ctx.Value("uuid").(string)
	return r.repo.GetArticleDraftList(ctx, uuid)
}

func (r *ArticleUseCase) SendArticle(ctx context.Context, id int32) error {
	uuid := ctx.Value("uuid").(string)
	return r.repo.SendArticle(ctx, uuid, id)
}
