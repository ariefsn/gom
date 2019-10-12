package demo

import (
	"strings"

	"github.com/ariefsn/gom"
	"github.com/ariefsn/gom/examples/models"
	"github.com/eaciit/toolkit"
	"go.mongodb.org/mongo-driver/bson"
)

// Demo = struct of Demo
type Demo struct {
	useParams bool
}

// NewDemo = init new demo
func NewDemo() *Demo {
	d := new(Demo)
	d.useParams = false

	return d
}

// UseParams = set use params
func (d *Demo) UseParams(useParams bool) {
	d.useParams = useParams
}

// InsertStruct = example insert struct data
func (d *Demo) InsertStruct(g *gom.Gom) {
	toolkit.Println("===== Insert Struct =====")
	hero := models.NewHero("Wolverine", "Hugh Jackman", 40)

	var err error
	if d.useParams {
		_, err = g.Set(&gom.SetParams{
			TableName: "hero",
			Timeout:   10,
		}).Cmd().Insert(hero)
	} else {
		_, err = g.Set(nil).Table("hero").Timeout(10).Cmd().Insert(hero)
	}

	if err != nil {
		toolkit.Println(err.Error())
		return
	}
}

// InsertMap = example insert struct map data
func (d *Demo) InsertMap(g *gom.Gom) {
	toolkit.Println("===== Insert Map =====")
	hero := bson.M{
		"Name":     "Black Widow",
		"RealName": "Scarlett Johansson",
		"Age":      32,
	}

	var err error
	if d.useParams {
		_, err = g.Set(&gom.SetParams{
			TableName: "hero",
			Timeout:   10,
		}).Cmd().Insert(&hero)
	} else {
		_, err = g.Set(nil).Table("hero").Timeout(10).Cmd().Insert(&hero)
	}

	if err != nil {
		toolkit.Println(err.Error())
		return
	}
}

// InsertAll = example insert multiple data
func (d *Demo) InsertAll(g *gom.Gom) {
	toolkit.Println("===== Insert All =====")
	heroes := models.DummyData()

	var err error
	if d.useParams {
		_, err = g.Set(&gom.SetParams{
			TableName: "hero",
			Timeout:   10,
		}).Cmd().InsertAll(&heroes)
	} else {
		_, err = g.Set(nil).Table("hero").Timeout(10).Cmd().InsertAll(&heroes)
	}

	if err != nil {
		toolkit.Println(err.Error())
		return
	}
}

// UpdateStruct = example update with struct
func (d *Demo) UpdateStruct(g *gom.Gom) {
	toolkit.Println("===== Update With Struct =====")
	hero := models.NewHero("Wonderwoman", "Gal Gadot", 34)

	var err error
	if d.useParams {
		err = g.Set(&gom.SetParams{
			TableName: "hero",
			Filter:    gom.Eq("RealName", "Scarlett Johansson"),
			Timeout:   10,
		}).Cmd().Update(hero)
	} else {
		err = g.Set(nil).Table("hero").Timeout(10).Filter(gom.Eq("RealName", "Scarlett Johansson")).Cmd().Update(hero)
	}

	if err != nil {
		toolkit.Println(err.Error())
		return
	}
}

// UpdateMap = example update with map
func (d *Demo) UpdateMap(g *gom.Gom) {
	toolkit.Println("===== Update With Map =====")
	hero := bson.M{
		"Name":     "Hawkeye",
		"RealName": "Jeremy Renner",
		"Age":      33,
	}

	var err error
	if d.useParams {
		err = g.Set(&gom.SetParams{
			TableName: "hero",
			Filter:    gom.Eq("Name", "Wolverine"),
			Timeout:   10,
		}).Cmd().Update(&hero)
	} else {
		err = g.Set(nil).Table("hero").Timeout(10).Filter(gom.Eq("Name", "Wolverine")).Cmd().Update(&hero)
	}

	if err != nil {
		toolkit.Println(err.Error())
		return
	}
}

