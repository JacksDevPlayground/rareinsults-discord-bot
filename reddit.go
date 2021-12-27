package main

import (
	"context"
	"math/rand"
	"time"

	"github.com/vartanbeno/go-reddit/v2/reddit"
)

var ctx = context.Background()
var MAX = 500

func random(max int) int {
	rand.Seed(time.Now().UnixNano())
	min := 0
	if max == -1 {
		max = MAX
	}

	return rand.Intn(max - min)
}

func getRareInsult() (cache []*reddit.Post, post reddit.Post, err error) {
	client, _ := reddit.NewReadonlyClient()
	posts, _, err := client.Subreddit.NewPosts(context.Background(), "rareinsults", &reddit.ListOptions{
		Limit: MAX,
	})

	if err != nil {
		return []*reddit.Post{}, reddit.Post{}, err
	}

	return posts, *posts[0], nil
}
