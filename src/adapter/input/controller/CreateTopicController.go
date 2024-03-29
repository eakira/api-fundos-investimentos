package controller

import (
	"api-fundos-investimentos/configuration/env"
	"api-fundos-investimentos/configuration/logger"
	"fmt"
	"log"

	"github.com/IBM/sarama"
	"github.com/gin-gonic/gin"
)

func (fc *fundosControllerInterface) CreateTopicController(c *gin.Context) {
	logger.Info("Init CreateTopicController", "sincronizarFundos")

	brokerAddrs := []string{env.GetKafkaHost()}
	config := sarama.NewConfig()

	admin, err := sarama.NewClusterAdmin(brokerAddrs, config)
	if err != nil {
		log.Fatal("Error while creating cluster admin: ", err.Error())
	}
	defer admin.Close()

	admin.CreateTopic(env.GetTopicProcessarArquivos(), &sarama.TopicDetail{
		NumPartitions:     env.GetNumParticoes(),
		ReplicationFactor: 1,
	}, false)

	admin.CreateTopic(env.GetTopicPersistenciaDados(), &sarama.TopicDetail{
		NumPartitions:     env.GetNumParticoes(),
		ReplicationFactor: 1,
	}, false)

	teste, _ := admin.ListTopics()
	fmt.Println(teste)
	logger.Info("Finish CreateTopicController", "sincronizarFundos")

}
