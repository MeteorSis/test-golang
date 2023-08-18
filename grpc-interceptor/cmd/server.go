package cmd

import (
	"log"

	"github.com/MeteorSis/test-golang/grpc-interceptor/internal/server"
)

func RunServer(port uint16) {
	log.Println(server.Serve(port))
}
