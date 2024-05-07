package controller

import (
	"api-fundos-investimentos/application/domain"
	"api-fundos-investimentos/configuration/env"
	"api-fundos-investimentos/configuration/logger"
)

func (fc *fundosControllerInterface) CreateTopicController() {
	logger.Info("Init CreateTopicController", "sincronizarFundos")

	qtd := env.GetNumParticoes()

	var topics = []string{
		env.GetTopicSincronizar(),
		env.GetTopicProcessarArquivos(),
		env.GetTopicPersistenciaDados(),
	}

	for _, topic := range topics {
		_, erro := fc.service.CreateTopicService(domain.TopicDomain{
			Topic: topic,
			Qtd:   qtd,
		})
		if erro != nil {
			logger.Error("Erro ao salvar o tópico", nil, "sincronizarFundos")
		}
	}

	logger.Info("Finish CreateTopicController", "sincronizarFundos")

}
