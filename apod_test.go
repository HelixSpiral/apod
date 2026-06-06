package apod_test

import (
	"testing"
	"time"

	"github.com/helixspiral/apod"
)

func TestNew(t *testing.T) {
	inputConfig := &apod.Config{}

	// We can't do a whole lot of checking for this function so we just do this to make sure they don't error.
	// If they don't, we're good for now.
	apodClient := apod.New(inputConfig)
	if apodClient == nil {
		t.Fatal("apod didn't return when not providing a key or domain")
	}

	inputConfig.APIKey = "DEMO_KEY1"

	apodClient = apod.New(inputConfig)
	if apodClient == nil {
		t.Fatal("apod didn't return when providing a key, but no domain")
	}

	inputConfig.APIKey = ""
	inputConfig.APODDomain = "https://api.nasa.gov/planetary/apod?api_key"

	apodClient = apod.New(inputConfig)
	if apodClient == nil {
		t.Fatal("apod didn't return when providing a domain, but no key")
	}

	inputConfig.APIKey = "DEMO_KEY1"
	inputConfig.APODDomain = "https://api.nasa.gov/planetary/apod?api_key"

	apodClient = apod.New(inputConfig)
	if apodClient == nil {
		t.Fatal("apod didn't return when providing both a domain and a key")
	}
}

func TestQuery(t *testing.T) {
	inputConfig := &apod.Config{
		APIKey: "DEMO_KEY",
	}

	apodClient := apod.New(inputConfig)
	date, _ := time.Parse("2006-01-02", "2022-02-01")

	// Test for Count with Date. Should error
	queryInput := &apod.ApodQueryInput{
		Date:  date,
		Count: 2,
	}
	_, err := apodClient.Query(queryInput)
	if err == nil {
		t.Fatal("Should have errored with count and date, but didn't.")
	}

	// Test for Date with StartDate. Should error
	queryInput = &apod.ApodQueryInput{
		Date:      date,
		StartDate: date.Add(time.Hour * 24),
	}
	_, err = apodClient.Query(queryInput)
	if err == nil {
		t.Fatal("Should have errored with Date and StartDate, but didn't.")
	}

	// Test for Date
	queryInput = &apod.ApodQueryInput{
		Date: date,
	}
	_, err = apodClient.Query(queryInput)
	if err != nil {
		t.Fatal(err)
	}

	// Test for Start and End Dates
	queryInput = &apod.ApodQueryInput{
		StartDate: date,
		EndDate:   date.Add((time.Hour * 24) * 5),
	}
	_, err = apodClient.Query(queryInput)
	if err != nil {
		t.Fatal(err)
	}

	// Test for count
	queryInput = &apod.ApodQueryInput{
		Count: 2,
	}
	_, err = apodClient.Query(queryInput)
	if err != nil {
		t.Fatal(err)
	}
}
