package queue

import (
	"api-fundos-investimentos/configuration/env"
	"api-fundos-investimentos/configuration/logger"
	"api-fundos-investimentos/configuration/resterrors"
	"api-fundos-investimentos/core/dto"
	"fmt"

	"github.com/IBM/sarama"
)

type queueProduce struct{}

func NewProduce() *queueProduce {

	return &queueProduce{}
}

func (nc *queueProduce) Produce(dto dto.FundosQueueDto) *resterrors.RestErr {
	logger.Info("Init Kafka Produce", "kafkaProduce")

	producer, _ := initProduce()
	defer producer.Close()

	msg := &sarama.ProducerMessage{Topic: dto.Topic, Key: nil, Value: sarama.StringEncoder(dto.Queue)}
	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		logger.Error("SendMessage err:", err, "kafkaProduce")
		return resterrors.NewNotFoundError("Failed to send message, err:")
	}

	logger.Info(
		fmt.Sprintf("Finish Kafka Produce partition id: %d; offset:%d, value: %s\n", partition, offset, dto.Queue),
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
