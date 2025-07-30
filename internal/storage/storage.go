package storage

import (
	"gorm.io/gorm"

	"url-shortener/internal/models"
)

type LinkStorage struct{ db *gorm.DB }

func NewLinkStorage(db *gorm.DB) *LinkStorage { return &LinkStorage{db: db} }

func (s *LinkStorage) AddLink(link models.Link) error {
	result := s.db.Create(&link)
	return result.Error
}

func (s *LinkStorage) DeleteLinkBySlug(slug string) error {
	result := s.db.Delete(&models.Link{}, "short = ?", slug)
	return result.Error
}

func (s *LinkStorage) GetLinkBySlug(slug string) (models.Link, bool) {
	var link models.Link
	result := s.db.First(&link, "short = ?", slug)
	if result.Error != nil {
		return models.Link{}, false
	}
	return link, true
}

func (s *LinkStorage) GetAllLinks() []models.Link {
	var links []models.Link
	s.db.Find(&links)
	return links
}

func (s *LinkStorage) LinksByUser(userID string) []models.Link {
	var links []models.Link
	s.db.Where("user_id = ?", userID).Find(&links)
	return links
}
