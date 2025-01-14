package application

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/eduardomassami/rest-api-dynamo/domain"
)

type UserService struct {
	repo   domain.UserRepository
	bucket domain.UserBucket
}

func NewUserService(repo domain.UserRepository, bucket domain.UserBucket) *UserService {
	return &UserService{repo: repo, bucket: bucket}
}

func (s *UserService) CreateUser(user domain.User) error {

	userJSON, err := json.Marshal(user)
	if err != nil {
		return fmt.Errorf("erro ao serializar usuário: %w", err)
	}

	var wg sync.WaitGroup
	var saveDynamoError error

	wg.Add(2)

	go func() {
		defer wg.Done()
		s.bucket.Save(userJSON)
	}()

	go func() {
		defer wg.Done() // Marcar como concluída
		saveDynamoError = s.repo.Save(user)
	}()

	wg.Wait()

	return saveDynamoError
}

func (s *UserService) GetUser(id string) (*domain.User, error) {
	return s.repo.FindByID(id)
}
