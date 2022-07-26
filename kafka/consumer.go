package kafka

import (
	"context"
	"errors"

	"github.com/twmb/franz-go/pkg/kgo"
)

var recordCountNotMatchedErr = errors.New("record count doesn't match")

type ConsumerOption struct {
	Brokers []string
	Topics  []string
}

type DefaultConsumer struct {
	client *kgo.Client
}

func NewDefaultConsumer(option ConsumerOption) (*DefaultConsumer, error) {
	options := []kgo.Opt{
		kgo.SeedBrokers(option.Brokers...),
		kgo.ConsumeTopics(option.Topics...),
	}

	client, err := kgo.NewClient(options...)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), pingTimeout)
	defer cancel()
	if err = client.Ping(ctx); err != nil {
		return nil, err
	}

	return &DefaultConsumer{client: client}, nil
}

func (c *DefaultConsumer) ConsumeFirstOne(ctx context.Context, recordCount int) ([]byte, error) {
	fetches := c.client.PollRecords(ctx, recordCount)
	records := fetches.Records()
	if len(records) != recordCount {
		return nil, recordCountNotMatchedErr
	}

	record := records[0]
	return record.Value, nil
}
