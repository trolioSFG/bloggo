package main

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"database/sql"
	"github.com/trolioSFG/database"
	"time"
	"strconv"
	"strings"
)

func handlerAddFeed(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 2 {
		return fmt.Errorf("Usage: addfeed <feed_name> <feed_url>")
	}


	data := database.CreateFeedParams {
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name: cmd.args[0],
		Url: cmd.args[1],
		UserID: user.ID,
	}

	feed, err := s.db.CreateFeed(context.Background(), data)
	if err != nil {
		return err
	}

	fmt.Println(feed)

	aux := command {
		name: "follow",
		args: []string{feed.Url},
	}

	err = handlerAddFeedFollow(s, aux, user)
	if err != nil {
		return fmt.Errorf("Error adding follow after add %w", err)
	}

	return nil
}

func handlerFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return err
	}

	for _, feed := range feeds {

		user, err := s.db.GetUserByID(context.Background(), feed.UserID)
		if err != nil {
			return nil
		}

		fmt.Println("Feed name:", feed.Name)
		fmt.Println("Feed URL:", feed.Url)
		fmt.Println("Creator:", user.Name)
	}

	return nil
}

func handlerAddFeedFollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("Usage follow <url>")
	}

	feed, err := s.db.GetFeedByURL(context.Background(), cmd.args[0])
	if err != nil {
		return err
	}


	data := database.CreateFeedFollowParams {
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID: user.ID,
		FeedID: feed.ID,
	}

	row, err := s.db.CreateFeedFollow(context.Background(), data)
	if err != nil {
		return err
	}

	fmt.Println("Feed name:", row.FeedName, " followed by:", row.UserName)
	return nil
}

func handlerFollowing(s *state, cmd command, user database.User) error {

	data, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return err
	}

	fmt.Println("Followed feeds for", user.Name)
	for _, followed := range data {
		fmt.Println(followed.FeedName)
	}

	return nil
}

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) < 1 {
		return fmt.Errorf("Usage: unfollow <feed_url>")
	}


	feed, err := s.db.GetFeedByURL(context.Background(), cmd.args[0])
	if err != nil {
		return err
	}

	err = s.db.DeleteFeedFollow(context.Background(), database.DeleteFeedFollowParams{UserID: user.ID, FeedID: feed.ID})
	if err != nil {
		return err
	}

	fmt.Println("Feed follow deleted!")
	return nil
}

func scrapeFeeds(s *state) error {
	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return err
	}
	fmt.Println("Feed ID to mark:", feed.ID)
	err = s.db.MarkFeedFetched(context.Background(),
		database.MarkFeedFetchedParams {
			LastFetchedAt: sql.NullTime{ Time: time.Now(), Valid: true },
			UpdatedAt: time.Now(),
			ID: feed.ID,
		})
	
	if err != nil {
		return err
	}

	fmt.Println("Marked")
	fmt.Println("Fetching feed...")

	rss, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		return err
	}

	for _, item := range rss.Channel.Item {
		fmt.Println("Saving:", item.Title)

		// OJO a esto...
		pubDate, _ := time.Parse(time.ANSIC, item.PubDate)

		err := s.db.AddPost(context.Background(), database.AddPostParams{ ID: uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Title: item.Title,
			Url: item.Link, 
			Description: sql.NullString{ String: item.Description, Valid: true },
			PublishedAt: pubDate, 
			FeedID: feed.ID,
		})

		if err != nil {
			// No hay mejor forma ??
			if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
				continue
			}			
			return fmt.Errorf("INSERT POST ERROR: %w %v\n", err, err)
		}

	}

	fmt.Println("======================================================")
	return nil

}

func handlerBrowse(s *state, cmd command, user database.User) error {
	if len(cmd.args) == 0 {
		cmd.args = append(cmd.args, "2")
	}

	numposts, err := strconv.Atoi(cmd.args[0])
	if err != nil {
		return fmt.Errorf("Error parsing int limit: %w", err)
	}


	posts, err := s.db.GetPostsForUser(context.Background(),
		database.GetPostsForUserParams { ID: user.ID, Limit: int32(numposts), })
	if err != nil {
		return err
	}

	for _, post := range posts {
		// Layouts must use the reference time Mon Jan 2 15:04:05 MST 2006
		// ... .Format("Mon Jan 2") !!!

		fmt.Println("Published at:", post.PublishedAt.Format(time.ANSIC))
		fmt.Println(post.Title)
		// Se guarda la struct entera al ser sql.NullString !!!
		fmt.Println(post.Description.String)
	}

	return nil
}
