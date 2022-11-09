package notification

import "net/http"

type Notification interface {
	HeaderKeys() []string
	GetType() string
	ValidHeader(req *http.Request) bool
	Process(req *http.Request, res http.ResponseWriter)
}

type baseNotification struct {
	notifType string
}

var NotificationList []Notification

// Will load all notification available
func InitNotification() {
	NotificationList = []Notification{
		&Telegram{
			baseNotification{
				notifType: "telegram",
			},
		},
	}
}
