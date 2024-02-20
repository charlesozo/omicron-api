package main

import (
	"errors"
	"fmt"
	emailverifier "github.com/AfterShip/email-verifier"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/store/sqlstore"
	waLog "go.mau.fi/whatsmeow/util/log"
	"regexp"
	"unicode"
)

var (
	verifier = emailverifier.NewVerifier()
)

func isValidEmail(email string) error {

	verifier = verifier.EnableDomainSuggest()
	verifier = verifier.AddDisposableDomains([]string{"tractorjj.com"})
	ret, err := verifier.Verify(email)
	if err != nil {
		return err
	}
	if !ret.Syntax.Valid {
		return errors.New("email address syntax is invalid")
	}
	if ret.Disposable {
		return errors.New("sorry, we do not accept disposable email addresses")
	}
	if ret.Suggestion != "" {
		return fmt.Errorf("email address is not reachable, looking for %v instead", ret.Suggestion)
	}
	if ret.Reachable == "no" {
		return errors.New("email address is unreachable")
	}
	if !ret.HasMxRecords {
		return errors.New("domain entered not properly setup to recieve emails, MX record not found")
	}
	return nil
}

func isStrongPassword(password string) error {
	// Check if password length is at least 7 characters
	if len(password) < 7 {
		return errors.New("password length is too short, should be above 7")
	}

	// Regular expressions to match uppercase, lowercase, and symbol characters
	hasUppercase := regexp.MustCompile(`[A-Z]`).MatchString
	hasLowercase := regexp.MustCompile(`[a-z]`).MatchString
	hasSymbol := regexp.MustCompile(`[!@#$%^&*()\-_=+{};:,<.>/'"\\|[\]~]`).MatchString

	// Check if password contains at least one uppercase, one lowercase, and one symbol character
	if !hasUppercase(password) || !hasLowercase(password) || !hasSymbol(password) {
		return errors.New("password must contain at least one uppercase, lowercase or symbol")
	}

	return nil
}
func containsNumbersOrSymbols(username string) bool {
	for _, char := range username {
		if !unicode.IsLetter(char) {
			return true
		}
	}
	return false
}

func isWhatsappNumbers(number, dburl string) bool {
	waNumbers := []string{number}
	dbLog := waLog.Stdout("Database", "DEBUG", true)
	clientLog := waLog.Stdout("Client", "DEBUG", true)
	container, err := sqlstore.New("postgres", dburl, dbLog)
	if err != nil {
		fmt.Printf("Unable to create a database store %v", err)
		return false
	}
	deviceStore, err := container.GetFirstDevice()
	if err != nil {
		fmt.Printf("Unable to create a device store %v", err)
		return false
	}
	client := whatsmeow.NewClient(deviceStore, clientLog)
	_, err = client.IsOnWhatsApp(waNumbers)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
