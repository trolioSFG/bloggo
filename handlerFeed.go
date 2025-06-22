package main

import (
	"context"
	"fmt"
	"github.com/google/uuid"
//	"database/sql"
	"github.com/trolioSFG/database"
	"time"
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

