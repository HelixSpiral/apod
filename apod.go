package apod

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// New returns an APOD object with the provided settings
// If no domain is provided it uses the default NASA domain.
func New(settings *Config) *Client {
	apodDomain := "https://science.nasa.gov/wp-json/wp/v2/apod-basic"
	if len(settings.APODDomain) > 0 {
		apodDomain = settings.APODDomain
	}

	return &Client{
		apiUrl: apodDomain,
	}
}

// Get returns a single APOD response
func (a *Client) Get(date time.Time) (Picture, error) {
	var apod Picture

	queryURL := fmt.Sprintf("%s/%s", a.apiUrl, date.Format("060102"))

	resp, err := http.Get(queryURL)
	if err != nil {
		return apod, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return apod, err
	}

	err = json.Unmarshal(body, &apod)

	return apod, err
}

// Query takes an ApodQueryInput and returns an slice of Pictures
func (a *Client) Query(queryParams *ApodQueryInput) ([]Picture, error) {
	var queryOutput []Picture

	endDate := queryParams.DateTo
	if endDate.IsZero() {
		endDate = time.Now()
	}

	baseURL := fmt.Sprintf("%s?date_from=%s&date_to=%s", a.apiUrl, queryParams.DateFrom.Format("060102"), endDate.Format("060102"))

	for x, page_empty := 1, false; page_empty != true && len(queryOutput) < queryParams.Limit; x++ {
		// Do the actual query
		resp, err := http.Get(fmt.Sprintf("%s&page=%d", baseURL, x))
		if err != nil {
			return queryOutput, err
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return queryOutput, err
		}

		var pageOutput []Picture
		err = json.Unmarshal(body, &pageOutput)
		if len(pageOutput) == 0 {
			page_empty = true

			break
		}

		queryOutput = append(queryOutput, pageOutput...)
	}

	return queryOutput, nil
}
