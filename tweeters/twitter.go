package tweeters

import (
    //"os"
	"io/ioutil"
	//"encoding/json"
	// "path/filepath"
	"github.com/dghubble/oauth1"
	"github.com/dghubble/go-twitter/twitter"
)

type error interface {
    Error() string
}

func SearchTweets(query string) string {

	// // Open json file to get keys later on.
	// jsonFile, err := os.Open("/keys.json")
	
	// // Check if there is an error.
	// if err != nil {
    //     return err.Error()
	// }
	
    // // defer the closing of our jsonFile so that we can parse it later on
	// defer jsonFile.Close()
	
	// // read our opened xmlFile as a byte array.
	// var keys map[string]interface{}
	// jsonKeys, _ := ioutil.ReadAll(jsonFile)

	// json.Unmarshal([]byte(jsonKeys), &keys)

	// config := oauth1.NewConfig(keys["consumer_key"].(string), keys["consumer_secret"].(string))
	// token := oauth1.NewToken(keys["token"].(string), keys["token_secret"].(string))
	// httpClient := config.Client(oauth1.NoContext, token)

	config := oauth1.NewConfig("r5OdTex1cn0cez6x6dJWNAMgS", "Ib9q2rj7CTPDLm2VPYPE27bYrciSoYoBnJmg9qv08geocffNDz")
	token := oauth1.NewToken("1322735486-IwPcAxffIBhWeHZT1XldW3VqD4LNdGrUqMF3ewA", "Quhyx9rtjiRB45NoGQmyWK8L6zq4Jxb3zHKxDHSMX5SWZ")
	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter client
	client := twitter.NewClient(httpClient)

	// Search Tweets
	_ , resp, err := client.Search.Tweets(&twitter.SearchTweetParams{
		Query: "elon",
		Count: 50,
	})
	if err != nil {
		return err.Error()
	}


	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err.Error()
	}
	defer resp.Body.Close()
	return string(body)
}