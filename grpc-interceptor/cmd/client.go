package cmd

import (
	"log"

	"github.com/MeteorSis/test-golang/grpc-interceptor/internal/client"
)

func RunClient(port uint16) {
	if err := client.DoRequest(port); err != nil {
		log.Println(err)
	}
}
