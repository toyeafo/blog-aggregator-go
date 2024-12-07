package main

import (
	"log"

	"github.com/toyeafo/blog-aggregator-go/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config file %v", err)
	}
	stg := state{cfg: &cfg}
}
