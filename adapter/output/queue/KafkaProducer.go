package queue

import (
	"api-fundos-investimentos/adapter/output/model/response"
	"api-fundos-investimentos/configuration/env"
	"api-fundos-investimentos/configuration/logger"
	"api-fundos-investimentos/configuration/resterrors"
	"encoding/json"
	"fmt"

	"github.com/IBM/sarama"
)

type queueProduce struct{}

func NewProduce() *queueProduce {

	return &queueProduce{}
}

func (nc *queueProduce) Produce(response response.FundosQueueResponse) *resterrors.RestErr {
	logger.Info("Init Kafka Produce", "kafkaProduce")

	producer, _ := initProduce()
	defer producer.Close()

	dtoJson, _ := json.Marshal(response)
	msg := &sarama.ProducerMessage{Topic: response.Topic, Key: nil, Value: sarama.StringEncoder(dtoJson)}
	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		logger.Error("SendMessage err:", err, "kafkaProduce")
		return resterrors.NewNotFoundError("Failed to send message, err:")
	}

	logger.Info(
		fmt.Sprintf("Finish Kafka Produce partition id: %d; offset:%d, value: %s\n", partition, offset, response.Queue),
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
