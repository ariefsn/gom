package demo

import (
	"strings"

	"github.com/ariefsn/gom"
	"github.com/ariefsn/gom/examples/models"
	"github.com/eaciit/toolkit"
	"go.mongodb.org/mongo-driver/bson"
)

// InsertStruct = example insert struct data
func InsertStruct(g *gom.Gom) {
	toolkit.Println("===== Insert Struct =====")
	hero := models.NewHero("Wolverine", "Hugh Jackman", 40)
	err := g.Set().Table("hero").Cmd().Insert(hero)
	if err != nil {
		toolkit.Println(err.Error())
		return
	}
}

// InsertMap = example insert struct map data
func InsertMap(g *gom.Gom) {
	toolkit.Println("===== Insert Map =====")
	hero := bson.M{
		"Name":     "Black Widow",
		"RealName": "Scarlett Johansson",
		"Age":      32,
	}
	err := g.Set().Table("hero").Cmd().Insert(&hero)
	if err != nil {
		toolkit.Println(err.Error())
		return
	}
}

// InsertAll = example insert multiple data
func InsertAll(g *gom.Gom) {
	toolkit.Println("===== Insert All =====")
	heroes := models.DummyData()
	err := g.Set().Table("hero").Cmd().InsertAll(&heroes)
	if err != nil {
		toolkit.Println(err.Error())
		return
	}
}

// UpdateStruct = example update with struct
func UpdateStruct(g *gom.Gom) {
	toolkit.Println("===== Update With Struct =====")
	hero := models.NewHero("Wonderwoman", "Gal Gadot", 34)
	err := g.Set().Table("hero").Filter(gom.Eq("RealName", "Scarlett Johansson")).Cmd().Update(hero)
	if err != nil {
		toolkit.Println(err.Error())
		return
	}
}

// UpdateMap = example update with map
func UpdateMap(g *gom.Gom) {
	toolkit.Println("===== Update With Map =====")
	hero := bson.M{
		"Name":     "Hawkeye",
		"RealName": "Jeremy Renner",
		"Age":      33,
	}
	err := g.Set().Table("hero").Filter(gom.Eq("Name", "Wolverine")).Cmd().Update(&hero)
	if err != nil {
		toolkit.Println(err.Error())
		return
	}
}

// DeleteOne = example delete one data
func DeleteOne(g *gom.Gom) {
	toolkit.Println("===== Delete One =====")
	err := g.Set().Table("hero").Filter(gom.Eq("Name", "Batman")).Cmd().DeleteOne()
	if err != nil {
		toolkit.Println(err.Error())
		return
	}
}

// DeleteAll = example delete all
func DeleteAll(g *gom.Gom) {
	toolkit.Println("===== Delete All =====")
	err := g.Set().Table("hero").Filter(gom.EndWith("Name", "man")).Cmd().DeleteAll()
	if err != nil {
		toolkit.Println(err.Error())
		return
	}
}

// GetAll = example get all data without filter or pipe
func GetAll(g *gom.Gom) int {
	toolkit.Println("===== Get All =====")
	res := []models.Hero{}
	err := g.Set().Table("hero").Result(&res).Cmd().Get()
	if err != nil {
		toolkit.Println(err.Error())
		return 0
	}
	for _, h := range res {
		toolkit.Println(h)
	}
	return len(res)
}

// GetOne = example get one data without filter or pipe
func GetOne(g *gom.Gom) {
	toolkit.Println("===== Get One =====")
	res := models.Hero{}
	err := g.Set().Table("hero").Result(&res).Cmd().GetOne()
	if err != nil {
		toolkit.Println(err.Error())
		return
	}
	toolkit.Println(res)
}

// Sort = example get all data with sort command
func Sort(g *gom.Gom, sortBy string) {
	toolkit.Println(toolkit.Sprintf("===== Sort %s =====", strings.ToUpper(sortBy)))
	res := []models.Hero{}
	err := g.Set().Table("hero").Result(&res).Sort("RealName", sortBy).Cmd().Get()
	if err != nil {
		toolkit.Println(err.Error())
	}
	for _, h := range res {
		toolkit.Println(h.RealName, "=>", h.Name, "=>", h.Age)
	}
}

// FilterEq = example get data with filter equal
func FilterEq(g *gom.Gom) {
	toolkit.Println("===== Equal =====")
	res := models.Hero{}
	err := g.Set().Table("hero").Result(&res).Filter(gom.Eq("Name", "Green Arrow")).Cmd().GetOne()
	if err != nil {
		toolkit.Println(err.Error())
		return
	}
	toolkit.Println(res)
}

// FilterNe = example get data with filter not equal
func FilterNe(g *gom.Gom) {
	toolkit.Println("===== Not Equal =====")
	res := []models.Hero{}
	err := g.Set().Table("hero").Result(&res).Filter(gom.Ne("Name", "Ironman")).Cmd().Get()
	if err != nil {
		toolkit.Println(err.Error())
		return
	}
	for _, h := range res {
		toolkit.Println(h)
	}
}

// FilterGt = example get data with filter greater than
func FilterGt(g *gom.Gom) {
	toolkit.Println("===== Greater Than =====")
	res := []models.Hero{}
	err := g.Set().Table("hero").Result(&res).Filter(gom.Gt("Age", 35)).Cmd().Get()
	if err != nil {
		toolkit.Println(err.Error())
		return
	}
	for _, h := range res {
		toolkit.Println(h)
	}
}

