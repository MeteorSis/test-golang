package main

import (
	"log"
	"os"
	"strconv"

	"github.com/MeteorSis/test-golang/grpc-interceptor/cmd"
)

func main() {
	log.Default().SetFlags(log.Lmicroseconds)
	if len(os.Args) < 3 {
		log.Fatal("Please specify 'server' or 'client' as first argument.")
	}

	port, err := strconv.ParseUint(os.Args[2], 10, 16)
	if err != nil {
		log.Fatal("Please specify port number as second argument.")
	}

	switch os.Args[1] {
	case "server":
		cmd.RunServer(uint16(port))
	case "client":
		cmd.RunClient(uint16(port))
	default:
		log.Fatal("Please specify 'server' or 'client' as first argument.")
	}
}
