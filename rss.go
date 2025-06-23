package main

import (
	"context"
	"net/http"
	"encoding/xml"
	"io"
	"html"
	"fmt"
	"time"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "gator")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	feed := RSSFeed{}
	err = xml.Unmarshal(data, &feed)
	if err != nil {
		return nil, err
	}

	// fmt.Println("PRE: ", feed.Channel.Description)
	feed.Channel.Title = html.UnescapeString(feed.Channel.Title)
	feed.Channel.Description = html.UnescapeString(feed.Channel.Description)
	// fmt.Println("POST: ", feed.Channel.Description)

	for i, item := range feed.Channel.Item {
	// fmt.Println("PRE: ", item.Description)
		item.Title = html.UnescapeString(item.Title)
		item.Description = html.UnescapeString(item.Description)
		// Porque ?? Range me da copia ??!!!
		feed.Channel.Item[i] = item
	// fmt.Println("POST: ", item.Description)
	}

	return &feed, nil
}

func handlerAgg(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("Usage: agg <duration>")
	}

	duration, err := time.ParseDuration(cmd.args[0])
	if err != nil {
		return fmt.Errorf("Error parsing duration: %s err: %w", cmd.args[0], err)
	}
	fmt.Printf("Collecting feeds every %s\n", duration)
	fmt.Printf("Now: %v\n", time.Now())

	ticker := time.NewTicker(duration)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}


	return nil
}



