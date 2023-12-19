package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/google/uuid"
)

const tableName = "lock-table"

func GetLock(lockID string, expiredAt time.Time) (locked bool, release func(), err error) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return false, nil, err
	}
	dbClient := dynamodb.NewFromConfig(cfg)

	now := time.Now().UTC()
	releaseID := uuid.NewString()
	putItemInput := &dynamodb.PutItemInput{
		TableName: aws.String(tableName),
		Item: map[string]types.AttributeValue{
			"LockID":    &types.AttributeValueMemberS{Value: lockID},
			"ExpiredAt": &types.AttributeValueMemberS{Value: expiredAt.Format(time.RFC3339)},
			"ReleaseID": &types.AttributeValueMemberS{Value: releaseID},
			"Time":      &types.AttributeValueMemberS{Value: now.Format(time.RFC3339)},
		},
		ConditionExpression: aws.String("attribute_not_exists(LockID) OR ExpiredAt <= :now"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":now": &types.AttributeValueMemberS{Value: now.Format(time.RFC3339)},
		},
	}

	_, err = dbClient.PutItem(context.TODO(), putItemInput)
	if err != nil {
		var dynamoErr *types.ConditionalCheckFailedException
		if errors.As(err, &dynamoErr) {
			// Failed to lock as it's already locked
			return false, nil, nil
		}
		// Failed to lock due to unexpected error
		return false, nil, err
	}

	release = func() {
		deleteItemInput := &dynamodb.DeleteItemInput{
			TableName: aws.String(tableName),
			Key: map[string]types.AttributeValue{
				"LockID": &types.AttributeValueMemberS{Value: lockID},
			},
		}
		_, err := dbClient.DeleteItem(context.TODO(), deleteItemInput)
		if err != nil {
			fmt.Println("Failed to release lock:", err)
		}
	}

	return true, release, nil
}

func main() {
	expiredAt := time.Now().UTC().Add(10 * time.Second)

	locked1, release1, err := GetLock("lock1", expiredAt)
	fmt.Println("locked1:", locked1)
	if err != nil {
		panic(err)
	}

	locked2, _, err := GetLock("lock1", expiredAt)
	fmt.Println("locked2:", locked2)
	if err != nil {
		panic(err)
	}

	release1()

	locked3, _, err := GetLock("lock1", expiredAt)
	if err != nil {
		panic(err)
	}
	fmt.Println("locked3:", locked3)
}
