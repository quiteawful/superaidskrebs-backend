package superaidskrebsbot

import (
	"context"
	"crypto/rand"
	"errors"
	"fmt"
	"log"
	"math/big"
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

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			log.Fatal("Crypto Rand Failed:", err)
		}
		b[i] = letters[n.Int64()]
	}
	return string(b)
}

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

func (s *SAKBot) getUserbyName(username string) (*postgresmodel.User, error) {
	u, err := postgresmodel.Users(postgresmodel.UserWhere.Username.EQ(username)).One(context.Background(), s.db.DB)
	if err != nil {
		return nil, errors.New("Error querying for User: " + err.Error())
	}
	return u, nil
}

func (s *SAKBot) userPreFlightCheck(username string, checkloggedin bool) (*postgresmodel.User, bool, error) {
	uexists, err := s.doesUserExist(username)
	auth := false
	if err != nil {
		return nil, auth, errors.New("Error checking for existence of user" + err.Error())
	}
	if !uexists {
		return nil, auth, errors.New("User does not exist" + err.Error())
	}
	u, err := s.getUserbyName(username)
	if err != nil {
		return nil, auth, errors.New("Error querying for User: " + err.Error())
	}
	if !u.Active {
		return nil, auth, errors.New("User not activated" + err.Error())
	}

	if checkloggedin {
		auth, err = u.UserIDOnetimePads(postgresmodel.OnetimePadWhere.Type.EQ(db.PadTypeAuth)).Exists(context.Background(), s.db.DB)
		if err != nil {
			return nil, auth, errors.New("Error querying for User auth: " + err.Error())
		}
	}
	return u, auth, nil
}

