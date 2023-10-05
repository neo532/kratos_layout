package main

import (
	"context"
	"flag"
	"os"
	"time"

	"github.com/go-kratos/kratos/v2"
	klog "github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/neo532/apitool/transport"
	"github.com/neo532/apitool/transport/http/xhttp"
	"github.com/neo532/gokit/log"
	"github.com/neo532/gokit/middleware"
	"github.com/neo532/gokit/middleware/tracing"
	"github.com/neo532/gokit/server"

	"github.com/neo532/kratos_layout/cmd"
	"github.com/neo532/kratos_layout/internal/conf"
	srv "github.com/neo532/kratos_layout/internal/server"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	//Name string
	// Version is the version of the compiled software.
	Version string
	// flagconf is the config flag.
	flagConf string
)

func init() {
	flag.StringVar(&flagConf, "conf", "./configs/config.yaml", "config path, eg: -conf config.yaml")
}

func newApp(
	c context.Context,
	bs *conf.Bootstrap,
	hs *http.Server,
	r *srv.Router,
	logger klog.Logger) *kratos.App {
	return kratos.New(
		kratos.ID(bs.General.Ip),
		kratos.Name(bs.General.Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Context(c),
		kratos.Server(
			hs,
		),
	)
}

func main() {
	flag.Parse()

	// config
	bs, err := cmd.InitConfig(flagConf)
	if err != nil {
		panic(err)
	}
	bs.General.Version = Version

	// log
	zapLogger := cmd.InitLogger(bs, middleware.EntryApi, bs.General.Logger.Filename)
	defer func() {
		zapLogger.Sync()
		time.Sleep(3 * time.Second)
	}()
	logger := log.AddGlobalVariable(zapLogger)

	// group
	tracing.SetGroupForTracing(bs.General.Group)
	tracing.SetNameForTracing(bs.General.Name)

	envEnv, er := transport.String2Env(bs.General.Env)
	if er != nil {
		panic(er)
	}

	// app
	app, cleanup, err := initApp(
		cmd.BootContext(),
		bs,
		xhttp.New(
			xhttp.WithLogger(log.NewXHttpLogger(logger)),
			xhttp.WithEnv(envEnv),
		),
		logger,
	)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// pid
	// prestop: ["cat","/home/www/be_activity/pid","|","xargs","kill","-2"]
	if err := server.WritePID(os.Getpid(), ""); err != nil {
		panic(err)
	}

	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}
}
