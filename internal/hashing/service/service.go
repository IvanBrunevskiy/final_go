package service

import (
	"crypto/sha256"
	"fmt"
)

type HashingService struct {
	store map[string]string
}

func NewHashingService() *HashingService {
	return &HashingService{
		store: make(map[string]string),
	}
}

// CheckHash проверяет, существует ли хеш для заданного payload
func (s *HashingService) CheckHash(payload string) (bool, error) {
	_, exists := s.store[payload]
	return exists, nil
}

// GetHash возвращает хеш для заданного payload, если он существует
func (s *HashingService) GetHash(payload string) (string, error) {
	hash, exists := s.store[payload]
	if !exists {
		return "", fmt.Errorf("hash not found for payload: %s", payload)
	}
	return hash, nil
}

// CreateHash создает новый хеш для заданного payload и сохраняет его
func (s *HashingService) CreateHash(payload string) (string, error) {
	if _, exists := s.store[payload]; exists {
		return "", fmt.Errorf("hash already exists for payload: %s", payload)
	}

	// Пример простого хеширования
	hash := fmt.Sprintf("%x", sha256.Sum256([]byte(payload)))
	s.store[payload] = hash

	return hash, nil
}
