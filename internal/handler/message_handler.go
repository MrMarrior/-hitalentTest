package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"hitalentTest/internal/service"
)

type MessageHandler struct {
	msgService *service.MessageService
}

func NewMessageHandler(msgService *service.MessageService) *MessageHandler {
	return &MessageHandler{msgService: msgService}
}

func (h *MessageHandler) SendMessage(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/chats/")
	path = strings.TrimSuffix(path, "/messages/")
	idStr := strings.Split(path, "/")[0]

	chatID, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid chat id", http.StatusBadRequest)
		return
	}

	var body struct {
		Text string `json:"text"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}

	msg, err := h.msgService.SendMessage(uint(chatID), body.Text)
	if err != nil {
		if err.Error() == "chat not found" {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(msg)
}
