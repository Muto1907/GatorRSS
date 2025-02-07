package config

import "github.com/Muto1907/GatorRSS/internal/database"

type State struct {
	Db     *database.Queries
	Config *Config
}
