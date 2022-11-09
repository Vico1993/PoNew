package notification

import (
	"fmt"
	"net/http"
)

// Looking for:
// x-ponew-type <- telegram
// x-ponew-telegram-chatId <- Chat ID where to send the notification
// x-ponew-telegram-token <- bot token

type Telegram struct {
	baseNotification
}

// List of the keys required by the Notification
func (t *Telegram) HeaderKeys() []string {
	return []string{
		"x-ponew-telegram-chatId",
		"x-ponew-telegram-token",
	}
}

// Return the type of the Notification struct
func (t *Telegram) GetType() string {
	return t.baseNotification.notifType
}

// Valid the header keys needed for the Telegram notification
func (t *Telegram) ValidHeader(req *http.Request) bool {
	for _, key := range t.HeaderKeys() {
		if req.Header.Get(key) == "" {
			return false
		}
	}

	return true
}

// Processing the notification
func (t *Telegram) Process(req *http.Request, res http.ResponseWriter) {
	fmt.Println("Processing ...")

	fmt.Println(req.FormValue("type"))
}
