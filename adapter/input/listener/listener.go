package listener

import (
	"api-fundos-investimentos/configuration/env"
	"api-fundos-investimentos/configuration/logger"
	"api-fundos-investimentos/configuration/resterrors"
	"fmt"
	"log"

	"github.com/IBM/sarama"
)

func initConsume() (sarama.Consumer, *resterrors.RestErr) {
	config := sarama.NewConfig()
	logger.Info(env.GetKafkaHost(), "kafka")

	consumer, err := sarama.NewConsumer([]string{env.GetKafkaHost()}, config)
	if err != nil {
		logger.Error("NewConsumer err: ", err, "listener")
		return nil, resterrors.NewNotFoundError("ConsumePartition err:")

	}
	return consumer, nil
}

func Consume(topic string) {
	logger.Info(fmt.Sprintf("Init Listener: %s", topic), "listener")
	consumer, _ := initConsume()
	defer consumer.Close()

	partitionConsumer, err := consumer.ConsumePartition(topic, 0, sarama.OffsetNewest)
	if err != nil {
		log.Fatal("ConsumePartition err: ", err)
		resterrors.NewNotFoundError("ConsumePartition err:")
	}
	defer partitionConsumer.Close()

	for message := range partitionConsumer.Messages() {
		err := switchCaseTopic(topic, string(message.Value))
		if err != nil {
			logger.Error("error trying switchCaseTopic", err, "listener")
		}

		logger.Info(
			fmt.Sprintf("[Consumer] partitionid: %d; offset:%d, value: %s\n",
				message.Partition, message.Offset, string(message.Value)),
			"sincronizarFundos")

	}
	logger.Info("Finish SincronizarFundos", "sincronizarFundos")
}

func switchCaseTopic(topic string, message string) *resterrors.RestErr {

	switch topic {
	case env.GetTopicSincronizar():
		return FundosSincroniszarListener(message)
	default:
		return resterrors.NewNotFoundError(
			fmt.Sprintf("Queue not found: %s", topic),
		)

	}

}
