package main

import (
	"fmt"
	"time"
	"unicode/utf8"

	"github.com/ChimeraCoder/anaconda"
	"github.com/ShriekBob/goismatic"
	"github.com/jasonlvhit/gocron"
	"github.com/shriekbob/gospirational/env"
	l "github.com/shriekbob/gospirational/logging"
)

var hashTag = "#InspirationalQuotes"
var maxQuoteLength = 280 - utf8.RuneCountInString(" "+hashTag)

func tweet() {
	// Twitter Auth
	anaconda.SetConsumerKey(env.ConsumerKey)
	anaconda.SetConsumerSecret(env.ConsumerSecret)
	twitter := anaconda.NewTwitterApi(env.AccessToken, env.AccessTokenSecret)

	// Get quote
	quote, err := goismatic.Get(goismatic.English)
	if err != nil {
		l.Log().Error(err)
	}
	if quote != nil {
		quoteLen := utf8.RuneCountInString(quote.String())

		// If it's too long, spin, then try again
		for quoteLen > maxQuoteLength {
			time.Sleep(time.Second * 10)
			quote, err := goismatic.Get(goismatic.English)
			if err != nil {
				l.Log().Error(err)
			}
			if quote != nil {
				quoteLen = utf8.RuneCountInString(quote.String())
			}
		}

		// Post quote to twitter
		body := quote.String() + " " + hashTag
		tweet, err := twitter.PostTweet(body, nil)
		if err != nil {
			l.Log().Error(err)
		}

		l.Log().Debug(tweet.Id)
	}
}

func main() {
	gocron.Every(1).Day().At("10:30").Do(tweet)
	gocron.Every(1).Day().At("13:00").Do(tweet)
	gocron.Every(1).Day().At("17:00").Do(tweet)

	_, time := gocron.NextRun()
	fmt.Println(time)

	<-gocron.Start()
}
