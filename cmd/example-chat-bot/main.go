package main

import (
	"fmt"

	wapi "github.com/sarthakjdev/wapi.go/pkg/client"
)

func main() {
	fmt.Println("This is the main package entry point of my golang file")
	whatsappClient := wapi.NewWapiClient(wapi.ClientConfig{})
	fmt.Print(whatsappClient)
}
