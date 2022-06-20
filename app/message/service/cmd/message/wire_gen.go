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
	txCode := data.NewPhoneCode(confData)
	goMail := data.NewGoMail(confData)
	dataData, err := data.NewData(logger, userClient, txCode, goMail)
	if err != nil {
		return nil, nil, err
	}
	userRepo := data.NewUserRepo(dataData, logger)
	userUseCase := biz.NewUserUseCase(userRepo, logger)
	messageService := service.NewMessageService(userUseCase, logger)
	httpServer := server.NewHTTPServer(confServer, messageService, logger)
	grpcServer := server.NewGRPCServer(confServer, messageService, logger)
	codeMqConsumerServer := server.NewCodeMqConsumerServer(confServer, messageService, logger)
	app := newApp(logger, registry, httpServer, grpcServer, codeMqConsumerServer)
	return app, func() {
	}, nil
}
