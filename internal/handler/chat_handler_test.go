package handler

import (
	"bytes"
	"encoding/json"
	"hitalentTest/internal/db"
	"hitalentTest/internal/models"
	"hitalentTest/internal/repository"
	"hitalentTest/internal/service"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateChat(t *testing.T) {
	database := db.ConnectTest()

	chatRepo := repository.NewChatRepository(database)
	chatService := service.NewChatService(chatRepo)

	chatHandler := NewChatHandler(chatService, nil)

	body := []byte(`{"title":"Test chat"}`)
	req := httptest.NewRequest(http.MethodPost, "/chats/", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	chatHandler.CreateChat(rr, req)

	require.Equal(t, http.StatusCreated, rr.Code)

	var chat models.Chat
	err := json.Unmarshal(rr.Body.Bytes(), &chat)
	require.NoError(t, err)

	require.Equal(t, "Test chat", chat.Title)

	require.NotZero(t, chat.ID)
}
