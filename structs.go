package apod

import "time"

// APOD is the main struct of our package
type APOD struct {
	apiUrl string
}

// ApodQueryInput is the input for an Apod Query
type ApodQueryInput struct {
	Date      time.Time `json:"date"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
	Count     int       `json:"count"`
	Thumbs    bool      `json:"thumbs"`
}

// ApodQueryOutput is the output from an Apod Query
type ApodQueryOutput struct {
	Title          string `json:"title,omitempty"`
	Explanation    string `json:"explanation,omitempty"`
	Date           string `json:"date,omitempty"`
	MediaType      string `json:"media_type,omitempty"`
	Url            string `json:"url,omitempty"`
	HdUrl          string `json:"hdurl,omitempty"`
	ThumbnailUrl   string `json:"thumbnail_url,omitempty"`
	Copyright      string `json:"Copyright,omitempty"`
	ServiceVersion string `json:"service_version,omitempty"`
}
