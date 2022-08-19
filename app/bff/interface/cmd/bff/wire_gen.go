// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/contrib/registry/nacos/v2"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/the-zion/matrix-core/app/bff/interface/internal/biz"
	"github.com/the-zion/matrix-core/app/bff/interface/internal/conf"
	"github.com/the-zion/matrix-core/app/bff/interface/internal/data"
	"github.com/the-zion/matrix-core/app/bff/interface/internal/server"
	"github.com/the-zion/matrix-core/app/bff/interface/internal/service"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, confData *conf.Data, logger log.Logger, registry *nacos.Registry) (*kratos.App, func(), error) {
	userClient := data.NewUserServiceClient(registry, logger)
	creationClient := data.NewCreationServiceClient(registry, logger)
	messageClient := data.NewMessageServiceClient(registry, logger)
	achievementClient := data.NewAchievementServiceClient(registry, logger)
	commentClient := data.NewCommentServiceClient(registry, logger)
	dataData, err := data.NewData(logger, userClient, creationClient, messageClient, achievementClient, commentClient)
	if err != nil {
		return nil, nil, err
	}
	userRepo := data.NewUserRepo(dataData, logger)
	achievementRepo := data.NewAchievementRepo(dataData, logger)
	creationRepo := data.NewCreationRepo(dataData, logger)
	recovery := data.NewRecovery(dataData)
	userUseCase := biz.NewUserUseCase(userRepo, achievementRepo, creationRepo, recovery, logger)
	creationUseCase := biz.NewCreationUseCase(creationRepo, recovery, logger)
	talkRepo := data.NewTalkRepo(dataData, logger)
	talkUseCase := biz.NewTalkUseCase(talkRepo, recovery, logger)
	articleRepo := data.NewArticleRepo(dataData, logger)
	articleUseCase := biz.NewArticleUseCase(articleRepo, recovery, logger)
	columnRepo := data.NewColumnRepo(dataData, logger)
	columnUseCase := biz.NewColumnUseCase(columnRepo, recovery, logger)
	commentRepo := data.NewCommentRepo(dataData, logger)
	achievementUseCase := biz.NewAchievementUseCase(achievementRepo, creationRepo, commentRepo, recovery, logger)
	newsRepo := data.NewNewsRepo(dataData, logger)
	newsUseCase := biz.NewNewsUseCase(newsRepo, recovery, logger)
	commentUseCase := biz.NewCommentUseCase(commentRepo, recovery, logger)
	bffService := service.NewBffService(userUseCase, creationUseCase, talkUseCase, articleUseCase, columnUseCase, achievementUseCase, newsUseCase, commentUseCase, logger)
	httpServer := server.NewHTTPServer(confServer, bffService, logger)
	grpcServer := server.NewGRPCServer(confServer, bffService, logger)
	app := newApp(logger, registry, httpServer, grpcServer)
	return app, func() {
	}, nil
}
