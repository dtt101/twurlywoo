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
	api := getClient()
	fmt.Fprintf(w, "URLs from your tweets. Woo. %s! \n", r.URL.Path[1:])
	fmt.Fprintln(w, "Twitter client initialized", *api.Credentials)
	searchResult, err := api.GetSearch("golang", nil)
	if err != nil {
		panic(err)
	}
	fmt.Fprintln(w, searchResult.Statuses)
	for _, tweet := range searchResult.Statuses {
		fmt.Fprintln(w, tweet.Text)
	}
}

func main() {
	gotenv.Load()
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
