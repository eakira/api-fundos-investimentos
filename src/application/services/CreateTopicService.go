package services

import (
	"api-fundos-investimentos/adapter/output/model/entity"
	"api-fundos-investimentos/application/domain"
	"api-fundos-investimentos/configuration/logger"
	"api-fundos-investimentos/configuration/resterrors"

	"github.com/jinzhu/copier"
)

func (fs *fundosDomainService) CreateTopicService(
	topicDomain domain.TopicDomain,
) (
	*domain.TopicDomain,
	*resterrors.RestErr,
) {
	logger.Info("Init CreateTopicService", "sincronizarFundos")

	entity := entity.TopicEntity{}
	copier.Copy(&entity, &topicDomain)

	err := fs.queue.CreateTopic(entity)
	if err != nil {
		logger.Error("Error calling CreateTopic", err, "sincronizarFundos")
		return nil, err
	}

	logger.Info("Finish CreateTopicService", "sincronizarFundos")
	return nil, nil

}
