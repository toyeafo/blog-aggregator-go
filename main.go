package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/toyeafo/blog-aggregator-go/internal/config"
	"github.com/toyeafo/blog-aggregator-go/internal/database"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config file %v", err)
	}
	stg := &state{cfg: &cfg}

	cmds := commands{commandMap: make(map[string]func(*state, command) error)}
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)

	if len(os.Args) < 2 {
		log.Fatal("not enough arguments")
		// return
	}

	argCommand := command{Name: os.Args[1], Args: os.Args[2:]}

	err = cmds.run(stg, argCommand)
	if err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open("postgres", stg.cfg.Db_url)
	dbQueries := database.New(db)
}
