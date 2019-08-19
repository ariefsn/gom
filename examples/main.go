package main

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/eaciit/toolkit"

	"github.com/ariefsn/gom"
	"github.com/ariefsn/gom/examples/models"
)

func main() {
	jos := gom.NewGom()

	jos.Init(gom.MongoConfig{
		Host:     "localhost",
		Port:     "27017",
		Username: "root",
		Password: "toor",
		Database: "test",
	})

	var res []models.Hero

	toolkit.Println("===== 01. And =====")
	jos.Data().Table("hero").Result(&res).Filter(gom.And(gom.Eq("Name", "Iron Man"), gom.Eq("RealName", "Tony Stark"))).Get()
	for _, h := range res {
		fmt.Println(h.RealName, "===>", h.Age)
	}
	toolkit.Println("===== 02. Or =====")
	jos.Data().Table("hero").Result(&res).Filter(gom.Or(gom.Eq("Name", "Superman"), gom.Eq("RealName", "Tony Stark"))).Get()
	for _, h := range res {
		fmt.Println(h.RealName)
	}
	// toolkit.Println("===== 03. Not =====")
	// jos.Data().Table("hero").Result(&res).Filter(gom.Not(gom.And(gom.Eq("Name", "Spiderman"), gom.Eq("Name", "Iron Man")))).Get()
	// for _, h := range res {
	// 	fmt.Println(h.RealName)
	// }
	toolkit.Println("===== 04. Equal =====")
	jos.Data().Table("hero").Result(&res).Filter(gom.Eq("Name", "Spiderman")).Get()
	for _, h := range res {
		fmt.Println(h.RealName)
	}
	toolkit.Println("===== 05. Not Equal =====")
	jos.Data().Table("hero").Result(&res).Filter(gom.Ne("Name", "Spiderman")).Get()
	for _, h := range res {
		fmt.Println(h.RealName)
	}
	toolkit.Println("===== 06. Greater Than or Equal =====")
	jos.Data().Table("hero").Result(&res).Filter(gom.Gte("Age", 43)).Get()
	for _, h := range res {
		fmt.Println(h.RealName, "===>", h.Age)
	}
	toolkit.Println("===== 07. Greater Than =====")
	jos.Data().Table("hero").Result(&res).Filter(gom.Gt("Age", 43)).Get()
	for _, h := range res {
		fmt.Println(h.RealName, "===>", h.Age)
	}
	toolkit.Println("===== 08. Less Than or Equal =====")
	jos.Data().Table("hero").Result(&res).Filter(gom.Lte("Age", 43)).Get()
	for _, h := range res {
		fmt.Println(h.RealName, "===>", h.Age)
	}
	toolkit.Println("===== 09. Less Than =====")
	jos.Data().Table("hero").Result(&res).Filter(gom.Lt("Age", 43)).Get()
	for _, h := range res {
		fmt.Println(h.RealName, "===>", h.Age)
	}
	toolkit.Println("===== 10. Between =====")
	jos.Data().Table("hero").Result(&res).Filter(gom.Between("Age", 25, 44)).Sort("Age", "desc").Get()
	for _, h := range res {
		fmt.Println(h.RealName, "===>", h.Age)
	}
	toolkit.Println("===== 11. Range =====")
	jos.Data().Table("hero").Result(&res).Filter(gom.Range("Age", 25, 44)).Get()
	for _, h := range res {
		fmt.Println(h.RealName, "===>", h.Age)
	}
	toolkit.Println("===== 12. Contains =====")
	jos.Data().Table("hero").Result(&res).Filter(gom.Contains("Name", "der")).Get()
	for _, h := range res {
		fmt.Println(h.RealName)
	}
	toolkit.Println("===== 13. StartWith =====")
	jos.Data().Table("hero").Result(&res).Filter(gom.StartWith("RealName", "T")).Get()
	for _, h := range res {
		fmt.Println(h.RealName)
	}
	toolkit.Println("===== 14. EndWith =====")
	jos.Data().Table("hero").Result(&res).Filter(gom.EndWith("Name", "Man")).Get()
	for _, h := range res {
		fmt.Println(h.RealName)
	}
	toolkit.Println("===== 15. In =====")
	jos.Data().Table("hero").Result(&res).Filter(gom.And(gom.In("Name", "Spiderman", "Iron Man", "Green Arrow", "Batman"))).Sort("RealName", "Desc").Get()
	for _, h := range res {
		fmt.Println(h.RealName)
	}
	toolkit.Println("===== 16. Nin =====")
	jos.Data().Table("hero").Result(&res).Filter(gom.And(gom.Nin("Name", "Spiderman", "Iron Man", "Green Arrow", "Batman"))).Sort("RealName", "Desc").Get()
	for _, h := range res {
		fmt.Println(h.RealName)
	}
	toolkit.Println("===== 17. Use Pipe =====")
	jos.Data().Table("hero").Result(&res).Pipe([]bson.M{
		bson.M{
			"$match": bson.M{
				"Name": bson.M{
					"$in": []string{"Green Arrow", "Red Arrow", "The Flash"},
				},
			},
		},
		bson.M{
			"$sort": bson.M{
				"RealName": -1,
			},
		},
	}).Limit(1).Get()
	for _, h := range res {
		fmt.Println(h.RealName)
	}
}
