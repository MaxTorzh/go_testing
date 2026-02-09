package storage

import (
	"errors"
	"github.com/MaxTorzh/go-practice/task5/internal/models"
	"sync"
)

type UserStorage struct {
	mu     sync.RWMutex
	users  map[int]models.User
	nextID int
}

func NewUserStorage() *UserStorage {
	return &UserStorage{
		users: map[int]models.User{
			1: {ID: 1, Name: "Max", Age: 35},
			2: {ID: 2, Name: "Mats", Age: 29},
		},
		nextID: 3,
	}
}

func (s *UserStorage) GetUser(id int) (models.User, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	user, exists := s.users[id]
	if !exists {
		return models.User{}, errors.New("user not found")
	}
	return user, nil
}

func (s *UserStorage) CreateUser(name string, age int) (models.User, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if name == "" {
		return models.User{}, errors.New("name cannot be empty")
	}

	if age < 0 {
		return models.User{}, errors.New("age must be positive")
	}

	user := models.User{
		ID:   s.nextID,
		Name: name,
		Age:  age,
	}
	s.users[user.ID] = user
	s.nextID++
	return user, nil
}
