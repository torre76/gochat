package main

import (
	"github.com/torre76/gochat/chat"
)

func main() {
	chat.StartServer("0.0.0.0", 10000)
}
