package listener

import (
	"api-fundos-investimentos/adapter/input/controller"
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

func Consume(topic string, partition int32, fundosController controller.FundosControllerInterface) {
	logger.Info(fmt.Sprintf("Init Listener: %s", topic), "listener")
	consumer, _ := initConsume()
	defer consumer.Close()

	partitionConsumer, err := consumer.ConsumePartition(topic, partition, sarama.OffsetNewest)
	if err != nil {
		log.Fatal("ConsumePartition err: ", err)
		resterrors.NewNotFoundError("ConsumePartition err:")
	}
	defer partitionConsumer.Close()

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
	logger.Info("Finish SincronizarFundos", "sincronizarFundos")
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
