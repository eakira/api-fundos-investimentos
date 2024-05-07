package queue

import (
	"api-fundos-investimentos/adapter/output/model/entity"
	"api-fundos-investimentos/configuration/logger"
	"api-fundos-investimentos/configuration/resterrors"
	"fmt"

	"github.com/IBM/sarama"
)

func (nc *queueProduce) CreateTopic(topicEntity entity.TopicEntity) *resterrors.RestErr {
	kafkaAdmin, err := NewKafkaAdminClient()
	if err != nil {
		return resterrors.NewInternalServerError("Error while creating Kafka admin client")
	}
	defer kafkaAdmin.Close()

	topics, err := kafkaAdmin.ListTopics()
	if err != nil {
		return resterrors.NewInternalServerError("Error listing Kafka topics")
	}

	if _, exists := topics[topicEntity.Topic]; !exists {
		topicDetail := &sarama.TopicDetail{
			NumPartitions:     int32(topicEntity.Qtd),
			ReplicationFactor: 1,
		}

		err = kafkaAdmin.CreateTopic(topicEntity.Topic, topicDetail, false)
		if err != nil {
			return resterrors.NewInternalServerError("Error creating Kafka topic")
		}
	}

	updatedTopics, _ := kafkaAdmin.ListTopics()
	fmt.Println("Updated Topics:", updatedTopics)

	logger.Info("Finish CreateTopic", "sincronizarFundos")
	return nil
}
