// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"context"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/neo532/apitool/transport/http/xhttp/client"
	"github.com/neo532/kratos_layout/cmd"
	"github.com/neo532/kratos_layout/internal/biz"
	"github.com/neo532/kratos_layout/internal/conf"
	"github.com/neo532/kratos_layout/internal/data"
	"github.com/neo532/kratos_layout/internal/server"
	"github.com/neo532/kratos_layout/internal/service/api"
)

// Injectors from wire.go:

// initApp init kratos application.
func initApp(contextContext context.Context, bootstrap *conf.Bootstrap, clientClient client.Client, logger log.Logger) (*kratos.App, func(), error) {
	httpServer := server.NewHTTPServer(bootstrap, logger)
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
	demoUsecase := biz.NewDemoUsecase(transactionDefaultRepo, distributedLock, demoRepo)
	demoApi := api.NewDemoApi(demoUsecase, logger)
	router := server.InitHTTPRouter(httpServer, demoApi)
	app := newApp(contextContext, bootstrap, httpServer, router, logger)
	return app, func() {
		cleanup2()
		cleanup()
	}, nil
}

func InitDemo() (*biz.DemoUsecase, func(), error) {
	contextContext := cmd.BootContext()
	bootstrap := cmd.ConfBootstap()
	logger := cmd.ConfLogger(bootstrap)
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
	demoUsecase := biz.NewDemoUsecase(transactionDefaultRepo, distributedLock, demoRepo)
	return demoUsecase, func() {
		cleanup2()
		cleanup()
	}, nil
}
