package apod_test

import (
	"testing"
	"time"

	"github.com/helixspiral/apod"
)

func TestNewAPOD(t *testing.T) {
	apodInput := &apod.NewAPODInput{}

	// We can't do a whole lot of checking for this function so we just do this to make sure they don't error.
	// If they don't, we're good for now.
	Apod := apod.NewAPOD(apodInput)
	if Apod == nil {
		t.Fatal("apod didn't return when not providing a key or domain")
	}

	apodInput.APIKey = "DEMO_KEY1"

	Apod = apod.NewAPOD(apodInput)
	if Apod == nil {
		t.Fatal("apod didn't return when providing a key, but no domain")
	}

	apodInput.APIKey = ""
	apodInput.APODDomain = "https://api.nasa.gov/planetary/apod?api_key"

	Apod = apod.NewAPOD(apodInput)
	if Apod == nil {
		t.Fatal("apod didn't return when providing a domain, but no key")
	}

	apodInput.APIKey = "DEMO_KEY1"
	apodInput.APODDomain = "https://api.nasa.gov/planetary/apod?api_key"

	Apod = apod.NewAPOD(apodInput)
	if Apod == nil {
		t.Fatal("apod didn't return when providing both a domain and a key")
	}
}

func TestQuery(t *testing.T) {
	apodInput := &apod.NewAPODInput{
		APIKey: "DEMO_KEY",
	}

	Apod := apod.NewAPOD(apodInput)
	date, _ := time.Parse("2006-01-02", "2022-02-01")

	// Test for Count with Date. Should error
	queryInput := &apod.ApodQueryInput{
		Date:  date,
		Count: 2,
	}
	_, err := Apod.Query(queryInput)
	if err == nil {
		t.Fatal("Should have errored with count and date, but didn't.")
	}

	// Test for Date with StartDate. Should error
	queryInput = &apod.ApodQueryInput{
		Date:      date,
		StartDate: date.Add(time.Hour * 24),
	}
	_, err = Apod.Query(queryInput)
	if err == nil {
		t.Fatal("Should have errored with Date and StartDate, but didn't.")
	}

	// Test for Date
	queryInput = &apod.ApodQueryInput{
		Date: date,
	}
	_, err = Apod.Query(queryInput)
	if err != nil {
		t.Fatal(err)
	}

	// Test for Start and End Dates
	queryInput = &apod.ApodQueryInput{
		StartDate: date,
		EndDate:   date.Add((time.Hour * 24) * 5),
	}
	_, err = Apod.Query(queryInput)
	if err != nil {
		t.Fatal(err)
	}

	// Test for count
	queryInput = &apod.ApodQueryInput{
		Count: 2,
	}
	_, err = Apod.Query(queryInput)
	if err != nil {
		t.Fatal(err)
	}
}
