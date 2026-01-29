package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"hitalentTest/internal/models"
	"hitalentTest/internal/service"
)

type ChatHandler struct {
	chatService *service.ChatService
	msgService  *service.MessageService
}

func NewChatHandler(
	chatService *service.ChatService,
	msgService *service.MessageService,
) *ChatHandler {
	return &ChatHandler{
		chatService: chatService,
		msgService:  msgService,
	}
}

// POST /chats/
func (h *ChatHandler) CreateChat(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Title string `json:"title"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}

	chat, err := h.chatService.CreateChat(body.Title)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(chat); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// GET /chats/{id}?limit=N
func (h *ChatHandler) GetChat(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/chats/")
	path = strings.TrimSuffix(path, "/")
	idStr := strings.Split(path, "/")[0]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid chat id", http.StatusBadRequest)
		return
	}

	limit := 20
	if q := r.URL.Query().Get("limit"); q != "" {
		if l, err := strconv.Atoi(q); err == nil && l > 0 && l <= 100 {
			limit = l
		}
	}

	chat, err := h.chatService.GetChat(uint(id))
	if err != nil {
		http.Error(w, "chat not found", http.StatusNotFound)
		return
	}

	messages, err := h.msgService.GetLastMessages(uint(id), limit)
	if err != nil {
		http.Error(w, "failed to get messages", http.StatusInternalServerError)
		return
	}

	resp := struct {
		Chat     *models.Chat     `json:"chat"`
		Messages []models.Message `json:"messages"`
	}{
		Chat:     chat,
		Messages: messages,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

// DELETE /chats/{id}
func (h *ChatHandler) DeleteChat(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/chats/")
	path = strings.TrimSuffix(path, "/")
	idStr := strings.Split(path, "/")[0]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid chat id", http.StatusBadRequest)
		return
	}

	if err := h.chatService.DeleteChat(uint(id)); err != nil {
		http.Error(w, "chat not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