// DeleteOne = example delete one data
func (d *Demo) DeleteOne(g *gom.Gom) {
	toolkit.Println("===== Delete One =====")

	var err error
	if d.useParams {
		err = g.Set(&gom.SetParams{
			TableName: "hero",
			Filter:    gom.Eq("Name", "Batman"),
			Timeout:   10,
		}).Cmd().DeleteOne()
	} else {
		err = g.Set(nil).Table("hero").Timeout(10).Filter(gom.Eq("Name", "Batman")).Cmd().DeleteOne()
	}

	if err != nil {
		toolkit.Println(err.Error())
		return
	}
}

// DeleteAll = example delete all
func (d *Demo) DeleteAll(g *gom.Gom) {
	toolkit.Println("===== Delete All =====")

	var err error
	if d.useParams {
		_, err = g.Set(&gom.SetParams{
			TableName: "hero",
			Filter:    gom.EndWith("Name", "man"),
			Timeout:   10,
		}).Cmd().DeleteAll()
	} else {
		_, err = g.Set(nil).Table("hero").Timeout(10).Filter(gom.EndWith("Name", "man")).Cmd().DeleteAll()
	}

	if err != nil {
		toolkit.Println(err.Error())
		return
	}
}

// GetAll = example get all data without filter or pipe
func (d *Demo) GetAll(g *gom.Gom) int64 {
	toolkit.Println("===== Get All =====")
	res := []models.Hero{}

	var cFilter, cTotal int64
	var err error
	if d.useParams {
		cFilter, cTotal, err = g.Set(&gom.SetParams{
			TableName: "hero",
			Result:    &res,
			Timeout:   10,
		}).Cmd().Get()
	} else {
		cFilter, cTotal, err = g.Set(nil).Timeout(10).Table("hero").Result(&res).Cmd().Get()
	}

	if err != nil {
		toolkit.Println(err.Error())
		return 0
	}

	toolkit.Println(cFilter, "of", cTotal)

	for _, h := range res {
		toolkit.Println(h)
	}

	return cFilter
}

// GetOne = example get one data without filter or pipe
func (d *Demo) GetOne(g *gom.Gom) {
	toolkit.Println("===== Get One =====")
	res := models.Hero{}

	var err error
	if d.useParams {
		err = g.Set(&gom.SetParams{
			TableName: "hero",
			Result:    &res,
			Timeout:   10,
		}).Cmd().GetOne()
	} else {
		err = g.Set(nil).Timeout(10).Table("hero").Result(&res).Cmd().GetOne()
	}

	if err != nil {
		toolkit.Println(err.Error())
		return
	}

	toolkit.Println(res)
}

// Sort = example get all data with sort command
func (d *Demo) Sort(g *gom.Gom, sortBy string) {
	toolkit.Println(toolkit.Sprintf("===== Sort %s =====", strings.ToUpper(sortBy)))
	res := []models.Hero{}

	var err error
	if d.useParams {
		_, _, err = g.Set(&gom.SetParams{
			TableName: "hero",
			Result:    &res,
			SortField: "RealName",
			SortBy:    sortBy,
			Timeout:   10,
		}).Cmd().Get()
	} else {
		_, _, err = g.Set(nil).Table("hero").Timeout(10).Result(&res).Sort("RealName", sortBy).Cmd().Get()
	}

	if err != nil {
		toolkit.Println(err.Error())
	}

	for _, h := range res {
		toolkit.Println(h.RealName, "=>", h.Name, "=>", h.Age)
	}
}

