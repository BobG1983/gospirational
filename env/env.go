package env

import (
	"github.com/gobuffalo/envy"
	l "github.com/shriekbob/gospirational/logging"
)

// Environment variables required for twitter login.  Available from dev.twitter.com.  Should be stored in a .env file or the environment
var (
	ConsumerKey       string
	ConsumerSecret    string
	AccessToken       string
	AccessTokenSecret string
)

func init() {
	// Get Env
	err := envy.Load()
	if err != nil {
		l.Log().Critical(err)
	}
	consumerKey, err := envy.MustGet("CONSUMER_KEY")
	if err != nil {
		l.Log().Critical(err)
	}
	consumerSecret, err := envy.MustGet("CONSUMER_SECRET")
	if err != nil {
		l.Log().Critical(err)
	}
	accessToken, err := envy.MustGet("ACCESS_TOKEN")
	if err != nil {
		l.Log().Critical(err)
	}
	accessTokenSecret, err := envy.MustGet("ACCESS_TOKEN_SECRET")
	if err != nil {
		l.Log().Critical(err)
	}

	ConsumerKey = consumerKey
	ConsumerSecret = consumerSecret
	AccessToken = accessToken
	AccessTokenSecret = accessTokenSecret
}
