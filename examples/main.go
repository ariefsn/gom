package main

import (
	"github.com/ariefsn/gom/examples/demo"
	"github.com/eaciit/toolkit"

	"github.com/ariefsn/gom"
)

func main() {
	g := gom.NewGom()

	cfg := gom.MongoConfig{
		Host:     "localhost",
		Port:     "27017",
		Username: "",
		Password: "",
		Database: "test",
	}

	g.Init(cfg)

	err := g.CheckClient()

	if err != nil {
		toolkit.Println(toolkit.Sprintf("Connection Error: %s", err.Error()))
	}

	if demo.GetAll(g) == 0 {
		demo.InsertStruct(g)
		demo.InsertMap(g)
		demo.InsertAll(g)
		demo.GetAll(g)
		demo.UpdateStruct(g)
		demo.GetAll(g)
		demo.UpdateMap(g)
		demo.GetAll(g)
		demo.DeleteOne(g)
		demo.GetAll(g)
	}
	demo.GetOne(g)
	demo.FilterEq(g)
	demo.FilterNe(g)
	demo.FilterGt(g)
	demo.FilterGte(g)
	demo.FilterLt(g)
	demo.FilterLte(g)
	demo.FilterBetweenOrRange(g)
	demo.FilterContains(g)
	demo.FilterStartWith(g)
	demo.FilterEndWith(g)
	demo.FilterIn(g)
	demo.FilterNin(g)
	demo.GetByPipe(g)
	demo.FilterAnd(g)
	demo.FilterOr(g)
	demo.Sort(g, "asc")
	demo.Sort(g, "desc")
	demo.DeleteAll(g)
	demo.GetAll(g)
}
