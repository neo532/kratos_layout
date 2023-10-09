//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"context"

	kratos "github.com/go-kratos/kratos/v2"
	klog "github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/neo532/apitool/transport/http/xhttp/client"

	"github.com/neo532/kratos_layout/cmd"
	"github.com/neo532/kratos_layout/internal/biz"
	"github.com/neo532/kratos_layout/internal/conf"
	"github.com/neo532/kratos_layout/internal/data"
	"github.com/neo532/kratos_layout/internal/server"
	"github.com/neo532/kratos_layout/internal/service/api"
)

// initApp init kratos application.
func initApp(
	context.Context,
	*conf.Bootstrap,
	client.Client,
	klog.Logger,
) (*kratos.App, func(), error) {
	panic(wire.Build(
		server.NewHTTPServer,
		server.InitHTTPRouter,

		newApp,
		api.ProviderSet,

		biz.ProviderSet,
		data.ProviderSet,
	))
}

func InitDemo() (*biz.DemoUsecase, func(), error) {
	panic(wire.Build(
		cmd.InitUnitTestSet,

		biz.ProviderSet,
		data.ProviderSet,
	))
}
