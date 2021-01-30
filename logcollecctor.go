package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/Ak-Army/cli"
	"github.com/Ak-Army/cli/command"
	"github.com/Ak-Army/xlog"

	_ "github.com/Ak-Army/logcollector/cmd"
)

var logger xlog.Logger

func main() {
	initLogger()

	logger.SetField("version", Version)
	logger.SetField("pid", fmt.Sprintf("%d", os.Getpid()))
	logger.Info("start...")
	ctx := xlog.NewContext(context.Background(), logger)

	c := cli.New("log collector", Version)
	cli.RootCommand().Authors = []string{"Hunyi"}
	cli.RootCommand().AddCommand("completion", command.New("log collector"))
	c.Run(ctx, os.Args)
}

func initLogger() {
	conf := xlog.Config{
		Level:  xlog.LevelDebug,
		Output: xlog.NewConsoleOutput(),
	}
	log.SetFlags(0)
	logger = xlog.New(conf)
	xlog.SetLogger(logger)
	log.SetOutput(logger)
}
