package input

type FundosDomainService interface {
	QueueFundosSincronizarService(string)
	GetFundosExternoService()
	DownloadArquivosCVMService(file string)
}
