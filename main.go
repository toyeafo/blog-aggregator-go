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

	cmds := commands{commandMap: make(map[string]func(*state, command) error)}
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", resetHandler)
	cmds.register("users", users)
	cmds.register("agg", handlerAgg)
	cmds.register("addfeed", middlewareLoggedIn(handlerAddFeed))
	cmds.register("feeds", handlerFeeds)
	cmds.register("follow", middlewareLoggedIn(handlerFollow))
	cmds.register("following", middlewareLoggedIn(handlerFollowing))
	cmds.register("unfollow", middlewareLoggedIn(handlerDeleteFollow))

	if len(os.Args) < 2 {
		log.Fatal("not enough arguments")
		// return
	}

	argCommand := command{Name: os.Args[1], Args: os.Args[2:]}

	db, err := sql.Open("postgres", cfg.Db_url)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	dbQueries := database.New(db)
	stg := &state{db: dbQueries, cfg: &cfg}

	err = cmds.run(stg, argCommand)
	if err != nil {
		log.Fatal(err)
	}
}
