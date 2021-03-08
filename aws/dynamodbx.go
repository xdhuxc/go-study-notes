package aws

import (
	"errors"
	"math"
	"os"
	"sync"

	"github.com/asaskevich/govalidator"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/ec2metadata"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

const (
	BatchLimit = 25
)

type DynamodbClient struct {
	dynamodb  *dynamodb.DynamoDB
	tableName string
}

// 创建 DynamoDB 客户端
func NewDynamoDBClient(region string, endpoint string, disableSSL bool) (*DynamodbClient, error) {
	if result := govalidator.Trim(region, " "); govalidator.IsNull(result) {
		sess, err := session.NewSession()
		if err != nil {
			return nil, err
		}
		metadata := ec2metadata.New(sess)
		eccRegion, err := metadata.Region()
		if err != nil {
			return nil, errors.New("the AWS service requires a region")
		}

		region = eccRegion
	}
	var c *aws.Config
	// 使用环境变量中的 AccessKey 和 SecretKey
	accessKey := os.Getenv("AWS_ACCESS_KEY_ID")
	secretKey := os.Getenv("AWS_SECRET_ACCESS_KEY")
	if accessKey != "" && secretKey != "" {
		creds := credentials.NewStaticCredentials(accessKey, secretKey, "")
		c = &aws.Config{
			Region:      aws.String(region),
			Credentials: creds,
			Endpoint:    aws.String(endpoint),
			DisableSSL:  aws.Bool(disableSSL),
		}
	} else if accessKey != "" && secretKey != "" { // 使用配置文件中的 AccessKey 和 SecretKey
		creds := credentials.NewStaticCredentials(accessKey, secretKey, "")
		c = &aws.Config{
			Region:      aws.String(region),
			Credentials: creds,
			Endpoint:    aws.String(endpoint),
			DisableSSL:  aws.Bool(disableSSL),
		}
	} else {
		c = &aws.Config{ // 使用 ~/.aws/credentials 文件中的 AccessKey 和 SecretKey 。
			Region:     aws.String(region),
			Endpoint:   aws.String(endpoint),
			DisableSSL: aws.Bool(disableSSL),
		}
	}

	sess, err := session.NewSession(c)
	if err != nil {
		return nil, err
	}

	db := dynamodb.New(sess)

	return &DynamodbClient{
		dynamodb: db,
	}, nil
}

func (dc *DynamodbClient) SetTableName(tableName string) {
	dc.tableName = tableName
}

// 创建 dynamodb 表
func (dc *DynamodbClient) CreateTable() error {
	itemInput := &dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("key"),
				AttributeType: aws.String("S"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("key"),
				KeyType:       aws.String("HASH"),
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(10),
			WriteCapacityUnits: aws.Int64(10),
		},
		TableName: aws.String(dc.tableName),
	}
	_, err := dc.dynamodb.CreateTable(itemInput)
	if err != nil {
		return err
	}

	return nil
}

// 删除 dynamodb 表
func (dc *DynamodbClient) DeleteTable() error {
	itemInput := &dynamodb.DeleteTableInput{
		TableName: aws.String(dc.tableName),
	}

	if _, err := dc.dynamodb.DeleteTable(itemInput); err != nil {
		return err
	}
	return nil
}

// 向指定的 dynamodb 表中插入单个键值对数据
func (dc *DynamodbClient) Insert(key string, value string) error {
	itemInput := &dynamodb.PutItemInput{
		TableName: aws.String(dc.tableName),
		Item: map[string]*dynamodb.AttributeValue{
			"key": {
				S: aws.String(key),
			},
			"value": {
				S: aws.String(value),
			},
		},
	}

	if _, err := dc.dynamodb.PutItem(itemInput); err != nil {
		return err
	}
	return nil
}

// 批量删除指定的 dynamodb 中的数据
func (dc *DynamodbClient) BatchDeleteItems() error {
	errc := make(chan error, 1)
	wg := sync.WaitGroup{}

	scannedData, err := dc.dynamodb.Scan(&dynamodb.ScanInput{TableName: aws.String(dc.tableName)})
	if err != nil {
		return err
	}

	length := len(scannedData.Items)
	wg.Add(int(math.Ceil((float64(length)) / float64(BatchLimit))))

	req := []*dynamodb.WriteRequest{}
	for i, a := range scannedData.Items {
		req = append(req, &dynamodb.WriteRequest{
			DeleteRequest: &dynamodb.DeleteRequest{
				Key: map[string]*dynamodb.AttributeValue{
					"key": a["key"],
				},
			},
		})
		if (i+1)%BatchLimit == 0 || i >= int(*scannedData.Count)-1 {
			go func(reqChunk []*dynamodb.WriteRequest) {
				defer wg.Done()
				_, err := dc.dynamodb.BatchWriteItem(&dynamodb.BatchWriteItemInput{
					RequestItems: map[string][]*dynamodb.WriteRequest{
						dc.tableName: reqChunk,
					},
				})
				if err != nil {
					errc <- err
				}
			}(req)
			req = []*dynamodb.WriteRequest{}
		}
	}
	go func() {
		wg.Wait()
		close(errc)
	}()
	return <-errc
}

// 向指定的 dynamodb 表中批量插入键值数据
func (dc *DynamodbClient) batchInsert(data map[string]string) error {
	requestItems := make(map[string][]*dynamodb.WriteRequest)
	length := len(data)
	var writeRequests []*dynamodb.WriteRequest
	count := 0
	writeRequests = []*dynamodb.WriteRequest{}
	for key, value := range data {
		var writeRequest *dynamodb.WriteRequest
		writeRequest = &dynamodb.WriteRequest{
			PutRequest: &dynamodb.PutRequest{
				Item: map[string]*dynamodb.AttributeValue{
					"key": {
						S: aws.String(key),
					},
					"value": {
						S: aws.String(value),
					},
				},
			},
		}
		writeRequests = append(writeRequests, writeRequest)
		count = count + 1

		if (count%BatchLimit == 0) || (count >= length) {
			requestItems[dc.tableName] = writeRequests
			batchInput := &dynamodb.BatchWriteItemInput{
				RequestItems: requestItems,
			}

			if _, err := dc.dynamodb.BatchWriteItem(batchInput); err != nil {
				return err
			}
			writeRequests = []*dynamodb.WriteRequest{}
		}
	}

	return nil
}
