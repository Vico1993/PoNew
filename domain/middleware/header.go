package middleware

import (
	"net/http"
	"strings"

	"github.com/Vico1993/Ponew/domain/notification"
	myUtil "github.com/Vico1993/Ponew/domain/util"
)

// Will check if the header match one of the one allowed
func CheckRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		notifType := req.Header.Get("x-ponew-type")

		if notifType == "" {
			myUtil.SendError(res, myUtil.NewValidationError("Missing type key in the header"))
			return
		}

		notificationFound := false
		for _, notif := range notification.NotificationList {
			if notif.GetType() != notifType {
				continue
			}

			if !notif.ValidHeader(req) {
				myUtil.SendError(res, myUtil.NewValidationError("Missing key for: "+notif.GetType()+", you need to provide: "+strings.Join(notif.HeaderKeys(), ", ")))
				return
			}

			notificationFound = true
		}

		if !(notificationFound) {
			myUtil.SendError(res, myUtil.NewValidationError("Couldn't process your request"))
			return
		}

		next.ServeHTTP(res, req)
	})
}
