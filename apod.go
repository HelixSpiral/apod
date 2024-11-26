package apod

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// NewAPOD returns an APOD object with the provided settings
// If no API Key is given we use the DEMO_KEY.
// If no domain is provided it uses the default NASA domain.
func NewAPOD(settings *NewAPODInput) *APOD {
	apiKey := "DEMO_KEY"
	if len(settings.APIKey) > 0 {
		apiKey = settings.APIKey
	}

	apodDomain := "https://api.nasa.gov/planetary/apod?api_key=%s"
	if len(settings.APODDomain) > 0 {
		apodDomain = settings.APODDomain
	}

	return &APOD{
		apiUrl: fmt.Sprintf(apodDomain, apiKey),
	}
}

// Query takes an ApodQueryInput and returns an slice of ApodQueryOutput
func (a *APOD) Query(queryParams *ApodQueryInput) ([]ApodQueryOutput, error) {
	var queryOutput []ApodQueryOutput

	// It doesn't matter what the value of thumbs is so we always include it.
	// If the user leaves it blank it defaults to false.
	queryUrl := fmt.Sprintf("%s&thumbs=%t", a.apiUrl, queryParams.Thumbs)

	if queryParams.Count != 0 {
		if !queryParams.Date.IsZero() || !queryParams.StartDate.IsZero() || !queryParams.EndDate.IsZero() {
			return queryOutput, fmt.Errorf("cannot use the following params with 'Count': 'Date', 'StartDate', 'EndDate'")
		}

		queryUrl += fmt.Sprintf("&count=%d", queryParams.Count)
	}

	if !queryParams.Date.IsZero() {
		if !queryParams.StartDate.IsZero() || !queryParams.EndDate.IsZero() {
			return queryOutput, fmt.Errorf("cannot use params 'Date' and 'StartDate' or 'EndDate' together")
		}

		queryUrl += fmt.Sprintf("&date=%s", queryParams.Date.Format("2006-01-02"))
	}

	if !queryParams.StartDate.IsZero() {
		queryUrl += fmt.Sprintf("&start_date=%s", queryParams.StartDate.Format("2006-01-02"))
	}

	if !queryParams.EndDate.IsZero() {
		queryUrl += fmt.Sprintf("&end_date=%s", queryParams.EndDate.Format("2006-01-02"))
	}

	// Do the actual query
	resp, err := http.Get(queryUrl)
	if err != nil {
		return queryOutput, err
	}
	defer resp.Body.Close()

	if resp.Header.Get("X-RateLimit-Remaining") == "0" {
		return queryOutput, fmt.Errorf("you have exceeded your rate limit")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return queryOutput, err
	}

	// Sometimes the header doesn't update properly so we have a second check for this here.
	if strings.Contains(string(body), "You have exceeded your rate limit.") {
		return queryOutput, fmt.Errorf("you have exceeded your rate limit")
	}

	// Since we're always returning an array we have to do a check for if it's not returning an array
	// We get the results, append to an empty array, and then return it back to the user.
	if !queryParams.Date.IsZero() {
		var queryOutputSingle ApodQueryOutput
		err = json.Unmarshal(body, &queryOutputSingle)
		queryOutput = append(queryOutput, queryOutputSingle)
	} else {
		err = json.Unmarshal(body, &queryOutput)
	}

	return queryOutput, err
}
