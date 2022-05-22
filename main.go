package main

import (
	"ascii-art/run"
	"log"
)

func main() {
	err := run.Run()
	if err != nil {
		log.Fatal(err)
	}
}
