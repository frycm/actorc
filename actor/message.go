package actor

import (
	"context"
	"time"
)

type Message struct {
	Ctx context.Context
	Sent time.Time
	Intention interface{}
	ResponseChan chan<- Message
}