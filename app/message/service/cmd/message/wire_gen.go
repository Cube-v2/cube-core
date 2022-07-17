// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/contrib/registry/nacos/v2"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/the-zion/matrix-core/app/message/service/internal/biz"
	"github.com/the-zion/matrix-core/app/message/service/internal/conf"
	"github.com/the-zion/matrix-core/app/message/service/internal/data"
	"github.com/the-zion/matrix-core/app/message/service/internal/server"
	"github.com/the-zion/matrix-core/app/message/service/internal/service"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, confData *conf.Data, logger log.Logger, registry *nacos.Registry) (*kratos.App, func(), error) {
	userClient := data.NewUserServiceClient(registry, logger)
	creationClient := data.NewCreationServiceClient(registry, logger)
	achievementClient := data.NewAchievementServiceClient(registry, logger)
	cosUser := data.NewCosUserClient(confData, logger)
	cosCreation := data.NewCosCreationClient(confData, logger)
	txCode := data.NewPhoneCode(confData)
	goMail := data.NewGoMail(confData)
	dataData, err := data.NewData(logger, userClient, creationClient, achievementClient, cosUser, cosCreation, txCode, goMail)
	if err != nil {
		return nil, nil, err
	}
	userRepo := data.NewUserRepo(dataData, logger)
	userUseCase := biz.NewUserUseCase(userRepo, logger)
	creationRepo := data.NewCreationRepo(dataData, logger)
	creationUseCase := biz.NewCreationUseCase(creationRepo, logger)
	achievementRepo := data.NewAchievementRepo(dataData, logger)
	achievementCase := biz.NewAchievementUseCase(achievementRepo, logger)
	messageService := service.NewMessageService(userUseCase, creationUseCase, achievementCase, logger)
	httpServer := server.NewHTTPServer(confServer, messageService, logger)
	grpcServer := server.NewGRPCServer(confServer, messageService, logger)
	codeMqConsumerServer := server.NewCodeMqConsumerServer(confServer, messageService, logger)
	profileMqConsumerServer := server.NewProfileMqConsumerServer(confServer, messageService, logger)
	articleReviewMqConsumerServer := server.NewArticleReviewMqConsumerServer(confServer, messageService, logger)
	articleMqConsumerServer := server.NewArticleMqConsumerServer(confServer, messageService, logger)
	talkReviewMqConsumerServer := server.NewTalkReviewMqConsumerServer(confServer, messageService, logger)
	talkMqConsumerServer := server.NewTalkMqConsumerServer(confServer, messageService, logger)
	achievementMqConsumerServer := server.NewAchievementMqConsumerServer(confServer, messageService, logger)
	app := newApp(logger, registry, httpServer, grpcServer, codeMqConsumerServer, profileMqConsumerServer, articleReviewMqConsumerServer, articleMqConsumerServer, talkReviewMqConsumerServer, talkMqConsumerServer, achievementMqConsumerServer)
	return app, func() {
	}, nil
}
