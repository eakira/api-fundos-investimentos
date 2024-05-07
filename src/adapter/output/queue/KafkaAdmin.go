package queue

import (
	"api-fundos-investimentos/configuration/env"

	"github.com/IBM/sarama"
)

// TODO: ADICIONAR NA CAMADA DE INFRA
type KafkaAdmin interface {
	ListTopics() (map[string]sarama.TopicDetail, error)
	CreateTopic(topic string, detail *sarama.TopicDetail, validateOnly bool) error
	Close() error
}

type KafkaAdminClient struct {
	Admin sarama.ClusterAdmin
}

func NewKafkaAdminClient() (KafkaAdmin, error) {
	brokerAddrs := []string{env.GetKafkaHost()}
	config := sarama.NewConfig()

	admin, err := sarama.NewClusterAdmin(brokerAddrs, config)
	if err != nil {
		return nil, err
	}

	return &KafkaAdminClient{Admin: admin}, nil
}

func (kc *KafkaAdminClient) ListTopics() (map[string]sarama.TopicDetail, error) {
	return kc.Admin.ListTopics()
}

func (kc *KafkaAdminClient) CreateTopic(topic string, detail *sarama.TopicDetail, validateOnly bool) error {
	return kc.Admin.CreateTopic(topic, detail, validateOnly)
}

func (kc *KafkaAdminClient) Close() error {
	return kc.Admin.Close()
}
