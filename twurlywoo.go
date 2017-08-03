package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/ChimeraCoder/anaconda"
	"github.com/subosito/gotenv"
)

func getClient() *anaconda.TwitterApi {
	anaconda.SetConsumerKey(os.Getenv("TWITTER_CONSUMER_KEY"))
	anaconda.SetConsumerSecret(os.Getenv("TWITTER_CONSUMER_SECRET"))
	api := anaconda.NewTwitterApi(
		os.Getenv("TWITTER_ACCESS_TOKEN"),
		os.Getenv("TWITTER_ACCESS_TOKEN_SECRET"))
	return api
}

func getTweetUrls(tweet *anaconda.Tweet) string {
	tweets := ""
	for _, url := range tweet.Entities.Urls {
		if url.Expanded_url != "" {
			if !strings.HasPrefix(url.Expanded_url, "https://twitter.com") {
				tweets += fmt.Sprintf("<a href=\"%s\">%[1]s</a>\n", url.Expanded_url)
				tweets += "<br>"
			}
		}
	}
	return tweets
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	api := getClient()
	fmt.Fprintf(w, "<p>URLs from my tweets. There are no words. Woo.</p>")
	homeResult, err := api.GetHomeTimeline(nil)
	if err != nil {
		panic(err)
	}
	for _, tweet := range homeResult {
		fmt.Fprint(w, getTweetUrls(&tweet))
		if tweet.QuotedStatus != nil {
			fmt.Fprint(w, getTweetUrls(tweet.QuotedStatus))
		} else if tweet.RetweetedStatus != nil {
			fmt.Fprint(w, getTweetUrls(tweet.RetweetedStatus))
		}
	}
}

func main() {
	gotenv.Load()
	http.HandleFunc("/", handler)
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
