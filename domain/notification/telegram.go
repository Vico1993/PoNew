package notification

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	myUtil "github.com/Vico1993/Ponew/domain/util"
)

// Looking for:
// type <- telegram
// telegram-chatId <- Chat ID where to send the notification
// telegram-token <- bot token

type Telegram struct {
	baseNotification
}

type TelegramBody struct {
	TelegramChatId string `json:"telegram-chatId"`
	TelegramToken  string `json:"telegram-token"`
}

// List of the keys required by the Notification
func (t *Telegram) BodyKeys() []string {
	return []string{
		"telegram-chatId",
		"telegram-token",
	}
}

// Return the type of the Notification struct
func (t *Telegram) GetType() string {
	return t.baseNotification.notifType
}

// Processing the notification
func (t *Telegram) Process(req *http.Request, res http.ResponseWriter) {
	var body TelegramBody
	err := json.NewDecoder(req.Body).Decode(&body)

	// Valid the body
	if err != nil || body.TelegramChatId == "" || body.TelegramToken == "" {
		myUtil.SendError(res, myUtil.NewValidationError("Keys incomplete or incorrect for Telegram Notification, make sure body contains: "+strings.Join(t.BodyKeys(), ", ")))
		return
	}

	fmt.Println("Valide BODY de ouf!", body, body.TelegramChatId, body.TelegramToken)

}