// FilterGte = example get data with filter greater than equal
func FilterGte(g *gom.Gom) {
	toolkit.Println("===== Greater Than Equal =====")
	res := []models.Hero{}
	err := g.Set().Table("hero").Result(&res).Filter(gom.Gte("Age", 43)).Cmd().Get()
	if err != nil {
		toolkit.Println(err.Error())
		return
	}
	for _, h := range res {
		toolkit.Println(h)
	}
}

// FilterLt = example get data with filter less than
func FilterLt(g *gom.Gom) {
	toolkit.Println("===== Less Than =====")
	res := []models.Hero{}
	err := g.Set().Table("hero").Result(&res).Filter(gom.Lt("Age", 35)).Cmd().Get()
	if err != nil {
		toolkit.Println(err.Error())
		return
	}
	for _, h := range res {
		toolkit.Println(h)
	}
}

// FilterLte = example get data with filter less than equal
func FilterLte(g *gom.Gom) {
	toolkit.Println("===== Less Than Equal =====")
	res := []models.Hero{}
	err := g.Set().Table("hero").Result(&res).Filter(gom.Lte("Age", 27)).Cmd().Get()
	if err != nil {
		toolkit.Println(err.Error())
		return
	}
	for _, h := range res {
		toolkit.Println(h)
	}
}

// FilterBetweenOrRange = example get data with filter between or range as alternative
func FilterBetweenOrRange(g *gom.Gom) {
	toolkit.Println("===== Between =====")
	res := []models.Hero{}
	err := g.Set().Table("hero").Result(&res).Filter(gom.Between("Age", 27, 38)).Cmd().Get()
	if err != nil {
		toolkit.Println(err.Error())
		return
	}
	for _, h := range res {
		toolkit.Println(h)
	}
}

// FilterContains = example get data with filter contains
func FilterContains(g *gom.Gom) {
	toolkit.Println("===== Contains =====")
	res := []models.Hero{}
	err := g.Set().Table("hero").Result(&res).Filter(gom.Contains("Name", "der")).Cmd().Get()
	if err != nil {
		toolkit.Println(err.Error())
		return
	}
	for _, h := range res {
		toolkit.Println(h)
	}
}

// FilterStartWith = example get data with filter startWith
func FilterStartWith(g *gom.Gom) {
	toolkit.Println("===== Start With =====")
	res := []models.Hero{}
	err := g.Set().Table("hero").Result(&res).Filter(gom.StartWith("Name", "S")).Cmd().Get()
	if err != nil {
		toolkit.Println(err.Error())
		return
	}
	for _, h := range res {
		toolkit.Println(h)
	}
}

// FilterEndWith = example get data with filter endWith
func FilterEndWith(g *gom.Gom) {
	toolkit.Println("===== End With =====")
	res := []models.Hero{}
	err := g.Set().Table("hero").Result(&res).Filter(gom.EndWith("Name", "man")).Cmd().Get()
	if err != nil {
		toolkit.Println(err.Error())
		return
	}
	for _, h := range res {
		toolkit.Println(h)
	}
}

// FilterIn = example get data with filter in
func FilterIn(g *gom.Gom) {
	toolkit.Println("===== In =====")
	res := []models.Hero{}
	err := g.Set().Table("hero").Result(&res).Filter(gom.In("Name", "Green Arrow", "Red Arrow")).Cmd().Get()
	if err != nil {
		toolkit.Println(err.Error())
		return
	}
	for _, h := range res {
		toolkit.Println(h)
	}
}

// FilterNin = example get data with filter not in
func FilterNin(g *gom.Gom) {
	toolkit.Println("===== Not In =====")
	res := []models.Hero{}
	names := []interface{}{"Green Arrow", "Red Arrow"}
	err := g.Set().Table("hero").Result(&res).Filter(gom.Nin("Name", names...)).Cmd().Get()
	if err != nil {
		toolkit.Println(err.Error())
		return
	}
	for _, h := range res {
		toolkit.Println(h)
	}
}

// GetByPipe = example get all data pipe
func GetByPipe(g *gom.Gom) {
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
	err := g.Set().Table("hero").Result(&res).Pipe(pipe).Cmd().Get()
	if err != nil {
		toolkit.Println(err.Error())
		return
	}
	for _, h := range res {
		toolkit.Println(h)
	}
}

// FilterAnd = example get data with filter and
func FilterAnd(g *gom.Gom) {
	toolkit.Println("===== And =====")
	res := []models.Hero{}
	filter := gom.And(gom.Eq("Age", 45), gom.StartWith("Name", "A"))
	err := g.Set().Table("hero").Result(&res).Filter(filter).Cmd().Get()
	if err != nil {
		toolkit.Println(err.Error())
		return
	}
	for _, h := range res {
		toolkit.Println(h)
	}
}

// FilterOr = example get data with filter or
func FilterOr(g *gom.Gom) {
	toolkit.Println("===== Or =====")
	res := []models.Hero{}
	filter := gom.Or(gom.Eq("Age", 45), gom.StartWith("Name", "A"))
	err := g.Set().Table("hero").Result(&res).Filter(filter).Cmd().Get()
	if err != nil {
		toolkit.Println(err.Error())
		return
	}
	for _, h := range res {
		toolkit.Println(h)
	}
}
