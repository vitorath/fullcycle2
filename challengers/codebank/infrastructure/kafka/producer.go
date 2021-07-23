package kafka

import ckafka "github.com/confluentinc/confluent-kafka-go/kafka"

type KafkaProducer struct {
	Producer *ckafka.Producer
}

func NewKafkaProducer() KafkaProducer {
	return KafkaProducer{}
}

func (k *KafkaProducer) SetupProducer(boostrapServer string) {
	configMap := &ckafka.ConfigMap{
		"boostrap.servers": boostrapServer,
	}

	k.Producer, _ = ckafka.NewProducer(configMap)
}

func (k *KafkaProducer) Publish(msg string, topic string) error {
	message := &ckafka.Message{
		TopicPartition: ckafka.TopicPartition{Topic: &topic, Partition: ckafka.PartitionAny},
		Value:          []byte(msg),
	}

	err := k.Producer.Produce(message, nil)
	if err != nil {
		return err
	}

	return nil
}