// FilterEq = example get data with filter equal
func (d *Demo) FilterEq(g *gom.Gom) {
	toolkit.Println("===== Equal =====")
	res := models.Hero{}

	var err error
	if d.useParams {
		err = g.Set(&gom.SetParams{
			TableName: "hero",
			Result:    &res,
			Filter:    gom.Eq("Name", "Green Arrow"),
			Timeout:   10,
		}).Cmd().GetOne()
	} else {
		err = g.Set(nil).Table("hero").Timeout(10).Result(&res).Filter(gom.Eq("Name", "Green Arrow")).Cmd().GetOne()
	}

	if err != nil {
		toolkit.Println(err.Error())
		return
	}

	toolkit.Println(res)
}

// FilterNe = example get data with filter not equal
func (d *Demo) FilterNe(g *gom.Gom) {
	toolkit.Println("===== Not Equal =====")
	res := []models.Hero{}

	var err error
	if d.useParams {
		_, _, err = g.Set(&gom.SetParams{
			TableName: "hero",
			Result:    &res,
			Filter:    gom.Ne("Name", "Ironman"),
			Timeout:   10,
		}).Cmd().Get()
	} else {
		_, _, err = g.Set(nil).Table("hero").Timeout(10).Result(&res).Filter(gom.Ne("Name", "Ironman")).Cmd().Get()
	}

	if err != nil {
		toolkit.Println(err.Error())
		return
	}

	for _, h := range res {
		toolkit.Println(h)
	}
}

// FilterGt = example get data with filter greater than
func (d *Demo) FilterGt(g *gom.Gom) {
	toolkit.Println("===== Greater Than =====")
	res := []models.Hero{}

	var cFilter, cTotal int64
	var err error
	if d.useParams {
		cFilter, cTotal, err = g.Set(&gom.SetParams{
			TableName: "hero",
			Result:    &res,
			Filter:    gom.Gt("Age", 35),
			Timeout:   10,
		}).Cmd().Get()
	} else {
		cFilter, cTotal, err = g.Set(nil).Table("hero").Timeout(10).Result(&res).Filter(gom.Gt("Age", 35)).Cmd().Get()
	}

	if err != nil {
		toolkit.Println(err.Error())
		return
	}

	toolkit.Println(cFilter, "of", cTotal)

	for _, h := range res {
		toolkit.Println(h)
	}
}

// FilterGte = example get data with filter greater than equal
func (d *Demo) FilterGte(g *gom.Gom) {
	toolkit.Println("===== Greater Than Equal =====")
	res := []models.Hero{}

	var err error
	if d.useParams {
		_, _, err = g.Set(&gom.SetParams{
			TableName: "hero",
			Result:    &res,
			Filter:    gom.Gte("Age", 43),
			Timeout:   10,
		}).Cmd().Get()
	} else {
		_, _, err = g.Set(nil).Table("hero").Timeout(10).Result(&res).Filter(gom.Gte("Age", 43)).Cmd().Get()
	}

	if err != nil {
		toolkit.Println(err.Error())
		return
	}

	for _, h := range res {
		toolkit.Println(h)
	}
}

// FilterLt = example get data with filter less than
func (d *Demo) FilterLt(g *gom.Gom) {
	toolkit.Println("===== Less Than =====")
	res := []models.Hero{}

	var err error
	if d.useParams {
		_, _, err = g.Set(&gom.SetParams{
			TableName: "hero",
			Result:    &res,
			Filter:    gom.Lt("Age", 35),
			Timeout:   10,
		}).Cmd().Get()
	} else {
		_, _, err = g.Set(nil).Table("hero").Timeout(10).Result(&res).Filter(gom.Lt("Age", 35)).Cmd().Get()
	}

	if err != nil {
		toolkit.Println(err.Error())
		return
	}

	for _, h := range res {
		toolkit.Println(h)
	}
}

