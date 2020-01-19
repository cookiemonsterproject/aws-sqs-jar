package awssqsjar

import (
	"github.com/aws/aws-sdk-go/service/sqs"

	"github.com/cookiemonsterproject/cookie-monster"
)

const receiptHandle = "receipt_handle"

type cookie struct {
	*sqs.Message
}

func newCookies(messages []*sqs.Message) []cookiemonster.Cookie {
	ret := make([]cookiemonster.Cookie, len(messages))
	for i, m := range messages {
		ret[i] = newCookie(m)
	}
	return ret
}

func newCookie(message *sqs.Message) cookiemonster.Cookie {
	return &cookie{Message: message}
}

func (c cookie) ID() string {
	return *c.MessageId
}

func (c cookie) Content() interface{} {
	return *c.Body
}

func (c cookie) Metadata() map[string]string {
	return map[string]string{
		receiptHandle: *c.ReceiptHandle,
	}
}
