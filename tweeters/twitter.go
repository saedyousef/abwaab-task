package tweeters

import (
	"github.com/dghubble/oauth1"
	"github.com/dghubble/go-twitter/twitter"
)


func SearchTweets(query string) []string {

	config := oauth1.NewConfig("r5OdTex1cn0cez6x6dJWNAMgS", "Ib9q2rj7CTPDLm2VPYPE27bYrciSoYoBnJmg9qv08geocffNDz")
	token := oauth1.NewToken("1322735486-IwPcAxffIBhWeHZT1XldW3VqD4LNdGrUqMF3ewA", "Quhyx9rtjiRB45NoGQmyWK8L6zq4Jxb3zHKxDHSMX5SWZ")
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