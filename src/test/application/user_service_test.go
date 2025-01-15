package application_test

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/eduardomassami/rest-api-dynamo/application"
	"github.com/eduardomassami/rest-api-dynamo/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock para o repositório
type MockRepo struct {
	mock.Mock
}

func (m *MockRepo) Save(user domain.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockRepo) FindByID(id string) (*domain.User, error) {
	args := m.Called(id)
	user, _ := args.Get(0).(*domain.User) // Convertendo o retorno para o tipo *domain.User
	return user, args.Error(1)
}

// Mock para o bucket
type MockBucket struct {
	mock.Mock
}

func (m *MockBucket) Save(data []byte) error {
	args := m.Called(data)
	return args.Error(0)
}

func Test_CreateUser(t *testing.T) {
	// Criar mocks
	mockRepo := new(MockRepo)
	mockBucket := new(MockBucket)

	// Criar o serviço usando os mocks
	service := application.NewUserService(mockRepo, mockBucket)

	// Dados do teste
	user := domain.User{Id: "123", Name: "Test", Email: "test@example.com"}
	userJSON, _ := json.Marshal(user)

	// Configurar os mocks
	mockRepo.On("Save", user).Return(nil)
	mockBucket.On("Save", userJSON).Return(nil)

	// Chamar o método
	err := service.CreateUser(user)

	// Verificar o resultado
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
	mockBucket.AssertExpectations(t)
}

func TestCreateUser_SaveRepoFails(t *testing.T) {
	// Criar mocks
	mockRepo := new(MockRepo)
	mockBucket := new(MockBucket)

	// Criar o serviço usando os mocks
	service := application.NewUserService(mockRepo, mockBucket)

	// Dados do teste
	user := domain.User{Id: "123", Name: "Test", Email: "test@example.com"}
	userJSON, _ := json.Marshal(user)

	// Configurar os mocks
	mockRepo.On("Save", user).Return(errors.New("repository save error"))
	mockBucket.On("Save", userJSON).Return(nil)

	// Chamar o método
	err := service.CreateUser(user)

	// Verificar o resultado
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "repository save error")
	mockRepo.AssertExpectations(t)
	mockBucket.AssertExpectations(t)
}
