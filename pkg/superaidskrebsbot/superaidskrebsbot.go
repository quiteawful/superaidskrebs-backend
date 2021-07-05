package superaidskrebsbot

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/quiteawful/superaidskrebs-backend/pkg/config"
	"github.com/quiteawful/superaidskrebs-backend/pkg/db/postgresmodel"
	tb "gopkg.in/tucnak/telebot.v2"
)

type SAKBot struct {
	bot       *tb.Bot
	db        *sqlx.DB
	filestore config.FilestoreConf
}

func logMessageError(err error) {
	if err != nil {
		log.Println("Error sending message:", err)
	}
}

func (s *SAKBot) sendErrorarguments(User *tb.User, errormessage string) {
	_, err := s.bot.Send(User, "Error:"+errormessage)
	logMessageError(err)
}

func (s *SAKBot) sendWrongarguments(User *tb.User, usage string) {
	s.sendErrorarguments(User, "Wrong amount of arguments. Usage:"+usage)
}

func (s *SAKBot) sendMessagetoUser(User *tb.User, message string) {
	_, err := s.bot.Send(User, message)
	logMessageError(err)
}

func (s *SAKBot) HandleMessages(m *tb.Message) {
	_, err := s.bot.Send(m.Sender, m.Text)
	logMessageError(err)
}

func (s *SAKBot) HandlePictures(m *tb.Message) {
	_, err := s.bot.Send(m.Sender, m.Caption)
	logMessageError(err)
	_, err = s.bot.Send(m.Sender, "Das ist ein Foto")
	logMessageError(err)
}

func CreateNewBot(config config.TelegramConf, db *sqlx.DB, filestoreconf config.FilestoreConf) (*SAKBot, error) {
	var sakbot SAKBot
	var err error
	sakbot.bot, err = tb.NewBot(tb.Settings{
		Token:  config.Token,
		Poller: &tb.LongPoller{Timeout: time.Duration(config.PollerTimeout) * time.Second},
	})
	if err != nil {
		return nil, errors.New("Error creating bot:" + err.Error())
	}
	sakbot.db = db
	sakbot.filestore = filestoreconf

	sakbot.bot.Handle("/user", sakbot.handleUserCommands)
	sakbot.bot.Handle(tb.OnText, sakbot.HandleMessages)
	sakbot.bot.Handle(tb.OnPhoto, sakbot.HandlePictures)

	return &sakbot, nil
}

func (s *SAKBot) RemovePads() {
	pp, err := postgresmodel.OnetimePads().All(context.Background(), s.db.DB)
	if err != nil {
		log.Printf("Ticker Error: Error getting OnetimePads: %v", err)
	}
	for _, p := range pp {
		if p.Time.After(time.Now()) {
			_, err := p.Delete(context.Background(), s.db.DB)
			if err != nil {
				log.Printf("Ticker Error: Error deleting OnetimePads: %v", err)
			}
		}
	}
}

func (s *SAKBot) RemovePadsTicker() {
	ticker := time.NewTicker(1 * time.Minute)
	for range ticker.C {
		s.RemovePads()
	}
}
func (s *SAKBot) Start() {
	s.RemovePads()
	go s.RemovePadsTicker()
	s.bot.Start()
}
