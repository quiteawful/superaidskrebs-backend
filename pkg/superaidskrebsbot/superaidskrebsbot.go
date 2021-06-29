package superaidskrebsbot

import (
	"errors"
	"log"
	"net/mail"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/quiteawful/superaidskrebs-backend/pkg/config"
	tb "gopkg.in/tucnak/telebot.v2"
)

type SAKBot struct {
	bot *tb.Bot
	db  *sqlx.DB
}

func logMessageError(err error) {
	if err != nil {
		log.Println("Error sending message:", err)
	}
}

func validateEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func (s *SAKBot) sendErrorarguments(User *tb.User, errormessage string) {
	_, err := s.bot.Send(User, "Error:"+errormessage)
	logMessageError(err)
}

func (s *SAKBot) sendWrongarguments(User *tb.User, usage string) {
	s.sendErrorarguments(User, "Wrong amount of arguments. Usage:"+usage)
}

func (s *SAKBot) HandleRegister(m *tb.Message) {
	if !m.FromChannel() {
		mysplit := strings.Split(m.Text, " ")
		if len(mysplit) != 3 {
			s.sendWrongarguments(m.Sender, "/register emailadress password")
			return
		}
		if !validateEmail(mysplit[1]) {
			s.sendErrorarguments(m.Sender, "Email Address not Vaild")
			return
		}
	}
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

func CreateNewBot(config config.TelegramConf, db *sqlx.DB) (*SAKBot, error) {
	var sakbot SAKBot
	var err error
	sakbot.bot, err = tb.NewBot(tb.Settings{
		Token:  config.Token,
		Poller: &tb.LongPoller{Timeout: time.Duration(config.PollerTimeout) * time.Second},
	})
	sakbot.db = db
	if err != nil {
		return nil, errors.New("Error creating bot:" + err.Error())
	}

	sakbot.bot.Handle("/register", sakbot.HandleRegister)
	sakbot.bot.Handle(tb.OnText, sakbot.HandleMessages)
	sakbot.bot.Handle(tb.OnPhoto, sakbot.HandlePictures)

	return &sakbot, nil
}

func (s *SAKBot) Start() {
	s.bot.Start()
}
