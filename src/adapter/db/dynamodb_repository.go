package db

import (
	"context"
	"errors"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/eduardomassami/rest-api-dynamo/domain"
)

type DynamoDBRepository struct {
	client    *dynamodb.Client
	tableName string
}

func NewDynamoDBRepository(dbClient *dynamodb.Client, tableName string) *DynamoDBRepository {
	return &DynamoDBRepository{
		client:    dbClient,
		tableName: tableName,
	}
}

func (r *DynamoDBRepository) Save(user domain.User) error {
	_, err := r.client.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: &r.tableName,
		Item: map[string]types.AttributeValue{
			"Id":    &types.AttributeValueMemberS{Value: user.Id},
			"Name":  &types.AttributeValueMemberS{Value: user.Name},
			"Email": &types.AttributeValueMemberS{Value: user.Email},
		},
	})
	return err
}

// FindByID retrieves a user by ID from DynamoDB.
func (r *DynamoDBRepository) FindByID(id string) (*domain.User, error) {
	result, err := r.client.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: &r.tableName,
		Key: map[string]types.AttributeValue{
			"Id": &types.AttributeValueMemberS{Value: id},
		},
	})
	if err != nil {
		return nil, err
	}

	if result.Item == nil {
		return nil, errors.New("user not found")
	}

	user := &domain.User{
		Id:    result.Item["Id"].(*types.AttributeValueMemberS).Value,
		Name:  result.Item["Name"].(*types.AttributeValueMemberS).Value,
		Email: result.Item["Email"].(*types.AttributeValueMemberS).Value,
	}

	return user, nil
}
