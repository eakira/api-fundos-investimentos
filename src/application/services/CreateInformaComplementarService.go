package services

import (
	"api-fundos-investimentos/application/domain"
	"api-fundos-investimentos/configuration/logger"
	"api-fundos-investimentos/configuration/resterrors"
)

func (fs *fundosDomainService) CreateInformacaoComplementarFundoService(informacacoesFundoDomain []domain.InformacoesFundoDomain) *resterrors.RestErr {
	logger.Info("Init CreateInformacaoComplementarFundo", "sincronizarFundos")

	if err := fs.repository.CreateManyInformacaoComplementarFundoRepository(informacacoesFundoDomain); err != nil {
		logger.Error("Error calling CreateManyInformacaoComplementarFundoRepository", err, "sincronizarFundos")
		return err
	}

	logger.Info("Finish CreateInformacaoComplementarFundo", "sincronizarFundos")
	return nil
}

func (fs *fundosDomainService) CreateInformacaoComplementarDivulgacaoService(informacoesDivulgacaoDomain []domain.InformacoesDivulgacaoDomain) *resterrors.RestErr {
	logger.Info("Init CreateInformacaoComplementarDivulgacao", "sincronizarFundos")

	if err := fs.repository.CreateManyInformacaoComplementarDivulgacaoRepository(informacoesDivulgacaoDomain); err != nil {
		logger.Error("Error calling CreateManyInformacaoComplementarFundoRepository", err, "sincronizarFundos")
		return err
	}

	logger.Info("Finish CreateInformacaoComplementarDivulgacao", "sincronizarFundos")
	return nil
}

func (fs *fundosDomainService) CreateInformacaoComplementarCotistaService(domain []domain.InformacoesCotistaDomain) *resterrors.RestErr {
	logger.Info("Init CreateInformacaoComplementarCotista", "sincronizarFundos")

	if err := fs.repository.CreateManyInformacaoComplementarCotistaRepository(domain); err != nil {
		logger.Error("Error calling CreateManyInformacaoComplementarFundoRepository", err, "sincronizarFundos")
		return err
	}

	logger.Info("Finish CreateInformacaoComplementarCotista", "sincronizarFundos")
	return nil
}

func (fs *fundosDomainService) CreateInformacaoComplementarServicoPrestadoService(servicoPrestadoDomain []domain.ServicoPrestadoDomain) *resterrors.RestErr {
	logger.Info("Init CreateInformacaoComplementarServicoPrestado", "sincronizarFundos")

	if err := fs.repository.CreateManyInformacaoComplementarServicoPrestadoRepository(servicoPrestadoDomain); err != nil {
		logger.Error("Error calling CreateManyInformacaoComplementarFundoRepository", err, "sincronizarFundos")
		return err
	}

	logger.Info("Finish CreateInformacaoComplementarServicoPrestado", "sincronizarFundos")
	return nil
}
