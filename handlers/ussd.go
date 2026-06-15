package handlers

import "net/http"

type USSDRequest struct {
	SessionID   string
	ServiceCode string
	PhoneNumber string
	Text        string
}

type USSDResponse struct {
	Action  string
	Message string
}

func HandleUSSD(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	_, _ = w.Write([]byte("USSD handler not yet implemented"))
}
