package main

import (
	"log"

	"github.com/quiteawful/superaidskrebs-backend/pkg/config"
	"github.com/quiteawful/superaidskrebs-backend/pkg/superaidskrebsbot"
)

func main() {
	conf, err := config.LoadConfig("config.toml")
	if err != nil {
		log.Fatalln("Error loading config:", err.Error())
	}

	bot, err := superaidskrebsbot.CreateNewBot(conf.Telegram)
	if err != nil {
		log.Fatalln("Error creating new bot:", err.Error())
	}
	bot.Start()
}
