package config

import (
	"github.com/segmentio/kafka-go"
)

var brokers = []string{"localhost:9093"}

func GetWriter(topic string) *kafka.Writer {
	return kafka.NewWriter(kafka.WriterConfig{
		Brokers: brokers,
		Topic:   topic,
	})
}