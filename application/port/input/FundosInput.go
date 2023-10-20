package input

type FundosDomainService interface {
	QueueFundosSincronizarService(string, string)
	GetFundosExternoService()
}
