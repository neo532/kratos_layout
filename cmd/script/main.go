package main

import (
	"context"
	"flag"
	"time"

	"github.com/go-kratos/kratos/v2"
	klog "github.com/go-kratos/kratos/v2/log"

	"github.com/neo532/gokit/log"
	"github.com/neo532/gokit/middleware"
	"github.com/neo532/gokit/middleware/tracing"

	"github.com/neo532/kratos_layout/cmd"
	"github.com/neo532/kratos_layout/internal/conf"
	"github.com/neo532/kratos_layout/internal/server"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	//Name string
	// Version is the version of the compiled software.
	Version string
	// flagConf is the config flag.
	flagConf string
	flagCmd  string
	flagArgs string
)

func init() {
	flag.StringVar(&flagConf, "conf", "../../configs", "config path, eg: -conf config.yaml")
	flag.StringVar(&flagCmd, "cmd", "", "command, eg: -cmd config.yaml")
	flag.StringVar(&flagArgs, "args", "", `args, eg: -args '{"a":1}'`)
}

func newApp(c context.Context, bs *conf.Bootstrap, logger klog.Logger, script *server.Script) *kratos.App {
	return kratos.New(
		kratos.ID(bs.General.Ip),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Context(c),
		kratos.Server(
			script,
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

	// log
	zapLogger := cmd.InitLogger(bs, middleware.EntryScript, bs.General.Logger.FilenameScript)
	defer func() {
		zapLogger.Sync()
		time.Sleep(3 * time.Second)
	}()
	logger := log.AddGlobalVariable(zapLogger)

	// discover && register
	tracing.SetGroupForTracing(bs.General.Group)
	tracing.SetNameForTracing(bs.General.Name)

	// stop event
	scriptFinishCh := make(chan struct{}, 1)

	// app init
	app, cleanup, err := initApp(
		cmd.BootContext(),
		scriptFinishCh,
		[]string{flagCmd, flagArgs},
		bs,
		logger,
	)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// listen stop
	go func() {
		<-scriptFinishCh
		app.Stop()
	}()

	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}
}
