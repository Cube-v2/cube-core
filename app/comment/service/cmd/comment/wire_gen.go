// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/contrib/registry/nacos/v2"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/the-zion/matrix-core/app/comment/service/internal/biz"
	"github.com/the-zion/matrix-core/app/comment/service/internal/conf"
	"github.com/the-zion/matrix-core/app/comment/service/internal/data"
	"github.com/the-zion/matrix-core/app/comment/service/internal/server"
	"github.com/the-zion/matrix-core/app/comment/service/internal/service"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, confData *conf.Data, logger log.Logger, registry *nacos.Registry) (*kratos.App, func(), error) {
	db := data.NewDB(confData, logger)
	client := data.NewCosServiceClient(confData, logger)
	cmdable := data.NewRedis(confData, logger)
	reviewMqPro := data.NewRocketmqReviewProducer(confData, logger)
	commentMqPro := data.NewRocketmqCommentProducer(confData, logger)
	dataData, cleanup, err := data.NewData(db, client, cmdable, reviewMqPro, commentMqPro, logger)
	if err != nil {
		return nil, nil, err
	}
	commentRepo := data.NewCommentRepo(dataData, logger)
	transaction := data.NewTransaction(dataData)
	commentUseCase := biz.NewCommentUseCase(commentRepo, transaction, logger)
	commentService := service.NewCommentService(commentUseCase, logger)
	httpServer := server.NewHTTPServer(confServer, commentService, logger)
	grpcServer := server.NewGRPCServer(confServer, commentService, logger)
	app := newApp(logger, registry, httpServer, grpcServer)
	return app, func() {
		cleanup()
	}, nil
}
