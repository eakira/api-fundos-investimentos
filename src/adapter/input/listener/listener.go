package listener

import (
	"api-fundos-investimentos/adapter/input/controller"
	"api-fundos-investimentos/configuration/env"
	"api-fundos-investimentos/configuration/logger"
	"api-fundos-investimentos/configuration/resterrors"
	"fmt"
	"strconv"
	"strings"
	"sync"

	"github.com/IBM/sarama"
)

// initConsume inicializa o consumidor Kafka
func initConsume() (sarama.Consumer, *resterrors.RestErr) {
	config := sarama.NewConfig()
	logger.Info("Init Consume", "kafka")

	consumer, err := sarama.NewConsumer([]string{env.GetKafkaHost()}, config)
	if err != nil {
		logger.Error("NewConsumer err: ", err, "kafka")
		return nil, resterrors.NewInternalServerError("Failed to create Kafka consumer")
	}
	logger.Info("Finish Init Consume", "kafka")
	return consumer, nil
}

// Consume inicia a escuta nos tópicos Kafka
func Consume(fundosController controller.FundosControllerInterface, shutdown chan bool) {
	topic := env.GetTopics()
	partition := env.GetPartitions()
	topics := strings.Split(topic, ",")
	partitions := strings.Split(partition, ",")

	var wg sync.WaitGroup

	for _, topic := range topics {
		for _, partition := range partitions {
			value, _ := strconv.ParseInt(partition, 10, 32)
			wg.Add(1)
			go consumeTopic(topic, int32(value), fundosController, shutdown, &wg)
		}
	}

	wg.Wait()
}

// consumeTopic consome mensagens de um tópico específico e uma partição específica
func consumeTopic(topic string, partition int32, fundosController controller.FundosControllerInterface, shutdown chan bool, wg *sync.WaitGroup) {
	defer wg.Done()

	logger.Info(fmt.Sprintf("Init Listener: %s", topic), "kafka")

	consumer, err := initConsume()
	if err != nil {
		logger.Error("Failed to initialize Kafka consumer: ", err, "kafka")
		return
	}
	defer consumer.Close()

	partitionConsumer, err2 := consumer.ConsumePartition(topic, partition, sarama.OffsetNewest)
	if err2 != nil {
		logger.Error("Failed to consume partition: ", err, "kafka")
		return
	}
	defer partitionConsumer.Close()

	logger.Info(fmt.Sprintf("Init Read Message: %s %d", topic, partition), "kafka")

	for {
		select {
		case <-shutdown:
			logger.Info(fmt.Sprintf("Shutting down listener for topic %s partition %d", topic, partition), "kafka")
			return
		case message := <-partitionConsumer.Messages():
			err := switchCaseTopic(topic, message.Value, fundosController)
			if err != nil {
				logger.Error("Error processing message: ", err, "kafka")
			}
			logger.Info(fmt.Sprintf("[Consumer] partitionid: %d; offset:%d, value: %s\n", message.Partition, message.Offset, string(message.Value)), "listener")
		}
	}
}

// switchCaseTopic decide qual função de tratamento de mensagem invocar com base no tópico
func switchCaseTopic(topic string, message []byte, fundosController controller.FundosControllerInterface) *resterrors.RestErr {
	switch topic {
	case env.GetTopicSincronizar():
		return FundosSincronizarListener(message, fundosController)
	case env.GetTopicProcessarArquivos():
		return FundosProcessarArquivosListener(message, fundosController)
	case env.GetTopicPersistenciaDados():
		return FundosPersistenciaDadosListener(message, fundosController)
	default:
		return resterrors.NewNotFoundError(fmt.Sprintf("Queue not found: %s", topic))
	}
}
