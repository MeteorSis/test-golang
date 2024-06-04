package cmd

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/MeteorSis/test-golang/grpc-interceptor/internal/server"
)

func RunServer(port uint16) {
	srv, err := server.New(port)
	if err != nil {
		log.Fatal(err)
	}

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-sigCh
		srv.GracefulStop()
		log.Println("server stopped")
	}()

	log.Println("server started")
	if err := srv.Serve(); err != nil {
		log.Fatal(err)
	}
	log.Println("server exited")
}
