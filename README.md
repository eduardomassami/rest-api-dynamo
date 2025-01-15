# rest-api-dynamo
Simple CRUD DynamoDB API 


go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
go tool cover -func=coverage.out