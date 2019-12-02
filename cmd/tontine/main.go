package main

import (
	"github.com/hieuphq/tontine/src/api"
	"github.com/hieuphq/tontine/src/config"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	api.InitRouter(config.DefaultConfig())
}
