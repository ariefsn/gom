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

- Import

  ```go
    import (
      "github.com/ariefsn/gom"
    )
  ```

- Create Instance

  ```go
    g := gom.NewGom()
  ```

- Set Config

  ```go
    cfg := gom.Config{
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

- Check

  > You can check connection with `CheckClient()`

  ```go
    err := g.CheckClient()

    if err != nil {
      toolkit.Println(toolkit.Sprintf("Connection Error: %s", err.Error()))
    }
  ```

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
  > You can choose want to use chain mode or with set params.
  > If you want to use chain mode, simply give `nil` value to Set. eg: `Set(nil)`
  > Or want to use set params, pass it to Set. eg `Set(&gom.SetParams{})`
  > Both of them are give the same result :)

  ```go
    SetParams {
      TableName string          // name of collection/table (required)
      Result    interface{}     // result (optional)
      Filter    *Filter         // gom filter (optional)
      Pipe      []bson.M        // pipe (optional)
      SortField string          // sort by field (optional)
      SortBy    string          // sort by asc/desc (optional)
      Skip      int             // skip result (optional)
      Limit     int             // limit result (optional)
      Timeout   time.Duration   // context timeout (optional)
    }
  ```

  > If Timeout not set it will give 30 second by default :)

  - **Get**
    > Get all data. It'll use Filter as default. if pipe not null then Filter will be ignored. This command returns countFilterData `int64`, countAllData `int64`, and `error`

    ```go
      res := []models.Hero{}

      // Chain
      countFilterData, countAllData, err := g.Set(nil).Table("hero").Timeout(10).Result(&res).Cmd().Get()

      // Use Set Params
      countFilterData, countAllData, err = g.Set(&gom.SetParams{
        TableName: "hero",
        Result:    &res,
        Timeout:   10,
      }).Cmd().Get()

      if err != nil {
        toolkit.Println(err.Error())
        return 0
      }

      toolkit.Println("Data found", countFilterData, "of", countAllData )

      for _, h := range res {
        toolkit.Println(h)
      }
    ```

  - **Get One**
    > Get one data. It'll use Filter as default, pipe ignored. This command return `error`.

    ```go
      res := models.Hero{}

      // Chain
      err := g.Set(nil).Table("hero").Timeout(10).Result(&res).Cmd().GetOne()

      // Use Set Params
      err = g.Set(&gom.SetParams{
        TableName: "hero",
        Result:    &res,
      Timeout:   10,
      }).Cmd().GetOne()

      if err != nil {
        toolkit.Println(err.Error())
        return
      }

      toolkit.Println(res)
    ```

  - **Insert**
    > Insert one data, for multiple data use InsertAll. This command returns insertedID `interface{}` and `error`.

    ```go
      hero := models.NewHero("Wolverine", "Hugh Jackman", 40)

      // Chain
      _, err := g.Set(nil).Table("hero").Timeout(10).Cmd().Insert(hero)

      // Use Set Params
    _, err = g.Set(&gom.SetParams{
      TableName: "hero",
      Timeout:   10,
      }).Cmd().Insert(hero)

      if err != nil {
        toolkit.Println(err.Error())
        return
      }
    ```
  
  - **Insert All**
    > Insert multiple data. This command returns insertedIDs `[]interface{}` and `error`

    ```go
      heroes := models.DummyData()

      // Chain
      _, err := g.Set(nil).Table("hero").Timeout(10).Cmd().InsertAll(&heroes)

      // Use Set Params
      _, err = g.Set(&gom.SetParams{
        TableName: "hero",
        Timeout:   10,
      }).Cmd().InsertAll(&heroes)

      if err != nil {
        toolkit.Println(err.Error())
        return
      }
    ```

  - **Update**
    > Update data with filter, pipe will ignored. This command return `error`.

    ```go
      hero := models.NewHero("Wonderwoman", "Gal Gadot", 34)

      // Chain
      err := g.Set(nil).Table("hero").Timeout(10).Filter(gom.Eq("RealName", "Scarlett Johansson")).Cmd().Update(hero)

      // Use Set Params
      err = g.Set(&gom.SetParams{
        TableName: "hero",
        Filter:    gom.Eq("RealName", "Scarlett Johansson"),
        Timeout:   10,
      }).Cmd().Update(hero)

      if err != nil {
        toolkit.Println(err.Error())
        return
      }
    ```

  - **Delete One**
    > Delete one data with filter, pipe will ignored. This command return `error`.

    ```go
      // Chain
      err := g.Set(nil).Table("hero").Timeout(10).Filter(gom.Eq("Name", "Batman")).Cmd().DeleteOne()

      // Use Set Params
      err = g.Set(&gom.SetParams{
        TableName: "hero",
        Filter:    gom.Eq("Name", "Batman"),
        Timeout:   10,
      }).Cmd().DeleteOne()

      if err != nil {
        toolkit.Println(err.Error())
        return
      }
    ```

  - **Delete All**
    > Delete all data with filter, pipe will ignored. This command return totalDeletedDocuments `int64` and `error`.

    ```go
      // Chain
      _, err := g.Set(nil).Table("hero").Filter(gom.EndWith("Name", "man")).Cmd().DeleteAll()

      // Use Set Params
      _, err = g.Set(&gom.SetParams{
        TableName: "hero",
        Filter:    gom.EndWith("Name", "man"),
        Timeout:   10,
      }).Cmd().DeleteAll()

      if err != nil {
        toolkit.Println(err.Error())
        return
      }
    ```

  - **Sort**
    > Sort results ascending or descending

    ```go
      res := []models.Hero{}

      // Chain
      _, _, err := g.Set(nil).Table("hero").Timeout(10).Result(&res).Sort("RealName", "asc").Cmd().Get()

      // Use Set Params
      _, _, err = g.Set(&gom.SetParams{
        TableName: "hero",
        Result:    &res,
        SortField: "RealName",
        SortBy:    "asc",
        Timeout:   10,
      }).Cmd().Get()

      if err != nil {
        toolkit.Println(err.Error())
      }

      for _, h := range res {
        toolkit.Println(h.RealName, "=>", h.Name, "=>", h.Age)
      }
    ```

  - **Skip & Limit**
    > Set skip & limit for results

    ```go
      res := []models.Hero{}

      // Chain
      err := g.Set(nil).Table("hero").Result(&res).Timeout(10).Skip(0).Limit(3).Cmd().Get()

      // Use Set Params
      err = g.Set(&gom.SetParams{
        TableName: "hero",
        Result:    &res,
        Skip:      0,
        Limit:     3,
        Timeout:   10,
      }).Cmd().Get()

      if err != nil {
        toolkit.Println(err.Error())
        return
      }

      for _, h := range res {
        toolkit.Println(h)
      }
    ```

  - **Filter**
    > Set filter data

    ```go
      res := []models.Hero{}
      filter := gom.And(gom.Eq("Age", 45), gom.StartWith("Name", "A"))

      // Chain
      _, _, err := g.Set(nil).Table("hero").Timeout(10).Result(&res).Filter(filter).Cmd().Get()

      // Use Set Params
      _, _, err = g.Set(&gom.SetParams{
        TableName: "hero",
        Result:    &res,
        Filter:    filter,
        Timeout:   10,
      }).Cmd().Get()

      if err != nil {
        toolkit.Println(err.Error())
        return
      }

      for _, h := range res {
        toolkit.Println(h)
      }
    ```

  - **Pipe**
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

      // Chain
      _, _, err = g.Set(nil).Table("hero").Result(&res).Timeout(10).Pipe(pipe).Cmd().Get()

      // Use Set Params
      _, _, err = g.Set(&gom.SetParams{
        TableName: "hero",
        Result:    &res,
        Pipe:      pipe,
        Timeout:   10,
      }).Cmd().Get()

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
