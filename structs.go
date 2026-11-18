package apod

import "time"

// Client is the main struct of our package
type Client struct {
	apiUrl string
}

// Config is the input for the New function
type Config struct {
	APODDomain string
}

// ApodQueryInput is the input for an Apod Query
type ApodQueryInput struct {
	DateFrom time.Time `json:"date_fom"`
	DateTo   time.Time `json:"date_to"`

	// Limit is the number of Pictures to return, this is rounded up to the nearest 25 pictures
	Limit int
}

// Picture is the information for a single APOD
type Picture struct {
	Date         string `json:"date,omitempty"`
	PostID       int    `json:"post_id,omitempty"`
	Title        string `json:"title,omitempty"`
	Permalink    string `json:"permalink,omitempty"`
	MediaType    string `json:"media_type,omitempty"`
	Explanation  string `json:"explanation,omitempty"`
	Url          string `json:"url,omitempty"`
	HdUrl        string `json:"hdurl,omitempty"`
	BasicHTML    string `json:"basic_html,omitempty"`
	BasicHTMLURL string `json:"basic_html_url,omitempty"`
}
