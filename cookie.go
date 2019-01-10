package awssqsjar

import (
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/cookiejars/cookiejar"
)

type cookie struct {
	*sqs.Message
}

func newCookies(messages []*sqs.Message) []cookiejar.Cookie {
	ret := make([]cookiejar.Cookie, len(messages))
	for i, m := range messages {
		ret[i] = newCookie(m)
	}
	return ret
}

func newCookie(message *sqs.Message) cookiejar.Cookie {
	return &cookie{message}
}

func (c cookie) ID() string {
	return *c.MessageId
}

func (c cookie) Content() interface{} {
	return *c.Body
}
