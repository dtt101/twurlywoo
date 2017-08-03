package main

import (
	"fmt"
	"net/http"
	"os"

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

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	api := getClient()
	fmt.Fprintf(w, "<p>URLs from my tweets. There are no words. Woo.</p>")
	homeResult, err := api.GetHomeTimeline(nil)
	if err != nil {
		panic(err)
	}
	for _, tweet := range homeResult {
		for _, url := range tweet.Entities.Urls {
			if url.Expanded_url != "" {
				fmt.Fprintf(w, "<a href=\"%s\">%[1]s</a>\n", url.Expanded_url)
				fmt.Fprintln(w, "<br>")
			}
		}
	}
}

func main() {
	gotenv.Load()
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
