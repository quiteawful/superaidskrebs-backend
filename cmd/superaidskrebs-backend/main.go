package main

import (
	"flag"
	"log"

	"github.com/quiteawful/superaidskrebs-backend/internal/hashing"
	"github.com/quiteawful/superaidskrebs-backend/internal/mailer"
	"github.com/quiteawful/superaidskrebs-backend/pkg/config"
	"github.com/quiteawful/superaidskrebs-backend/pkg/db"
	"github.com/quiteawful/superaidskrebs-backend/pkg/superaidskrebsbot"
)

var confpath string

func init() {
	flag.StringVar(&confpath, "c", "config.toml", "Path to the config for the server")
}

func main() {
	conf, err := config.LoadConfig(confpath)
	if err != nil {
		log.Fatalln("Error loading config:", err.Error())
	}

	db, err := db.Initialisation(conf.Database)
	if err != nil {
		log.Fatal("Error initalizing Database:", err.Error())
	}

	hashing.SetParams(conf.Argon)
	mailer.SetParams(conf.Mail)

	bot, err := superaidskrebsbot.CreateNewBot(conf.Telegram, db, conf.Filestore)
	if err != nil {
		log.Fatalln("Error creating new bot:", err.Error())
	}
	log.Println("Starting Bot")
	bot.Start()
}