func (s *SAKBot) handleUserCommands(m *tb.Message) {
	if !m.FromChannel() {
		mysplit := strings.Split(m.Text, " ")
		if len(mysplit) > 1 {
			switch mysplit[1] {
			case "register":
				s.handleRegister(m, mysplit)
				return
			case "activate":
				s.handleActivate(m, mysplit)
				return
			case "login":
				s.handleLogin(m, mysplit)
				return
			case "resetpw":
				s.handleResetPassword(m)
				return
			case "changepw":
				s.handleChangePassword(m, mysplit)
				return
			case "changeemail":
				s.handleChangeEmail(m, mysplit)
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

func (s *SAKBot) handleActivate(m *tb.Message, mysplit []string) {
	if len(mysplit) != 3 {
		s.sendWrongarguments(m.Sender, "/user activate code")
		return
	}
	code := mysplit[2]
	uname := m.Sender.Username
	o, err := postgresmodel.OnetimePads(postgresmodel.OnetimePadWhere.Key.EQ(code)).One(context.Background(), s.db.DB)
	if err != nil {
		log.Printf("Error querying for activation code: %v Error: %v", uname, err)
		s.sendErrorarguments(m.Sender, "Error activating User, contact Bot Admin")
		return
	}
	if o.Type != db.PadTypeActivation {
		log.Printf("Error code is not a activation code: %v User: %v", code, uname)
		s.sendErrorarguments(m.Sender, "Error activating User, contact Bot Admin")
		return
	}
	u, err := postgresmodel.FindUser(context.Background(), s.db.DB, o.UserID)
	if err != nil {
		log.Printf("Error querying for activation User: %v Error: %v", uname, err)
		s.sendErrorarguments(m.Sender, "Error activating User, contact Bot Admin")
		return
	}
	if uname != u.Username {
		log.Printf("Error activtion User and username differen: %v vs %v", u.Username, uname)
		s.sendErrorarguments(m.Sender, "Error activating User, contact Bot Admin")
		return
	}
	u.Active = true
	_, err = u.Update(context.Background(), s.db.DB, boil.Whitelist(postgresmodel.UserColumns.Active))
	if err != nil {
		log.Printf("Error setting User active: %v Error: %v", uname, err)
		s.sendErrorarguments(m.Sender, "Error activating User, contact Bot Admin")
		return
	}
	_, err = o.Delete(context.Background(), s.db.DB)
	if err != nil {
		log.Printf("Error deleting activation pad: Error: %v", err)
	}
	s.sendMessagetoUser(m.Sender, "User account activated")
}

func (s *SAKBot) handleLogin(m *tb.Message, mysplit []string) {
	if len(mysplit) != 3 {
		s.sendWrongarguments(m.Sender, "/user login password")
		return
	}
	pw := mysplit[2]
	uname := m.Sender.Username
	u, _, err := s.userPreFlightCheck(uname, false)
	if err != nil {
		log.Printf("Error on User Preflight for User: %v Error: %v", uname, err)
		s.sendErrorarguments(m.Sender, "Error logging User in, contact Bot Admin")
		return
	}
	match, err := hashing.ComparePasswordAndHash(pw, u.Passwort)
	if err != nil {
		log.Printf("Error matching Password for User: %v Error: %v", uname, err)
		s.sendErrorarguments(m.Sender, "Error logging User in, contact Bot Admin")
		return
	}
	if !match {
		log.Printf("Error Passwords not matching for User: %v", uname)
		s.sendErrorarguments(m.Sender, "Error logging User in, contact Bot Admin")
		return
	}
	id := uuid.New()
	a := &postgresmodel.OnetimePad{Time: time.Now().Add(1 * time.Minute), Key: id.String(), UserID: u.UserID, Type: db.PadTypeAuth}
	err = a.Insert(context.Background(), s.db.DB, boil.Infer())
	if err != nil {
		log.Printf("Error creating authpad for User: %v Error: %v", uname, err)
		s.sendErrorarguments(m.Sender, "Error logging User in, contact Bot Admin")
		return
	}
	s.sendMessagetoUser(m.Sender, "User logged in")
}

func (s *SAKBot) setNewUserPassword(u *postgresmodel.User, newpw string) error {
	hash, err := hashing.GenerateFromPassword(newpw)
	if err != nil {
		return fmt.Errorf("Error creating new pw hash for User: %v Error: %v", u.Username, err)
	}
	u.Passwort = hash
	_, err = u.Update(context.Background(), s.db.DB, boil.Whitelist(postgresmodel.UserColumns.Passwort))
	if err != nil {
		return fmt.Errorf("Error setting User password: %v Error: %v", u.Username, err)
	}
	return nil
}

func (s *SAKBot) handleResetPassword(m *tb.Message) {
	uname := m.Sender.Username
	u, _, err := s.userPreFlightCheck(uname, false)
	if err != nil {
		log.Printf("Error on User Preflight for User: %v Error: %v", uname, err)
		s.sendErrorarguments(m.Sender, "Error resetting password contact Bot Admin")
		return
	}
	newpw := randSeq(20)
	err = s.setNewUserPassword(u, newpw)
	if err != nil {
		log.Printf("Error setting new Password for User: %v Error: %v", uname, err)
		s.sendErrorarguments(m.Sender, "Error resetting password, contact Bot Admin")
		return
	}
	err = mailer.SendPassword(u.Email, u.Username, newpw)
	if err != nil {
		log.Printf("Error sending Mail for User: %v Error: %v", uname, err)
		s.sendErrorarguments(m.Sender, "Error resetting password, contact Bot Admin")
		return
	}
	s.sendMessagetoUser(m.Sender, "Password resetted, check your Mail")
}

func (s *SAKBot) handleChangePassword(m *tb.Message, mysplit []string) {
	if len(mysplit) != 3 {
		s.sendWrongarguments(m.Sender, "/user changepw newpassword")
		return
	}
	newpw := mysplit[2]
	uname := m.Sender.Username
	u, auth, err := s.userPreFlightCheck(uname, true)
	if err != nil {
		log.Printf("Error on User Preflight for User: %v Error: %v", uname, err)
		s.sendErrorarguments(m.Sender, "Error setting password, contact Bot Admin")
		return
	}
	if !auth {
		s.sendErrorarguments(m.Sender, "User not logged in")
		return
	}
	err = s.setNewUserPassword(u, newpw)
	if err != nil {
		log.Printf("Error setting new Password for User: %v Error: %v", uname, err)
		s.sendErrorarguments(m.Sender, "Error setting password, contact Bot Admin")
		return
	}
	s.sendMessagetoUser(m.Sender, "Password set")
}

func (s *SAKBot) handleChangeEmail(m *tb.Message, mysplit []string) {
	if len(mysplit) != 3 {
		s.sendWrongarguments(m.Sender, "/user changeemail newemail")
		return
	}
	newmail := mysplit[2]
	uname := m.Sender.Username
	u, auth, err := s.userPreFlightCheck(uname, true)
	if err != nil {
		log.Printf("Error on User Preflight for User: %v Error: %v", uname, err)
		s.sendErrorarguments(m.Sender, "Error changing email, contact Bot Admin")
		return
	}
	if !auth {
		s.sendErrorarguments(m.Sender, "User not logged in")
		return
	}
	u.Email = newmail
	_, err = u.Update(context.Background(), s.db.DB, boil.Whitelist(postgresmodel.UserColumns.Email))
	if err != nil {
		log.Printf("Error setting new Email for User: %v Error: %v", uname, err)
		s.sendErrorarguments(m.Sender, "Error changing email, contact Bot Admin")
		return
	}
	s.sendMessagetoUser(m.Sender, "Email set")
}
