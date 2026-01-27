package service

import (
	"errors"
	"strings"

	"hitalentTest/internal/models"
	"hitalentTest/internal/repository"
)

type ChatService struct {
	repo *repository.ChatRepository
}

func NewChatService(repo *repository.ChatRepository) *ChatService {
	return &ChatService{repo: repo}
}

func (s *ChatService) CreateChat(title string) (*models.Chat, error) {
	title = strings.TrimSpace(title)
	if len(title) == 0 || len(title) > 200 {
		return nil, errors.New("title must be 1-200 characters")
	}
	chat := &models.Chat{Title: title}
	return s.repo.Create(chat)
}

func (s *ChatService) GetChat(id uint) (*models.Chat, error) {
	return s.repo.GetByID(id)
}

func (s *ChatService) DeleteChat(id uint) error {
	return s.repo.Delete(id)
}
