package service

import (
	"errors"
	"strings"

	"hitalentTest/internal/models"
	"hitalentTest/internal/repository"
)

type MessageService struct {
	repo     *repository.MessageRepository
	chatRepo *repository.ChatRepository
}

func NewMessageService(repo *repository.MessageRepository, chatRepo *repository.ChatRepository) *MessageService {
	return &MessageService{repo: repo, chatRepo: chatRepo}
}

func (s *MessageService) SendMessage(chatID uint, text string) (*models.Message, error) {
	text = strings.TrimSpace(text)
	if len(text) == 0 || len(text) > 5000 {
		return nil, errors.New("text must be 1-5000 characters")
	}

	if _, err := s.chatRepo.GetByID(chatID); err != nil {
		return nil, errors.New("chat not found")
	}

	msg := &models.Message{
		ChatID: chatID,
		Text:   text,
	}
	return s.repo.Create(msg)
}

func (s *MessageService) GetLastMessages(chatID uint, limit int) ([]models.Message, error) {
	if limit <= 0 || limit > 100 {
		limit = 20
	}
	return s.repo.GetLastMessages(chatID, limit)
}
