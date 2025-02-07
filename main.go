package main

import (
	"database/sql"
	"fmt"
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
	fmt.Printf("%s\n", cfg.DbUrl)
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
	cmd := command.Command{Name: args[1], Args: args[2:]}
	err = commands.Run(state, cmd)

	if err != nil {
		log.Fatal(err)
	}

}
