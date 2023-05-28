package main

import (
	"encoding/xml"
	"io"
	"net/http"
	"time"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"tilte"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Language    string    `xml:"language"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

// urlToFeed create http client and make the request to the url provided and get the rss feed
func urlToFeed(url string) (RSSFeed, error) {
	httpClient := http.Client{
		Timeout: 10 * time.Second,
	}
	res, err := httpClient.Get(url)
	if err != nil {
		return RSSFeed{}, err
	}
	// By using defer here, we guarantee that this close function will always get to run last
	// Because other part of this function maybe can throw an error and panic, without defer the Close
	// function might not get called.
	// It a good practice to do this to ensure that this always get called and prevent potential issue
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)

	if err != nil {
		return RSSFeed{}, err
	}

	rssFeed := RSSFeed{}
	xmlMarshalErr := xml.Unmarshal(data, &rssFeed)
	if xmlMarshalErr != nil {
		return RSSFeed{}, xmlMarshalErr
	}
	return rssFeed, nil
}