// FilterLte = example get data with filter less than equal
func (d *Demo) FilterLte(g *gom.Gom) {
	toolkit.Println("===== Less Than Equal =====")
	res := []models.Hero{}

	var err error
	if d.useParams {
		_, _, err = g.Set(&gom.SetParams{
			TableName: "hero",
			Result:    &res,
			Filter:    gom.Lte("Age", 27),
			Timeout:   10,
		}).Cmd().Get()
	} else {
		_, _, err = g.Set(nil).Table("hero").Timeout(10).Result(&res).Filter(gom.Lte("Age", 27)).Cmd().Get()
	}

	if err != nil {
		toolkit.Println(err.Error())
		return
	}

	for _, h := range res {
		toolkit.Println(h)
	}
}

// FilterBetweenOrRange = example get data with filter between or range as alternative
func (d *Demo) FilterBetweenOrRange(g *gom.Gom) {
	toolkit.Println("===== Between =====")
	res := []models.Hero{}

	var err error
	if d.useParams {
		_, _, err = g.Set(&gom.SetParams{
			TableName: "hero",
			Result:    &res,
			Filter:    gom.Between("Age", 27, 38),
			Timeout:   10,
		}).Cmd().Get()
	} else {
		_, _, err = g.Set(nil).Table("hero").Timeout(10).Result(&res).Filter(gom.Between("Age", 27, 38)).Cmd().Get()
	}

	if err != nil {
		toolkit.Println(err.Error())
		return
	}

	for _, h := range res {
		toolkit.Println(h)
	}
}

// FilterContains = example get data with filter contains
func (d *Demo) FilterContains(g *gom.Gom) {
	toolkit.Println("===== Contains =====")
	res := []models.Hero{}

	var err error
	if d.useParams {
		_, _, err = g.Set(&gom.SetParams{
			TableName: "hero",
			Result:    &res,
			Filter:    gom.Contains("Name", "der"),
			Timeout:   10,
		}).Cmd().Get()
	} else {
		_, _, err = g.Set(nil).Table("hero").Timeout(10).Result(&res).Filter(gom.Contains("Name", "der")).Cmd().Get()
	}

	if err != nil {
		toolkit.Println(err.Error())
		return
	}

	for _, h := range res {
		toolkit.Println(h)
	}
}

// FilterStartWith = example get data with filter startWith
func (d *Demo) FilterStartWith(g *gom.Gom) {
	toolkit.Println("===== Start With =====")
	res := []models.Hero{}

	var err error
	if d.useParams {
		_, _, err = g.Set(&gom.SetParams{
			TableName: "hero",
			Result:    &res,
			Filter:    gom.StartWith("Name", "S"),
			Timeout:   10,
		}).Cmd().Get()
	} else {
		_, _, err = g.Set(nil).Table("hero").Timeout(10).Result(&res).Filter(gom.StartWith("Name", "S")).Cmd().Get()
	}

	if err != nil {
		toolkit.Println(err.Error())
		return
	}

	for _, h := range res {
		toolkit.Println(h)
	}
}

// FilterEndWith = example get data with filter endWith
func (d *Demo) FilterEndWith(g *gom.Gom) {
	toolkit.Println("===== End With =====")
	res := []models.Hero{}

	var err error
	if d.useParams {
		_, _, err = g.Set(&gom.SetParams{
			TableName: "hero",
			Result:    &res,
			Filter:    gom.EndWith("Name", "man"),
			Timeout:   10,
		}).Cmd().Get()
	} else {
		_, _, err = g.Set(nil).Table("hero").Timeout(10).Result(&res).Filter(gom.EndWith("Name", "man")).Cmd().Get()
	}

	if err != nil {
		toolkit.Println(err.Error())
		return
	}

	for _, h := range res {
		toolkit.Println(h)
	}
}

// FilterIn = example get data with filter in
func (d *Demo) FilterIn(g *gom.Gom) {
	toolkit.Println("===== In =====")
	res := []models.Hero{}

	var err error
	if d.useParams {
		_, _, err = g.Set(&gom.SetParams{
			TableName: "hero",
			Result:    &res,
			Filter:    gom.In("Name", "Green Arrow", "Red Arrow"),
			Timeout:   10,
		}).Cmd().Get()
	} else {
		_, _, err = g.Set(nil).Table("hero").Timeout(10).Result(&res).Filter(gom.In("Name", "Green Arrow", "Red Arrow")).Cmd().Get()
	}

	if err != nil {
		toolkit.Println(err.Error())
		return
	}

	for _, h := range res {
		toolkit.Println(h)
	}
}

