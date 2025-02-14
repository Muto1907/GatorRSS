package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/Muto1907/GatorRSS/command"
	"github.com/Muto1907/GatorRSS/internal/config"
	"github.com/Muto1907/GatorRSS/internal/database"
	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}
	state := &config.State{Config: &cfg}
	db, err := sql.Open("postgres", cfg.DbUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	dbQueries := database.New(db)
	state.Db = dbQueries
	args := os.Args
	if len(args) < 2 {
		log.Fatalf("enter a command name")
	}
	commands := command.Commands{Handlers: make(map[string]func(*config.State, command.Command) error)}
	commands.Register("login", command.HandlerLogin)
	commands.Register("register", command.HandlerRegister)
	commands.Register("reset", command.HandlerReset)
	commands.Register("users", command.HandlerListUsers)
	commands.Register("agg", command.HandleAgg)
	commands.Register("addfeed", command.MiddleWareLoggedIn(command.HandlerAddFeed))
	commands.Register("feeds", command.HandlerListFeeds)
	commands.Register("follow", command.MiddleWareLoggedIn(command.HandlerFollow))
	commands.Register("following", command.MiddleWareLoggedIn(command.HandlerFollowing))
	cmd := command.Command{Name: args[1], Args: args[2:]}
	err = commands.Run(state, cmd)

	if err != nil {
		log.Fatal(err)
	}

}
