package kafka

import (
	"context"
	"testing"
	"time"

	"go.uber.org/zap"
)

func TestKafkaProducingAndConsuming(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()
	zap.ReplaceGlobals(logger)
	testKafkaBroker := []string{"localhost:9092"}

	ctx := context.Background()
	if deadline, ok := t.Deadline(); ok {
		var cancel context.CancelFunc
		ctx, cancel = context.WithDeadline(ctx, deadline)
		defer cancel()
	}

	now := time.Now()

	producerOption := ProducerOption{
		brokers:         testKafkaBroker,
		interval:        3 * time.Second,
		timeout:         3 * time.Second,
		retry:           3,
		autoCreateTopic: true,
	}
	p, err := NewDefaultProducer(producerOption)
	if err != nil {
		t.Error(err)
	}
	defer p.Close()

	topic := "test_topic"
	testMessage := "kafka TESTING"

	p.Produce(ctx, topic, []byte(testMessage), now)

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	consumerOption := ConsumerOption{
		Brokers: testKafkaBroker,
		Topics:  []string{topic},
	}

	c, err := NewDefaultConsumer(consumerOption)
	if err != nil {
		t.Error(err)
	}

	recordValue, err := c.ConsumeFirstOne(ctx, 1)
	if err != nil {
		t.Error(err)
	}

	if string(recordValue) != testMessage {
		t.Error("record value doesn't matches testMessage", recordValue, testMessage)
	}
}
