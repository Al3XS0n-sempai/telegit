package bot

import (
	"fmt"

	tele "gopkg.in/telebot.v3"
)

// TestHandler is handler for testing Telebot initalizing
func TestHandler(c tele.Context) error {
	fmt.Println("Recieved message with args ", c.Args())
	fmt.Println(c.Sender().ID)
	fmt.Println(c.Sender().FirstName)
	fmt.Println(c.Sender().LastName)
	fmt.Println(c.Sender().Username)
	fmt.Println(c.Sender().Usernames)
	fmt.Println(c.Sender().IsBot)
	fmt.Println(c.Sender().IsPremium)
	fmt.Println(c.Sender().LanguageCode)
	fmt.Println(c.Sender().SupportsInline)
	fmt.Println(c.Sender().Recipient())
	return c.Send("Test message response")
}
