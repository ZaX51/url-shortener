package main

import (
	"github.com/ZaX51/url-shortener/internal/app"
	"github.com/spf13/viper"
)

func main() {
	viper.SetDefault("DOMAIN", "localhost:3000")
	viper.SetDefault("URL_LENGTH", 7)
	viper.SetDefault("REDIS_ADDR", "localhost:6379")
	viper.SetDefault("REDIS_EXPIRATION", "20s")

	viper.AutomaticEnv()

	app := app.New(app.Config{
		Domain:          viper.GetString("DOMAIN"),
		UrlLength:       viper.GetInt("URL_LENGTH"),
		RedisAddr:       viper.GetString("REDIS_ADDR"),
		RedisExpiration: viper.GetDuration("REDIS_EXPIRATION"),
	})

	app.Run(":3000")
}
