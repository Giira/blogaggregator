package main

import (
	"context"
	"database/sql"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"net/http"
	"time"

	"github.com/Giira/blogaggregator/internal/database"
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
	client := http.DefaultClient
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var feed RSSFeed
	err = xml.Unmarshal(body, &feed)
	if err != nil {
		return nil, err
	}

	return &feed, nil
}

func scrapeFeeds(s *state) error {
	db_url, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return fmt.Errorf("error: %v", err)
	}

	feed_details, err := s.db.GetFeed(context.Background(), db_url)
	if err != nil {
		return fmt.Errorf("error: %v", err)
	}

	s.db.MarkFeedFetched(context.Background(), database.MarkFeedFetchedParams{
		LastFetchedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
		ID: feed_details.ID,
	})

	feed, err := fetchFeed(context.Background(), db_url)
	if err != nil {
		return fmt.Errorf("error: %v", err)
	}

	title := html.UnescapeString(feed.Channel.Title)
	fmt.Printf("Title: %v\n", title)
	fmt.Printf("Link: %v\n", feed.Channel.Link)
	desc := html.UnescapeString(feed.Channel.Description)
	fmt.Printf("Description: %v\n", desc)
	for i, item := range feed.Channel.Item {
		fmt.Printf("\nItem %v:\n\n", i+1)
		title = html.UnescapeString(item.Title)
		fmt.Printf("Title: %v\n", title)
		fmt.Printf("Link: %v\n", item.Link)
		desc = html.UnescapeString(item.Description)
		fmt.Printf("Description: %v\n", desc)
		fmt.Printf("Publication Date: %v\n", item.PubDate)
	}
	return nil
}
