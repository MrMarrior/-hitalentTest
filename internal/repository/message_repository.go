package repository

import (
	"hitalentTest/internal/models"

	"gorm.io/gorm"
)

type MessageRepository struct {
	db *gorm.DB
}

func NewMessageRepository(db *gorm.DB) *MessageRepository {
	return &MessageRepository{db: db}
}

func (r *MessageRepository) Create(msg *models.Message) (*models.Message, error) {
	if err := r.db.Create(msg).Error; err != nil {
		return nil, err
	}
	return msg, nil
}

func (r *MessageRepository) GetLastMessages(chatID uint, limit int) ([]models.Message, error) {
	var messages []models.Message
	if err := r.db.Where("chat_id = ?", chatID).Order("created_at desc").Limit(limit).Find(&messages).Error; err != nil {
		return nil, err
	}
	return messages, nil
}
