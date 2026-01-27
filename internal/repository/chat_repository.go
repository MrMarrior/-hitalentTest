package repository

import (
	"hitalentTest/internal/models"

	"gorm.io/gorm"
)

type ChatRepository struct {
	db *gorm.DB
}

func NewChatRepository(db *gorm.DB) *ChatRepository {
	return &ChatRepository{db: db}
}

func (r *ChatRepository) Create(chat *models.Chat) (*models.Chat, error) {
	if err := r.db.Create(chat).Error; err != nil {
		return nil, err
	}
	return chat, nil
}

func (r *ChatRepository) GetByID(id uint) (*models.Chat, error) {
	var chat models.Chat
	if err := r.db.Preload("Messages").First(&chat, id).Error; err != nil {
		return nil, err
	}
	return &chat, nil
}

func (r *ChatRepository) Delete(id uint) error {
	return r.db.Delete(&models.Chat{}, id).Error
}
