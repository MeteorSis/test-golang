package main

import (
	"context"
	"fmt"

	"github.com/MeteorSis/test-golang/interface-embedding/internal/usecase"
	"github.com/MeteorSis/test-golang/interface-embedding/internal/usecase/message"
)

func main() {
	var msgUsecase usecase.MessageUsecase
	msgUsecase = message.New("msg content")

	msg, err := msgUsecase.GetMessage(context.Background(), "cID", 1) // panic!!
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("get message: %#v\n", msg)
}
