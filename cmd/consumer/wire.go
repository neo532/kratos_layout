//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"context"

	"github.com/google/wire"

	kratos "github.com/go-kratos/kratos/v2"
	klog "github.com/go-kratos/kratos/v2/log"

	"github.com/neo532/kratos_layout/internal/biz"
	"github.com/neo532/kratos_layout/internal/conf"
	"github.com/neo532/kratos_layout/internal/data"
	"github.com/neo532/kratos_layout/internal/server"
	"github.com/neo532/kratos_layout/internal/service/consumer"
)

// initApp init kratos application.
func initApp(
	context.Context,
	*conf.Bootstrap,
	klog.Logger,
) (*kratos.App, func(), error) {
	panic(wire.Build(
		server.NewConsumerDefault,
		newApp,
		consumer.ProviderSet,

		biz.ProviderSet,
		data.ProviderSet,
	))
}