// FilterNin = example get data with filter not in
func (d *Demo) FilterNin(g *gom.Gom) {
	toolkit.Println("===== Not In =====")
	res := []models.Hero{}
	names := []interface{}{"Green Arrow", "Red Arrow"}

	var err error
	if d.useParams {
		_, _, err = g.Set(&gom.SetParams{
			TableName: "hero",
			Result:    &res,
			Filter:    gom.Nin("Name", names...),
			Timeout:   10,
		}).Cmd().Get()
	} else {
		_, _, err = g.Set(nil).Table("hero").Timeout(10).Result(&res).Filter(gom.Nin("Name", names...)).Cmd().Get()
	}

	if err != nil {
		toolkit.Println(err.Error())
		return
	}

	for _, h := range res {
		toolkit.Println(h)
	}
}

// GetByPipe = example get all data pipe
func (d *Demo) GetByPipe(g *gom.Gom) {
	toolkit.Println("===== Get By Pipe =====")
	res := []models.Hero{}
	pipe := []bson.M{
		bson.M{
			"$match": bson.M{
				"Name": bson.M{
					"$in": []string{"Superman", "Batman", "Flash"},
				},
			},
		},
		bson.M{
			"$sort": bson.M{
				"RealName": -1,
			},
		},
	}

	var cFilter, cTotal int64
	var err error
	if d.useParams {
		cFilter, cTotal, err = g.Set(&gom.SetParams{
			TableName: "hero",
			Result:    &res,
			Pipe:      pipe,
			Timeout:   10,
		}).Cmd().Get()
	} else {
		cFilter, cTotal, err = g.Set(nil).Table("hero").Result(&res).Timeout(10).Pipe(pipe).Cmd().Get()
	}

	if err != nil {
		toolkit.Println(err.Error())
		return
	}

	toolkit.Println(cFilter, "of", cTotal)

	for _, h := range res {
		toolkit.Println(h)
	}
}

// FilterAnd = example get data with filter and
func (d *Demo) FilterAnd(g *gom.Gom) {
	toolkit.Println("===== And =====")
	res := []models.Hero{}
	filter := gom.And(gom.Eq("Age", 45), gom.StartWith("Name", "A"))

	var err error
	if d.useParams {
		_, _, err = g.Set(&gom.SetParams{
			TableName: "hero",
			Result:    &res,
			Filter:    filter,
			Timeout:   10,
		}).Cmd().Get()
	} else {
		_, _, err = g.Set(nil).Table("hero").Timeout(10).Result(&res).Filter(filter).Cmd().Get()
	}

	if err != nil {
		toolkit.Println(err.Error())
		return
	}

	for _, h := range res {
		toolkit.Println(h)
	}
}

// FilterOr = example get data with filter or
func (d *Demo) FilterOr(g *gom.Gom) {
	toolkit.Println("===== Or =====")
	res := []models.Hero{}
	filter := gom.Or(gom.Eq("Age", 45), gom.StartWith("Name", "A"))

	var err error
	if d.useParams {
		_, _, err = g.Set(&gom.SetParams{
			TableName: "hero",
			Result:    &res,
			Filter:    filter,
			Timeout:   10,
		}).Cmd().Get()
	} else {
		_, _, err = g.Set(nil).Table("hero").Timeout(10).Result(&res).Filter(filter).Cmd().Get()
	}

	if err != nil {
		toolkit.Println(err.Error())
		return
	}

	for _, h := range res {
		toolkit.Println(h)
	}
}
