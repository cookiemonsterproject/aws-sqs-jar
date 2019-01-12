package awssqsjar

import (
	"errors"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/aws/aws-sdk-go/service/sqs/sqsiface"
	"github.com/cookiejars/cookiejar"
)

type Config struct {
	QueueName         string
	MaxNumberMessages int64
}

type jar struct {
	sqsiface.SQSAPI
	cfg Config
}

func NewWithSession(session *session.Session, config Config) (cookiejar.Jar, error) {
	return &jar{
		SQSAPI: sqs.New(session),
		cfg:    config,
	}, nil
}

func (j jar) Retrieve() ([]cookiejar.Cookie, error) {
	queueURL, err := j.getQueueURL()
	if err != nil {
		return nil, err
	}

	receiveInput := &sqs.ReceiveMessageInput{
		MaxNumberOfMessages: aws.Int64(j.cfg.MaxNumberMessages),
		QueueUrl:            aws.String(queueURL),
	}

	result, err := j.ReceiveMessage(receiveInput)
	if err != nil {
		return nil, err
	}

	return newCookies(result.Messages), nil
}

func (j jar) Retire(cookie cookiejar.Cookie) error {
	queueURL, err := j.getQueueURL()
	if err != nil {
		return err
	}

	metadata := cookie.Metadata()
	receiptHandle := metadata[receiptHandle]

	inputDelete := &sqs.DeleteMessageInput{QueueUrl: &queueURL, ReceiptHandle: &receiptHandle}

	_, err = j.DeleteMessage(inputDelete)

	return err
}

func (j jar) getQueueURL() (string, error) {
	input := &sqs.GetQueueUrlInput{QueueName: &j.cfg.QueueName}

	output, err := j.GetQueueUrl(input)
	if err != nil {
		return "", err
	}

	if output.QueueUrl == nil {
		return "", errors.New("queue url is nil")
	}

	return *output.QueueUrl, nil
}
