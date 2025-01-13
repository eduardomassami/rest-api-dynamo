package application

import (
	"encoding/json"
	"fmt"

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
	s.bucket.Save(userJSON)

	return s.repo.Save(user)
}

func (s *UserService) GetUser(id string) (*domain.User, error) {
	return s.repo.FindByID(id)
}
