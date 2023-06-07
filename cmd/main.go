package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"test-golang-check-sites/internal/server"
	"test-golang-check-sites/internal/services"
	"test-golang-check-sites/internal/storage"
)

const (
	pathFileSites = "../data/sites.txt"
)

var (
	Port     = "80"
	interval = 1 // interval in minutes
)

func main() {
	if r := recover(); r != nil {
		log.Fatal("RECOVER:", r)
	}

	ctx, cancel := context.WithCancel(context.Background())

	path := pathFileSites

	str := storage.NewStorage()
	err := str.ParseDomainsPromFile(path)
	if err != nil {
		log.Fatal(err)
	}

	services := services.NewServices(services.Deps{
		Ctx:      ctx,
		Storage:  str,
		Interval: interval,
	})

	// start server
	go func() {
		server := server.NewServer(services)
		log.Fatal(server.Run(Port))
	}()

	// graceful shutdown
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt, syscall.SIGTERM)
	<-exit
	cancel()
	log.Fatal("EXIT APP")
}
