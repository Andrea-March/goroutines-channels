package server

import (
	"context"
	"fmt"
	"github.com/config"
	"github.com/golang/glog"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

func Listen(cfg *config.Configuration, handler http.Handler)  {
	stopMain := make(chan os.Signal)
	stopSignals := make(chan os.Signal)
	signal.Notify(stopSignals, syscall.SIGTERM, syscall.SIGINT)
	done := make(chan struct{})
	mainServer := newMainServer(cfg, handler)
	mainListener, err := newListener(mainServer.Addr)
	if err != nil {
		glog.Errorf("Error listening for TCP connections on %s: %v for main server", mainServer.Addr, err)
		return
	}

	go shutdownAfterSignals(mainServer, stopMain, done)
	go runServer(mainServer, "Main", mainListener)
	wait(stopSignals, done, stopMain)
	return
}

func wait(inbound <-chan os.Signal, done <-chan struct{}, outbound ...chan<- os.Signal) {
	sig := <-inbound

	for i := 0; i < len(outbound); i++ {
		go sendSignal(outbound[i], sig)
	}

	for i := 0; i < len(outbound); i++ {
		<-done
	}
}

func sendSignal(to chan<- os.Signal, sig os.Signal) {
	to <- sig
}

func newMainServer(cfg *config.Configuration, handler http.Handler) *http.Server {
	var serverHandler = handler

	return &http.Server{
		Addr:         cfg.Host + ":" + strconv.Itoa(cfg.Port),
		Handler:      serverHandler,
	}

}


func newListener(address string) (net.Listener, error) {
	ln, err := net.Listen("tcp", address)
	if err != nil {
		return nil, fmt.Errorf("Error listening for TCP connections on %s: %v", address, err)
	}
	return ln, nil
}

func shutdownAfterSignals(server *http.Server, stopper <-chan os.Signal, done chan<- struct{}) {
	sig := <-stopper

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var s struct{}
	glog.Infof("Stopping %s because of signal: %s", server.Addr, sig.String())




	if err := server.Shutdown(ctx); err != nil {
		glog.Errorf("Failed to shutdown %s: %v", server.Addr, err)
	}
	done <- s
}

func runServer(server *http.Server, name string, listener net.Listener) {
	glog.Infof("%s server starting on: %s", name, server.Addr)
	err := server.Serve(listener)
	glog.Errorf("%s server quit with error: %v", name, err)
}