# [GOM](http://github.com/ariefsn/gom)

 1. [Overview](#overview)
 2. [Dependencies](#dependencies)
 3. [Installation](#installation)
 4. [How to use](#how-to-use)
 5. [Thanks to](#thanks-to)

## Overview

> Gom is a simple mongodb library for golang, inspired by dbflex.

## Dependencies

> 1. [Mongo Go Driver](https://github.com/mongodb/mongo-go-driver)
> 2. [Toolkit](https://github.com/eaciit/toolkit)

## Installation

  ```go
    go get -u github.com/ariefsn/gom
  ```

## How to use

- Create Instance

```go
  g := gom.NewGom()
```

- Set Config

```go
  cfg := gom.MongoConfig{
    Host:     "localhost",
    Port:     "27017",
    Username: "",
    Password: "",
    Database: "test",
  }
```

- Initialize

```go
  g.Init(cfg)
```

> That's it! Gom has ready to use! :)
> Check full [demo](https://github.com/ariefsn/gom/blob/master/examples/demo/demo.go)

- Gom Filter

```go
  // Equal
  // gom.Eq(<Field>, <Value>)
  gom.Eq("Name", "Ironman")

  // Not Equal
  // gom.Ne(<Field>, <Value>)
  gom.Ne("Name", "Batman")

  // Greater Than
  // gom.Gt(<Field>, <Value>)
  gom.Gt("Age", 27)

  // Greater Than Equal
  // gom.Gte(<Field>, <Value>)
  gom.Gte("Age", 28)

  // Less Than
  // gom.Lt(<Field>, <Value>)
  gom.Lt("Age", 32)

  // Less Than Equal
  // gom.Lte(<Field>, <Value>)
  gom.Lte("Age", 33)

  // Range
  // gom.Range(<Field>, <From Value>, <To Value>)
  gom.Range("Age", 20, 28)

  // Between
  // gom.Between(<Field>, <From Value>, <To Value>)
  gom.Between("Age", 20, 28)

  // In
  // gom.In(<Field>, <Values...>)
  gom.In("Name", "Green Arrow", "Red Arrow")

  // Not In
  // gom.Nin(<Field>, <Values...>)
  gom.Nin("Name", "Batman", "Superman")

  // Contains
  // gom.Contains(<Field>, <Value>)
  gom.Contains("Name", "der")

  // Start With
  // gom.StartWith(<Field>, <Value>)
  gom.StartWith("Real Name", "Tony")

  // End With
  // gom.EndWith(<Field>, <Value>)
  gom.EndWith("Name", "man")

  // And
  // gom.And(<Filters...>)
  gom.And(gom.Eq("Age", 45), gom.StartWith("Name", "A"))

  // Or
  // gom.Or(<Filters...>)
  gom.Or(gom.Eq("Age", 45), gom.StartWith("Name", "A"))

```

- Gom Command

  - Get
  > Get all data. It'll use Filter as default. if pipe not null then Filter will be ignored. This command returns countFilterData `int64`, countAllData `int64`, and `error`

  ```go
  
    res := []models.Hero{}
    countFilterData, countAllData, err := g.Set().Table("hero").Result(&res).Cmd().Get()
    if err != nil {
      toolkit.Println(err.Error())
      return 0
    }
    toolkit.Println("Data found", countFilterData, "of", countAllData )
    for _, h := range res {
      toolkit.Println(h)
    }
  ```

  - Get One
  > Get one data. It'll use Filter as default, pipe ignored. This command return `error`.

  ```go
    res := models.Hero{}
    err := g.Set().Table("hero").Result(&res).Cmd().GetOne()
    if err != nil {
      toolkit.Println(err.Error())
      return
    }
    toolkit.Println(res)
  ```

  - Insert
  > Insert one data, for multiple data use InsertAll. This command returns insertedID `interface{}` and `error`.

  ```go
    hero := models.NewHero("Wolverine", "Hugh Jackman", 40)
    _, err := g.Set().Table("hero").Cmd().Insert(hero)
    if err != nil {
      toolkit.Println(err.Error())
      return
    }
  ```
  
  - Insert All
  > Insert multiple data. This command returns insertedIDs `[]interface{}` and `error`

  ```go
    heroes := models.DummyData()
    _, err := g.Set().Table("hero").Cmd().InsertAll(&heroes)
    if err != nil {
      toolkit.Println(err.Error())
      return
    }
  ```

  - Update
  > Update data with filter, pipe will ignored. This command return `error`.

  ```go
    hero := models.NewHero("Wonderwoman", "Gal Gadot", 34)
    err := g.Set().Table("hero").Filter(gom.Eq("RealName", "Scarlett Johansson")).Cmd().Update(hero)
    if err != nil {
      toolkit.Println(err.Error())
      return
    }
  ```

  - Delete One
  > Delete one data with filter, pipe will ignored. This command return `error`.

  ```go
    err := g.Set().Table("hero").Filter(gom.Eq("Name", "Batman")).Cmd().DeleteOne()
    if err != nil {
      toolkit.Println(err.Error())
      return
    }
  ```

  - Delete All
  > Delete all data with filter, pipe will ignored. This command return totalDeletedDocuments `int64` and `error`.

  ```go
    _, err := g.Set().Table("hero").Filter(gom.EndWith("Name", "man")).Cmd().DeleteAll()
    if err != nil {
      toolkit.Println(err.Error())
      return
    }
  ```

  - Sort
  > Sort results ascending or descending

  ```go
    res := []models.Hero{}
    err := g.Set().Table("hero").Result(&res).Sort("RealName", sortBy).Cmd().Get()
    if err != nil {
      toolkit.Println(err.Error())
    }
    for _, h := range res {
      toolkit.Println(h.RealName, "=>", h.Name, "=>", h.Age)
    }
  ```

  - Skip & Limit
  > Set skip & limit for results

  ```go
    res := []models.Hero{}
    err := g.Set().Table("hero").Result(&res).Skip(0).Limit(3).Cmd().Get()
    if err != nil {
      toolkit.Println(err.Error())
      return
    }
    for _, h := range res {
      toolkit.Println(h)
    }
  ```

  - Filter
  > Set filter data

  ```go
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
  ```

  - Pipe
  > Set custom pipe if want to more flexible aggregate

  ```go
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
  ```

## Thanks to

> - Allah :blush:
> - [MongoDB](https://github.com/mongodb/)
> - [Eaciit](https://github.com/eaciit/)
