package main

import (
	"github.com/ariefsn/gom/examples/demo"
	"github.com/eaciit/toolkit"

	"github.com/ariefsn/gom"
)

func main() {
	g := gom.NewGom()

	cfg := gom.Config{
		Host:     "localhost",
		Port:     27017,
		Username: "",
		Password: "",
		Database: "test",
	}

	g.Init(cfg)

	err := g.CheckClient()

	if err != nil {
		toolkit.Println(toolkit.Sprintf("Connection Error: %s", err.Error()))
		return
	}

	d := demo.NewDemo()
	// false => chaining set
	// true => use SetParams
	d.UseParams(true)

	if d.GetAll(g) == 0 {
		d.InsertStruct(g)
		d.InsertMap(g)
		d.InsertAll(g)
		d.GetAll(g)
		d.UpdateStruct(g)
		d.GetAll(g)
		d.UpdateMap(g)
		d.GetAll(g)
		d.DeleteOne(g)
		d.GetAll(g)
	}
	d.GetOne(g)
	d.FilterEq(g)
	d.FilterNe(g)
	d.FilterGt(g)
	d.FilterGte(g)
	d.FilterLt(g)
	d.FilterLte(g)
	d.FilterBetweenOrRange(g)
	d.FilterContains(g)
	d.FilterStartWith(g)
	d.FilterEndWith(g)
	d.FilterIn(g)
	d.FilterNin(g)
	d.FilterExists(g)
	d.GetByPipe(g)
	d.FilterAnd(g)
	d.FilterOr(g)
	d.Sort(g, "asc")
	d.Sort(g, "desc")
	d.DeleteAll(g)
	d.GetAll(g)
}
