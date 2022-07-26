package kafka

import (
	"context"
	"time"

	"google.golang.org/protobuf/proto"

	eventpb "github.com/jinwoo1225/kafka-test-suite/gen/github.com/jinwoo1225/kafka-test-suite/event"
)

type EventProducer struct {
	DefaultProducer
}

func (ep *EventProducer) ProduceEvent(ctx context.Context, topic string, ts time.Time, eventMessage *eventpb.Message) (<-chan error, error) {
	b, err := proto.Marshal(eventMessage)
	if err != nil {
		return nil, err
	}

	return ep.DefaultProducer.Produce(ctx, topic, b, ts), nil
}
