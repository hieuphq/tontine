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
		Name:            "HTTP",
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

	newG.Name = "HTTP 1"
	updatedG, err := r.Group.Update(ctx, s, *newG)
	if err != nil {
		fmt.Println("unable update group : ", err)
		return
	}
	fmt.Println(updatedG)
	err = r.Group.Delete(ctx, s, *updatedG)
	if err != nil {
		fmt.Println("unable delete group : ", err)
		return
	}

	dts, err := r.Group.GetList(ctx, s)
	if err != nil {
		fmt.Println("unable get group list: ", err)
		return
	}
	fmt.Println(dts)
}
