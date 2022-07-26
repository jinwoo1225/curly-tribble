package config

import "github.com/jinwoo1225/kafka-test-suite/kafka"

type Config interface {
	GetConfig() *BaseConfig
}

type BaseConfig struct {
	Producer *kafka.ProducerOption
	Consumer *kafka.ConsumerOption
}

func (c BaseConfig) GetConfig() *BaseConfig {
	return &c
}
