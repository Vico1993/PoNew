package handler

import (
	"net/http"

	"github.com/Vico1993/Ponew/domain/notification"
	myUtil "github.com/Vico1993/Ponew/domain/util"
	"github.com/gorilla/mux"
)

func PostNotification(res http.ResponseWriter, req *http.Request) {
	notifType := mux.Vars(req)["type"]

	if notifType == "" {
		// Shoult not happend
		myUtil.SendError(res, myUtil.NewValidationError("Missing type in the path"))
		return
	}

	for _, notif := range notification.NotificationList {
		if notif.GetType() == notifType {
			notif.Process(req, res)
			break
		}
	}

	// Default
	myUtil.SendError(res, myUtil.NewValidationError("Couldn't process your request, type not recognized"))
}
