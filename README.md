# Twurlywoo

Cuts to the end and just shows the links from your twitter timeline.

## Development

You need to have installed go and godeps.

Create a new twitter app via dev.twitter.com and copy the values you'll need for the .env below

Create a .env in the root of your project.

```
TWITTER_CONSUMER_KEY=MIND
TWITTER_CONSUMER_SECRET=YOUR
TWITTER_ACCESS_TOKEN=OWN
TWITTER_ACCESS_TOKEN_SECRET=BUSINESS
PORT=8080
```

Run with:

```
go run twurlywoo.go
```

Visit http://localhost:8080

## Production

Deployed to heroku at https://twurlywoo.herokuapp.com/
