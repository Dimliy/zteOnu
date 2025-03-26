package main

import (
	"log"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Panicln(err)
	}
}
