# [GOM](http://github.com/ariefsn/gom)

 1. [Overview](#overview)
 2. [Dependencies](#dependencies)
 3. [Installation](#installation)
 4. [How to use](#how-to-use)
 5. [Thanks to](#thanks-to)

## Overview

> Gom is a simple mongodb library for golang, inspired by dbflex.

## Dependencies

> [Mongo Go Driver](https://github.com/mongodb/mongo-go-driver)
> [Toolkit](https://github.com/eaciit/toolkit)

## Installation

  ```go
    go get -u github.com/ariefsn/gom
  ```

## How to use

- Initialize

```go
  g := gom.NewGom()
```

- Set config

```go
  cfg := gom.MongoConfig{
    Host:     "localhost",
    Port:     "27017",
    Username: "",
    Password: "",
    Database: "test",
  }
```

- Init Gom

```go
  g.Init(cfg)
```

> That's it! Gom has ready to launch! :)
> Check full documentation

- Gom Filter

```go
  // Equal
  gom.Eq(<Field>, <Value>)

  // Not Equal
  gom.Ne(<Field>, <Value>)

  // Greater Than
  gom.Gt(<Field>, <Value>)

  // Greater Than Equal
  gom.Gte(<Field>, <Value>)

  // Less Than
  gom.Lt(<Field>, <Value>)

  // Less Than Equal
  gom.Lte(<Field>, <Value>)

  // Range
  gom.Range(<Field>, <From Value>, <To Value>)

  // Between
  gom.Between(<Field>, <From Value>, <To Value>)

  // In
  gom.In(<Field>, <Values...>)

  // Not In
  gom.Nin(<Field>, <Values...>)

  // Contains
  gom.Contains(<Field>, <Values...>)

  // Start With
  gom.StartWith(<Field>, <Values...>)

  // End With
  gom.EndWith(<Field>, <Values...>)

  // And
  gom.And(<Filters...>)

  // Or
  gom.And(<Filters...>)

  // Sort
  gom.Sort(<Field>, <SortType>)

```

- Gom Command
  - Get

  ```go
    res := []models.Hero{}
    err := g.Set().Table("hero").Result(&res).Cmd().Get()
    if err != nil {
      toolkit.Println(err.Error())
      return 0
    }
    for _, h := range res {
      toolkit.Println(h)
    }
  ```

  - Get One

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

  ```go
    hero := models.NewHero("Wolverine", "Hugh Jackman", 40)
    err := g.Set().Table("hero").Cmd().Insert(hero)
    if err != nil {
      toolkit.Println(err.Error())
      return
    }
  ```
  
  - Insert All

  ```go
    heroes := models.DummyData()
    err := g.Set().Table("hero").Cmd().InsertAll(&heroes)
    if err != nil {
      toolkit.Println(err.Error())
      return
    }
  ```

  - Update

  ```go
    hero := models.NewHero("Wonderwoman", "Gal Gadot", 34)
    err := g.Set().Table("hero").Filter(gom.Eq("RealName", "Scarlett Johansson")).Cmd().Update(hero)
    if err != nil {
      toolkit.Println(err.Error())
      return
    }
  ```

  - Delete

  ```go
    err := g.Set().Table("hero").Filter(gom.Eq("Name", "Batman")).Cmd().Delete()
    if err != nil {
      toolkit.Println(err.Error())
      return
    }
  ```

  - Sort

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

  - Filter

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

> [MongoDB](https://github.com/mongodb/)
> [Eaciit](https://github.com/eaciit/)
