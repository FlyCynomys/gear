package main

import (
	"os"
	"os/signal"
	"syscall"

	"flag"

	"github.com/FlyCynomys/gear/conf"
	"github.com/FlyCynomys/gear/handles"
	"github.com/FlyCynomys/tools/log"
)

var cfgpath = flag.String("-c", "cfg.josn", "config file path")

func Init() {
	flag.Parse()
	log.Info("start server")
	log.Info(*cfgpath)
	conf.Init(*cfgpath)
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
