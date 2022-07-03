// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/contrib/registry/nacos/v2"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/the-zion/matrix-core/app/creation/service/internal/biz"
	"github.com/the-zion/matrix-core/app/creation/service/internal/conf"
	"github.com/the-zion/matrix-core/app/creation/service/internal/data"
	"github.com/the-zion/matrix-core/app/creation/service/internal/server"
	"github.com/the-zion/matrix-core/app/creation/service/internal/service"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, confData *conf.Data, logger log.Logger, registry *nacos.Registry) (*kratos.App, func(), error) {
	db := data.NewDB(confData, logger)
	cmdable := data.NewRedis(confData, logger)
	client := data.NewCosServiceClient(confData, logger)
	articleDraftMqPro := data.NewRocketmqArticleDraftProducer(confData, logger)
	dataData, cleanup, err := data.NewData(db, cmdable, client, articleDraftMqPro, logger)
	if err != nil {
		return nil, nil, err
	}
	articleRepo := data.NewArticleRepo(dataData, logger)
	transaction := data.NewTransaction(dataData)
	articleUseCase := biz.NewArticleUseCase(articleRepo, transaction, logger)
	creationService := service.NewCreationService(articleUseCase, logger)
	httpServer := server.NewHTTPServer(confServer, creationService, logger)
	grpcServer := server.NewGRPCServer(confServer, creationService, logger)
	app := newApp(logger, registry, httpServer, grpcServer)
	return app, func() {
		cleanup()
	}, nil
}
