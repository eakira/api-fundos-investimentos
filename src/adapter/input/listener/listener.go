package listener

import (
	"api-fundos-investimentos/adapter/input/controller"
	"api-fundos-investimentos/configuration/env"
	"api-fundos-investimentos/configuration/logger"
	"api-fundos-investimentos/configuration/resterrors"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/IBM/sarama"
)

func initConsume() (sarama.Consumer, *resterrors.RestErr) {
	config := sarama.NewConfig()
	logger.Info("Init Consume", "kafka")

	consumer, err := sarama.NewConsumer([]string{env.GetKafkaHost()}, config)
	if err != nil {
		logger.Error("NewConsumer err: ", err, "listener")
		return nil, resterrors.NewNotFoundError("ConsumePartition err:")

	}
	logger.Info("Finish Init Consume", "kafka")
	return consumer, nil
}

func Consume(
	fundosController controller.FundosControllerInterface,
	shutdown chan bool,
) {

	topic := env.GetTopics()
	partition := env.GetPartitions()
	topics := strings.Split(topic, ",")
	partitions := strings.Split(partition, ",")

	for _, topic := range topics {
		for _, partition := range partitions {
			value, _ := strconv.ParseInt(partition, 10, 32)
			go consumeTopic(
				topic,
				int32(value),
				fundosController,
				shutdown,
			)
		}
	}

}

func consumeTopic(
	topic string,
	partition int32,
	fundosController controller.FundosControllerInterface,
	shutdown chan bool,
) {
	logger.Info(fmt.Sprintf("Init Listener: %s", topic), "listener")

	consumer, _ := initConsume()
	defer consumer.Close()

	partitionConsumer, err := consumer.ConsumePartition(topic, partition, sarama.OffsetNewest)
	if err != nil {
		log.Fatal("ConsumePartition err: ", err)
		resterrors.NewNotFoundError("ConsumePartition err:")
	}
	defer partitionConsumer.Close()

	logger.Info(fmt.Sprintf("Init Read Message: %s %d", topic, partition), "listener")

	for message := range partitionConsumer.Messages() {

		err := switchCaseTopic(topic, message.Value, fundosController)
		if err != nil {
			logger.Error("error trying switchCaseTopic", err, "listener")
		}

		logger.Info(
			fmt.Sprintf("[Consumer] partitionid: %d; offset:%d, value: %s\n",
				message.Partition, message.Offset, string(message.Value)),
			"sincronizarFundos")

	}
	logger.Info("Finish consumeTopic", "listener")
}

func switchCaseTopic(
	topic string,
	message []byte,
	fundosController controller.FundosControllerInterface,
) *resterrors.RestErr {

	switch topic {
	case env.GetTopicSincronizar():
		return FundosSincronizarListener(message, fundosController)
	case env.GetTopicProcessarArquivos():
		return FundosProcessarArquivosListener(message, fundosController)
	case env.GetTopicPersistenciaDados():
		return FundosPersistenciaDadosListener(message, fundosController)
	default:
		return resterrors.NewNotFoundError(
			fmt.Sprintf("Queue not found: %s", topic),
		)
	}

}
