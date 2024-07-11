package main

import (
	"log"
	"os"
	"os/signal"
	"server/internal/app"
	"syscall"
)

func main() {

	log.Printf("starting")
	application := app.New(8888)

	//cacheMem := cache.New(10)

	go application.GRPCSrv.MustRun()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	stopping := <-stop

	log.Printf("application stopping, signal: %s", stopping)

	application.GRPCSrv.Stop()

	log.Println("application stopped")
}
