package main

import (
	"log"

	"github.com/hyperversal-blocks/bmlloopbe/cmd/server"
)

func main() {
	if err := server.Init(); err != nil {
		log.Fatal(err)
	}
}
