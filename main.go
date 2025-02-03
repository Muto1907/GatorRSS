package main

import (
	"fmt"

	"github.com/Muto1907/GatorRSS/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(cfg.DbUrl)
	cfg.SetUser("Mahmut")
	cfg, err = config.Read()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(cfg.DbUrl)
	fmt.Println(cfg.Username)
}
