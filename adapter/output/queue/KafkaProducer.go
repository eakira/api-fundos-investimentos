package queue

import (
	"api-fundos-investimentos/configuration/env"
	"api-fundos-investimentos/configuration/logger"
	"api-fundos-investimentos/configuration/resterrors"
	"fmt"

	"github.com/IBM/sarama"
)

type queueProduce struct{}

func NewProduce() *queueProduce {

	return &queueProduce{}
}

func (nc *queueProduce) Produce(topic string, message string) *resterrors.RestErr {
	logger.Info("Init Kafka Produce", "kafkaProduce")

	producer, _ := initProduce()
	defer producer.Close()

	msg := &sarama.ProducerMessage{Topic: topic, Key: nil, Value: sarama.StringEncoder(message)}
	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		logger.Error("SendMessage err:", err, "kafkaProduce")
		return resterrors.NewNotFoundError("Failed to send message, err:")
	}

	logger.Info(
		fmt.Sprintf("Finish Kafka Produce partition id: %d; offset:%d, value: %s\n", partition, offset, message),
		"kafkaProduce",
	)
	return nil
}

func initProduce() (sarama.SyncProducer, *resterrors.RestErr) {
	KAFKA_HOST := env.GetKafkaHost()

	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true

	producer, err := sarama.NewSyncProducer([]string{KAFKA_HOST}, config)
	if err != nil {
		logger.Error("Failed to initialize NewSyncProducer, err:", err, "kafkaProduce")
		return nil, resterrors.NewNotFoundError("Failed to initialize NewSyncProducer, err:")
	}
	return producer, nil
}
