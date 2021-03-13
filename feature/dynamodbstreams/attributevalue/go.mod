module github.com/aws/aws-sdk-go-v2/feature/dynamodbstreams/attributevalue

go 1.15

require (
	github.com/aws/aws-sdk-go-v2 v1.2.1
	github.com/aws/aws-sdk-go-v2/service/dynamodb v1.1.2
	github.com/aws/aws-sdk-go-v2/service/dynamodbstreams v1.1.2
	github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding v1.0.2
	github.com/google/go-cmp v0.5.4
)

replace github.com/aws/aws-sdk-go-v2/service/dynamodb => ../../../service/dynamodb/

replace github.com/aws/aws-sdk-go-v2 => ../../../

replace github.com/aws/aws-sdk-go-v2/service/dynamodbstreams => ../../../service/dynamodbstreams/

replace github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding => ../../../service/internal/accept-encoding/
