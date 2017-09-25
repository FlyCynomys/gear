package gracefull

// build 1.8

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Event struct {
	Mask    uint32
	Message string
}

func (e *Event) String() string {
	return e.Message
}

func (e *Event) Signal() {

}

type GraceServer struct {
	http.Server
	OperationTimeout int64 // s
	ShutEvent        chan os.Signal
}

func NewGraceServer_(s http.Server, ShutTimeout int64) (*GraceServer, error) {
	g := new(GraceServer)
	g.OperationTimeout = ShutTimeout
	g.ShutEvent = make(chan os.Signal, 1)
	g.Server = s
	return g, nil
}

func (g *GraceServer) ListenAndServe() {
	go func() {
		if err := g.Server.ListenAndServe(); err != nil {
			log.Println("start server failed ", err)
			g.ShutEvent <- &Event{
				Mask:    10,
				Message: "server start failed",
			}
		}
	}()
	g.Wait()
	g.Shutdown()
}

func (g *GraceServer) Wait() {
	go g.SystemEvent()
	select {
	case e := <-g.ShutEvent:
		log.Println("shut down event : ", e.String())
	}
	err := g.Shutdown()
	if err != nil {
		log.Println("shotdown error : ", err.Error())
	} else {
		log.Println("shotdown ok ")
	}
}

func (g *GraceServer) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(g.OperationTimeout))
	defer cancel()

	err := g.Server.Shutdown(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (g *GraceServer) ProgramEvent(e *Event) {
	g.ShutEvent <- e
}

func (g *GraceServer) SystemEvent() {
	sc := make(chan os.Signal, 1)
	signal.Notify(sc,
		os.Interrupt,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)
	g.ShutEvent <- <-sc
}
