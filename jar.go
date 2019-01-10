package awssqsjar

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/aws/aws-sdk-go/service/sqs/sqsiface"
	"github.com/cookiejars/cookiejar"
)

type jar struct {
	sqsiface.SQSAPI
}

func New(session *session.Session) (cookiejar.Jar, error) {
	return &jar{
		SQSAPI: sqs.New(session),
	}, nil
}

func (jar) Retrieve() ([]cookiejar.Cookie, error) {
	panic("implement me")
}

func (jar) Retire(cookie cookiejar.Cookie) error {
	panic("implement me")
}
