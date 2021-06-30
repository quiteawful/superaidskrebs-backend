package superaidskrebsbot

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/mail"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/quiteawful/superaidskrebs-backend/internal/hashing"
	"github.com/quiteawful/superaidskrebs-backend/internal/mailer"
	"github.com/quiteawful/superaidskrebs-backend/pkg/db"
	"github.com/quiteawful/superaidskrebs-backend/pkg/db/postgresmodel"
	"github.com/volatiletech/sqlboiler/v4/boil"
	tb "gopkg.in/tucnak/telebot.v2"
)

func validateEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func (s *SAKBot) doesUserExist(username string) (bool, error) {
	uexists, err := postgresmodel.Users(postgresmodel.UserWhere.Username.EQ(username)).Exists(context.Background(), s.db.DB)
	if err != nil {
		return uexists, errors.New("Error querying for User: " + err.Error())
	}
	return uexists, nil
}

func (s *SAKBot) handleUserCommands(m *tb.Message) {
	if !m.FromChannel() {
		mysplit := strings.Split(m.Text, " ")
		if len(mysplit) > 2 {
			switch mysplit[1] {
			case "register":
				s.handleRegister(m, mysplit)
				return
			}
		}
		s.sendErrorarguments(m.Sender, "Wrong command. Use /help user to see usage")

	}
}

func (s *SAKBot) handleRegister(m *tb.Message, mysplit []string) {
	if len(mysplit) != 4 {
		s.sendWrongarguments(m.Sender, "/user register emailadress password")
		return
	}
	uname := m.Sender.Username
	email := mysplit[2]
	if !validateEmail(email) {
		s.sendErrorarguments(m.Sender, "Email Address not Vaild")
		return
	}
	eexists, err := postgresmodel.Users(postgresmodel.UserWhere.Email.EQ(email)).Exists(context.Background(), s.db.DB)
	if err != nil {
		log.Printf("Error querying DB for Email: %v Error: %v", email, err)
		s.sendErrorarguments(m.Sender, "Error creating User, contact Bot Admin")
		return
	}
	if eexists {
		s.sendErrorarguments(m.Sender, "Email address already registered")
		return
	}

	uexists, err := s.doesUserExist(uname)
	if err != nil {
		log.Printf("Error querying DB for Username: %v Error: %v", uname, err)
		s.sendErrorarguments(m.Sender, "Error creating User, contact Bot Admin")
		return
	}
	if uexists {
		s.sendErrorarguments(m.Sender, "User with this Username already exists")
		return
	}

	hash, err := hashing.GenerateFromPassword(mysplit[3])
	if err != nil {
		log.Printf("Error hashing Passwort for User: %v Error: %v", uname, err)
		s.sendErrorarguments(m.Sender, "Error creating User, contact Bot Admin")
		return
	}
	u := &postgresmodel.User{Username: uname, Email: email, Passwort: hash}
	err = u.Insert(context.Background(), s.db.DB, boil.Infer())
	if err != nil {
		log.Printf("Error Inserting User: %v Error: %v", uname, err)
		s.sendErrorarguments(m.Sender, "Error creating User, contact Bot Admin")
		return
	}
	fmt.Println(u)
	id := uuid.New()
	a := &postgresmodel.OnetimePad{Time: time.Now().Add(2 * 24 * time.Hour), Key: id.String(), UserID: u.UserID, Type: db.PadTypeActivation}
	err = a.Insert(context.Background(), s.db.DB, boil.Infer())
	if err != nil {
		log.Printf("Error Activation for User: %v Error: %v", uname, err)
		s.sendErrorarguments(m.Sender, "Error creating User, contact Bot Admin")
		return
	}
	err = mailer.SendCode(u.Email, u.Username, a.Key)
	if err != nil {
		log.Printf("Error sending Mail for User: %v Error: %v", uname, err)
		s.sendErrorarguments(m.Sender, "Error sending Activation Mail, contact Bot Admin")
		return
	}
	s.sendMessagetoUser(m.Sender, "User created and activation Mail sent")
}
