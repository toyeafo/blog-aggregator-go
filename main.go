package main

import (
	"fmt"

	"github.com/toyeafo/blog-aggregator-go/internal/config"
)

func main() {
	cfg := config.Read()
	cfg.SetUser("toye")
	cfg = config.Read()
	fmt.Println(cfg)
}
