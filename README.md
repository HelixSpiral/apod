Astronomy Picture of the Day API Wrapper
---

This is a API wrapper written in Golang for the [Astronomy Picture of the Day](https://apod.nasa.gov/apod/astropix.html) service that NASA hosts.

Usage
---

Basic date query:
```go
    apodCfg := &apod.Config{
        APIKey: "DEMO_KEY",
    }

    apodClient := apod.New(apodCfg)
    date, _ := time.Parse("2006-01-02", "2022-02-11")
    queryInput := &apod.ApodQueryInput{
        Date: date,
    }

    resp, err := apodClient.Query(queryInput)
    if err != nil {
        panic(err)
    }

    fmt.Println(resp)
```

You can provide a start and end date, with end date defaulting to the current date
```go
    apodCfg := &apod.Config{
        APIKey: "DEMO_KEY",
    }

    apodClient := apod.New(apodCfg)

    date, _ := time.Parse("2006-01-02", "2022-02-01")
    queryInput := &apod.ApodQueryInput{
        StartDate: date,
        EndDate: date.Add((time.Hour*24) * 5),
    }
    
    resp, err := apodClient.Query(queryInput)
    if err != nil {
        panic(err)
    }

    fmt.Println(resp)
```

Providing a count gives you that many random selections
```go
    apodCfg := &apod.Config{
        APIKey: "DEMO_KEY",
    }

    apodClient := apod.New(apodCfg)

    queryInput := &apod.ApodQueryInput{
        Count: 5,
    }
    
    resp, err := apodClient.Query(queryInput)
    if err != nil {
        panic(err)
    }

    fmt.Println(resp)
```
