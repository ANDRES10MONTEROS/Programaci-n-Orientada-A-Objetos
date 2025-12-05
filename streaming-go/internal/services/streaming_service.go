package services

import (
	"fmt"
	"time"

	customErr "streaming/internal/errors"
	"streaming/internal/models"
)

// ===== Interfaces que debe implementar el repo =====

type UserRepo interface {
	CreateUser(u *models.User) error
	GetUserByID(id string) (*models.User, error)
}

type ContentRepo interface {
	CreateContent(c *models.Content) error
	GetContentByID(id string) (*models.Content, error)
	ListContent() []*models.Content
}

// ===== Servicio principal =====

type StreamingService struct {
	userRepo    UserRepo
	contentRepo ContentRepo
	playHistory map[string][]string
}

func NewStreamingService(u UserRepo, c ContentRepo) *StreamingService {
	return &StreamingService{
		userRepo:    u,
		contentRepo: c,
		playHistory: make(map[string][]string),
	}
}

// ===== Crear usuario =====

func (s *StreamingService) RegisterUser(id, name, email string) (*models.User, error) {
	user, err := models.NewUser(id, name, email)
	if err != nil {
		return nil, err
	}

	err = s.userRepo.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// ===== Crear contenido =====

func (s *StreamingService) AddContent(id, title string, ctype models.ContentType, duration int) (*models.Content, error) {
	content, err := models.NewContent(id, title, ctype, duration)
	if err != nil {
		return nil, err
	}

	err = s.contentRepo.CreateContent(content)
	if err != nil {
		return nil, err
	}

	return content, nil
}

// ===== Reproducir contenido =====

func (s *StreamingService) Play(userID, contentID string) (string, error) {
	_, err := s.userRepo.GetUserByID(userID)
	if err != nil {
		return "", customErr.ErrNotFound{Entity: "User", ID: userID}
	}

	content, err := s.contentRepo.GetContentByID(contentID)
	if err != nil {
		return "", customErr.ErrNotFound{Entity: "Content", ID: contentID}
	}

	s.playHistory[userID] = append(s.playHistory[userID], contentID)

	msg := fmt.Sprintf(
		"User %s está reproduciendo: %s (%ds) - %s",
		userID,
		content.GetTitle(),
		content.GetDuration(),
		time.Now().Format("2006-01-02 15:04:05"),
	)

	return msg, nil
}

// ===== Historial =====

func (s *StreamingService) GetHistory(userID string) ([]*models.Content, error) {
	ids, exists := s.playHistory[userID]
	if !exists {
		return []*models.Content{}, nil
	}

	var result []*models.Content

	for _, id := range ids {
		if content, err := s.contentRepo.GetContentByID(id); err == nil {
			result = append(result, content)
		}
	}

	return result, nil
}
