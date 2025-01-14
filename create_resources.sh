
#!/bin/bash

echo $PWD

awslocal dynamodb create-table --cli-input-json file:///tmp/dynamo/dynamoTable.json --region=sa-east-1
    
awslocal --endpoint-url=http://localhost:4566 --region=sa-east-1  dynamodb list-tables

awslocal --endpoint-url=http://localhost:4566 --region=sa-east-1 dynamodb put-item --table-name TestTable --item "{\"Id\":{\"S\":\"123456789abc\"}}"

awslocal --endpoint-url=http://localhost:4566 --region=sa-east-1 dynamodb scan --table-name TestTable

awslocal s3api create-bucket --endpoint-url=http://localhost:4566 --bucket test-bucket --region sa-east-1 --create-bucket-configuration LocationConstraint=sa-east-1
