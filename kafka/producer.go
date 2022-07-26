package kafka

import (
	"context"
	"time"

	"github.com/twmb/franz-go/pkg/kgo"
	"go.uber.org/zap"
)

type ProducerOption struct {
	brokers         []string
	interval        time.Duration
	timeout         time.Duration
	retry           int
	autoCreateTopic bool
}

type DefaultProducer struct {
	client *kgo.Client
}

func NewDefaultProducer(option ProducerOption) (*DefaultProducer, error) {
	zap.L().Sugar().Debug(option.brokers)
	options := []kgo.Opt{
		kgo.SeedBrokers(option.brokers...),
		kgo.ProduceRequestTimeout(option.timeout),
		kgo.RecordDeliveryTimeout(option.interval),
		kgo.RecordRetries(option.retry),
	}

	if option.autoCreateTopic {
		options = append(options, kgo.AllowAutoTopicCreation())
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

	return &DefaultProducer{client: client}, nil
}

func (d *DefaultProducer) Produce(ctx context.Context, topic string, data []byte, ts time.Time) {
	record := &kgo.Record{
		Topic:     topic,
		Timestamp: ts,
		Value:     data,
	}
	d.client.Produce(ctx, record, func(record *kgo.Record, err error) {
		if err != nil {
			zap.L().Sugar().Error(err)
		} else {
			zap.L().Sugar().Debug("record sent : ", string(record.Value))
		}
	})
}

func (d *DefaultProducer) Close() {
	d.client.Close()
}
