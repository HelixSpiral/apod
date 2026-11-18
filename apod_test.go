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
		t.Fatal("apod didn't return when not providing a domain")
	}

	inputConfig.APODDomain = "https://api.nasa.gov/planetary/apod"

	apodClient = apod.New(inputConfig)
	if apodClient == nil {
		t.Fatal("apod didn't return when providing a domain")
	}
}

func TestGet(t *testing.T) {
	inputConfig := &apod.Config{}

	apodClient := apod.New(inputConfig)
	date, _ := time.Parse("2006-01-02", "2022-02-01")

	_, err := apodClient.Get(date)
	if err != nil {
		t.Fatal(err)
	}
}

func TestQuery(t *testing.T) {
	inputConfig := &apod.Config{}

	apodClient := apod.New(inputConfig)
	date, _ := time.Parse("2006-01-02", "2022-02-01")

	queryInput := &apod.ApodQueryInput{
		DateFrom: date,
		DateTo:   date.Add((time.Hour * 24) * 5),
	}
	_, err := apodClient.Query(queryInput)
	if err != nil {
		t.Fatal(err)
	}
}
