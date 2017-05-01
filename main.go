package main

import (
	"os"
	"os/signal"
	"syscall"

	"flag"

	"github.com/FlyCynomys/gear/conf"
	"github.com/FlyCynomys/gear/functions"
	"github.com/FlyCynomys/gear/handles"
	"github.com/FlyCynomys/gear/models"
	"github.com/FlyCynomys/gear/service"
	"github.com/FlyCynomys/tools/log"
)

var cfgpath = flag.String("c", "cfg.json", "config file path")

func Init() {
	flag.Parse()
	log.Info("start server")
	conf.Init(*cfgpath)
	service.Init()
	models.Init(conf.GetCfg())
	functions.Init()
	handles.Init("9000")

}

func main() {
	Init()
	sc := make(chan os.Signal, 1)
	signal.Notify(sc,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)
	<-sc
}
