Astronomy Picture of the Day API Wrapper
---

This is a API wrapper written in Golang for the [Astronomy Picture of the Day](https://apod.nasa.gov/apod/astropix.html) service that NASA hosts.

Usage
---

Basic date query:
```go
	Apod := apod.NewAPOD("DEMO_KEY")
	date, _ := time.Parse("2006-01-02", "2022-02-11")
    queryInput := &apod.ApodQueryInput{
		Date: date,
	}

    resp, err := Apod.Query(queryInput)
    if err != nil {
        panic(err)
    }

    fmt.Println(resp)
```

You can provide a start and end date, with end date defaulting to the current date
```go
	Apod := apod.NewAPOD("DEMO_KEY")
	date, _ := time.Parse("2006-01-02", "2022-02-01")
    queryInput := &apod.ApodQueryInput{
		StartDate: date,
        EndDate: date.Add((time.Hour*24) * 5)),
	}
    
    resp, err := Apod.Query(queryInput)
    if err != nil {
        panic(err)
    }

    fmt.Println(resp)
```

Providing a count gives you that many random selections
```go
	Apod := apod.NewAPOD("DEMO_KEY")
    queryInput := &apod.ApodQueryInput{
		Count: 5,
	}
    
    resp, err := Apod.Query(queryInput)
    if err != nil {
        panic(err)
    }

    fmt.Println(resp)
```