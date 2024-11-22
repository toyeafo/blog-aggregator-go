package main

import (
	"fmt"

	"github.com/toyeafo/blog-aggregator-go/internal/config"
)

type state struct {
	cfg *config.Config
}

type command struct {
	name string
	args []string
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("please provide a username")
	}
	s.cfg.User_name = cmd.args[1]
	fmt.Println("User has been set")
	return nil
}
