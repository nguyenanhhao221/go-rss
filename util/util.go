package util

import "time"

// ParsePubDate parse the time base of different format.
func ParsePubDate(pubDate string) (time.Time, error) {

	// Define the possible date formats
	dateFormats := []string{
		time.RFC1123Z,     // format used in Feed A
		"02 Jan 2006 MST", // format used in Feed 2
	}
	var t time.Time
	var err error
	// Try parsing the pubDate using different formats
	for _, format := range dateFormats {
		t, err := time.Parse(format, pubDate)
		if err == nil {
			return t, nil
		}
	}
	return t, err
}
