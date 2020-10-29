package aws

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/codingnagger/cloudready-todolist-blog-example/pkg/completion"
	"github.com/codingnagger/cloudready-todolist-blog-example/pkg/creation"
	"github.com/codingnagger/cloudready-todolist-blog-example/pkg/listing"
	"github.com/google/uuid"
)

var (
	tableName *string
	awsConfig *aws.Config
)

// Storage keeps tasks data in memory
type Storage struct {
}

// NewStorage creates and return a new storage
func NewStorage(config TodoConfig) *Storage {
	c := aws.NewConfig()

	if len(config.Region) > 0 {
		c = c.WithRegion(config.Region)
	}

	if len(config.Endpoint) > 0 {
		c = c.WithEndpoint(config.Endpoint)
	}

	tableName = &config.TableName
	awsConfig = c

	return new(Storage)
}

// NewDynamoDBClient creates a new DynamoDB client
func (s Storage) NewDynamoDBClient() *dynamodb.DynamoDB {
	session := session.Must(session.NewSession(awsConfig))
	return dynamodb.New(session)
}

// CreateTask adds a new task in memory
func (s *Storage) CreateTask(task creation.Task) error {
	svc := s.NewDynamoDBClient()

	newTask := Task{
		ID:          uuid.New().String(),
		Description: task.Description,
		Completed:   false,
	}

	t, err := dynamodbattribute.MarshalMap(newTask)

	if err != nil {
		return err
	}

	_, err = svc.PutItem(&dynamodb.PutItemInput{
		TableName: tableName,
		Item:      t,
	})

	return err
}

// CompleteTask marks a task as completed
func (s *Storage) CompleteTask(task completion.Task) error {
	svc := s.NewDynamoDBClient()

	input := &dynamodb.UpdateItemInput{
		TableName: tableName,
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(task.ID),
			},
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":completed": {
				BOOL: aws.Bool(true),
			},
		},
		UpdateExpression: aws.String("SET completed = :completed"),
		ReturnValues:     aws.String("UPDATED_OLD"),
	}

	res, err := svc.UpdateItem(input)

	if err != nil {
		if aerr, ok := err.(awserr.Error); ok && aerr.Code() == dynamodb.ErrCodeResourceNotFoundException {
			return completion.ErrorTaskNotFound
		}
		return err
	}

	if *res.Attributes["completed"].BOOL {
		return completion.ErrorTaskAlreadyCompleted
	}

	return nil
}

// ListCompletedTasks lists all the tasks marked as completed
func (s *Storage) ListCompletedTasks() (listing.ListedTasks, error) {
	return s.listTasksFromStatus(true)
}

// ListPendingTasks lists all the tasks not marked as completed
func (s *Storage) ListPendingTasks() (listing.ListedTasks, error) {
	return s.listTasksFromStatus(false)
}

func (s *Storage) listTasksFromStatus(completed bool) (listing.ListedTasks, error) {
	input := &dynamodb.ScanInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":completed": {
				BOOL: aws.Bool(completed),
			},
		},
		FilterExpression:     aws.String("completed = :completed"),
		ProjectionExpression: aws.String("id, description, completed"),
		TableName:            tableName,
	}

	return s.collectResult(input)
}

// ListAllTasks returns all tasks
func (s *Storage) ListAllTasks() (listing.ListedTasks, error) {
	input := &dynamodb.ScanInput{
		ProjectionExpression: aws.String("id, description, completed"),
		TableName:            tableName,
	}

	return s.collectResult(input)
}

func (s *Storage) collectResult(input *dynamodb.ScanInput) (listing.ListedTasks, error) {
	svc := s.NewDynamoDBClient()

	r, err := svc.Scan(input)

	if err != nil {
		fmt.Printf("%s\n", err)
		return nil, err
	}

	result := listing.ListedTasks{}

	for _, item := range r.Items {
		var t listing.Task

		dynamodbattribute.UnmarshalMap(item, &t)

		task := listing.Task{
			ID:          t.ID,
			Description: t.Description,
		}

		result = append(result, task)
	}

	return result, nil
}
