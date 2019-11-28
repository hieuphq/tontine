package main

import (
	"context"
	"fmt"

	"github.com/hieuphq/tontine/src/interfaces/repo/sqlite"
	"github.com/hieuphq/tontine/src/model"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	g := model.Group{
		Name:            "HP",
		StrategyPercent: 20.0,
	}

	s := sqlite.NewStore("./bin/db.db")

	r := sqlite.NewRepo()

	ctx := context.Background()
	newG, err := r.Group.Create(ctx, s, g)
	if err != nil {
		fmt.Println("unable create group: ", err)
		return
	}
	fmt.Println(newG)

	dt, err := r.Group.GetList(ctx, s)
	if err != nil {
		fmt.Println("unable get group list: ", err)
		return
	}
	fmt.Println(dt)

}
