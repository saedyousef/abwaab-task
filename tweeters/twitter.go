package tweeters

import (
	"os"
	"github.com/joho/godotenv"
	"github.com/dghubble/oauth1"
	"github.com/dghubble/go-twitter/twitter"
)
type error interface {
    Error() string
}

func SearchTweets(query string) []string {

	var e []string
	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		er := append(e, err.Error())
		return er
	}

	config := oauth1.NewConfig(os.Getenv("CONSUMER_KEY"), os.Getenv("CONSUMER_SECRET"))
	token := oauth1.NewToken(os.Getenv("ACCESS_TOKEN"), os.Getenv("TOKEN_SECRET"))
	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter client
	client := twitter.NewClient(httpClient)

	// Search Tweets
	search, _, _ := client.Search.Tweets(&twitter.SearchTweetParams{
		Query: query,
		Count: 50,
		Lang: "en",
	})

	// Tweets Array.
	var tweets []string

	// Filling the tweets array.
	for _, tweet := range search.Statuses {
		tweets = append(tweets, tweet.Text, tweet.CreatedAt)
	}

	return tweets
}