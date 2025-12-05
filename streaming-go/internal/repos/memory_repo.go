package repos

import (
	"sync"

	customErr "streaming/internal/errors"
	"streaming/internal/models"
)

// Interfaces esperadas por el servicio
type UserRepo interface {
	CreateUser(u *models.User) error
	GetUserByID(id string) (*models.User, error)
	ListUsers() []*models.User
}

type ContentRepo interface {
	CreateContent(c *models.Content) error
	GetContentByID(id string) (*models.Content, error)
	ListContent() []*models.Content
}

// Repositorio único en memoria
type MemoryRepo struct {
	users    map[string]*models.User
	contents map[string]*models.Content
	mu       sync.RWMutex
}

func NewMemoryRepo() *MemoryRepo {
	return &MemoryRepo{
		users:    make(map[string]*models.User),
		contents: make(map[string]*models.Content),
	}
}

// ------------------ USUARIOS ------------------

func (r *MemoryRepo) CreateUser(u *models.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, exists := r.users[u.GetID()]; exists {
		return customErr.ErrAlreadyExists{Entity: "User", ID: u.GetID()}
	}
	r.users[u.GetID()] = u
	return nil
}

func (r *MemoryRepo) GetUserByID(id string) (*models.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	if u, ok := r.users[id]; ok {
		return u, nil
	}
	return nil, customErr.ErrNotFound{Entity: "User", ID: id}
}

func (r *MemoryRepo) ListUsers() []*models.User {
	r.mu.RLock()
	defer r.mu.RUnlock()
	out := make([]*models.User, 0, len(r.users))
	for _, user := range r.users {
		out = append(out, user)
	}
	return out
}

// ------------------ CONTENIDO ------------------

func (r *MemoryRepo) CreateContent(c *models.Content) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, exists := r.contents[c.GetID()]; exists {
		return customErr.ErrAlreadyExists{Entity: "Content", ID: c.GetID()}
	}
	r.contents[c.GetID()] = c
	return nil
}

func (r *MemoryRepo) GetContentByID(id string) (*models.Content, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	if c, ok := r.contents[id]; ok {
		return c, nil
	}
	return nil, customErr.ErrNotFound{Entity: "Content", ID: id}
}

func (r *MemoryRepo) ListContent() []*models.Content {
	r.mu.RLock()
	defer r.mu.RUnlock()
	out := make([]*models.Content, 0, len(r.contents))
	for _, c := range r.contents {
		out = append(out, c)
	}
	return out
}
