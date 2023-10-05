// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"context"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/neo532/kratos_layout/internal/biz"
	"github.com/neo532/kratos_layout/internal/conf"
	"github.com/neo532/kratos_layout/internal/data"
	"github.com/neo532/kratos_layout/internal/server"
	"github.com/neo532/kratos_layout/internal/service/consumer"
)

// Injectors from wire.go:

// initApp init kratos application.
func initApp(contextContext context.Context, bootstrap *conf.Bootstrap, logger log.Logger) (*kratos.App, func(), error) {
	databaseDefault, cleanup, err := data.NewDatabaseDefault(contextContext, bootstrap, logger)
	if err != nil {
		return nil, nil, err
	}
	transactionDefaultRepo := data.NewTransactionDefaultRepo(databaseDefault, logger)
	redisLock, cleanup2, err := data.NewRedisLock(contextContext, bootstrap, logger)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	distributedLock := data.NewToolDistributedLock(redisLock)
	demoRepo := data.NewDemoRepo(databaseDefault)
	demoUsecase := biz.NewDemoUsecase(transactionDefaultRepo, distributedLock, demoRepo, logger)
	demoConsumer := consumer.NewDemoConsumer(demoUsecase, logger)
	queueConsumer := server.NewConsumerDefault(contextContext, bootstrap, logger, demoConsumer)
	app := newApp(contextContext, bootstrap, logger, queueConsumer)
	return app, func() {
		cleanup2()
		cleanup()
	}, nil
}
