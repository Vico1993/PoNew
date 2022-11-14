package notification

import "net/http"

type Notification interface {
	BodyKeys() []string
	GetType() string
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
