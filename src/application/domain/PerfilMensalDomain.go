package domain

type PerfilMensalDomain struct {
	CenarioFPRCupom                    string // Cenário de fator de risco para cupom
	CenarioFPRDolar                    string // Cenário de fator de risco para dólar
	CenarioFPRIBovespa                 string // Cenário de fator de risco para Ibovespa
	CenarioFPRJuros                    string // Cenário de fator de risco para juros
	CenarioFPROutro                    string // Outro cenário de fator de risco
	FundoCnpj                          string // CNPJ do fundo
	ComitenteLigado1                   string // Comitente ligado 1
	ComitenteLigado2                   string // Comitente ligado 2
	ComitenteLigado3                   string // Comitente ligado 3
	CPFCNPJComitente1                  string // CPF/CNPJ do comitente 1
	CPFCNPJComitente2                  string // CPF/CNPJ do comitente 2
	CPFCNPJComitente3                  string // CPF/CNPJ do comitente 3
	CPFCNPJEmissor1                    string // CPF/CNPJ do emissor 1
	CPFCNPJEmissor2                    string // CPF/CNPJ do emissor 2
	CPFCNPJEmissor3                    string // CPF/CNPJ do emissor 3
	DelibAssemb                        string // Deliberação da assembleia
	DenomSocial                        string // Denominação social
	DataComptc                         string // Data de competência
	DataCotaTaxaPerfm                  string // Data da cota de taxa de performance
	EmissorLigado1                     string // Emissor ligado 1
	EmissorLigado2                     string // Emissor ligado 2
	EmissorLigado3                     string // Emissor ligado 3
	FatorRiscoNocional                 string // Fator de risco nocional
	FatorRiscoOutro                    string // Outro fator de risco
	FPR                                string // FPR
	JustifVotoAdminAssemb              string // Justificativa de voto na assembleia administrativa
	ModVar                             string // Modelo variável
	NrCotstBanco                       string // Número de cotistas de banco
	NrCotstCapitaliz                   string // Número de cotistas de capitalização
	NrCotstCorretoraDistrib            string // Número de cotistas de corretora distribuidora
	NrCotstDistrib                     string // Número de cotistas distribuidores
	NrCotstEAPC                        string // Número de cotistas EAPC
	NrCotstEFPC                        string // Número de cotistas EFPC
	NrCotstEntidPrevidCompl            string // Número de cotistas entidade de previdência complementar
	NrCotstFIClube                     string // Número de cotistas de FIC clube
	NrCotstInvnr                       string // Número de cotistas investidores não residentes
	NrCotstOutro                       string // Número de cotistas de outro tipo
	NrCotstPFPB                        string // Número de cotistas de pessoa física PB
	NrCotstPFVarejo                    string // Número de cotistas de pessoa física varejo
	NrCotstPJFinanc                    string // Número de cotistas de pessoa jurídica financeira
	NrCotstPJNaoFinancPB               string // Número de cotistas de pessoa jurídica não financeira PB
	NrCotstPJNaoFinancVarejo           string // Número de cotistas de pessoa jurídica não financeira varejo
	NrCotstRPPS                        string // Número de cotistas RPPS
	NrCotstSegur                       string // Número de cotistas de seguradora
	NrDiaCemPerc                       string // Número de dias 100% de aplicação
	NrDiaCinquPerc                     string // Número de dias 50% de aplicação
	PFPJComitente1                     string // PF/PJ comitente 1
	PFPJComitente2                     string // PF/PJ comitente 2
	PFPJComitente3                     string // PF/PJ comitente 3
	QtdAcoesEmitidasMaiorQtdAutoriz    string // Quantidade de ações emitidas maior que quantidade autorizada
	QtdCotstEmit                       string // Quantidade de cotistas emitidas
	QtdCotstResgatadas                 string // Quantidade de cotistas resgatadas
	QtdDiasCorridos                    string // Quantidade de dias corridos
	QtdEmissaoResgate                  string // Quantidade de emissão/resgate
	QtdEmissaoResgateAgrupado          string // Quantidade de emissão/resgate agrupado
	QtdRecolhimentoIntegralizacao      string // Quantidade de recolhimento para integralização
	QtdReembolsoDividendosJSCP         string // Quantidade de reembolso de dividendos JSCP
	QtdReembolsoDividendosJSCPAgrupado string // Quantidade de reembolso de dividendos JSCP agrupado
	QtdReembolsoJurosCP                string // Quantidade de reembolso de juros CP
	QtdReembolsoJurosCPAgrupado        string // Quantidade de reembolso de juros CP agrupado
	QtdReembolsoRendimentos            string // Quantidade de reembolso de rendimentos
	QtdReembolsoRendimentosAgrupado    string // Quantidade de reembolso de rendimentos agrupado
	QtdResgateIntegralizacao           string // Quantidade de resgate para integralização
	QtdResgateParcial                  string // Quantidade de resgate parcial
	QtdRevolucaoCotst                  string // Quantidade de revolução de cotistas
	QtdRendimentosAuferidos            string // Quantidade de rendimentos auferidos
	QtdResgates                        string // Quantidade de resgates
	QtdSuspensaoCota                   string // Quantidade de suspensão de cota
	QtdVendaCota                       string // Quantidade de venda de cota
	QtdAmortizacaoPrecoAdquirido       string // Quantidade de amortização de preço adquirido
	QtdReinvestimento                  string // Quantidade de reinvestimento
	QtdSubscricao                      string // Quantidade de subscrição
	TipoComitente1                     string // Tipo de comitente 1
	TipoComitente2                     string // Tipo de comitente 2
	TipoComitente3                     string // Tipo de comitente 3
	TipoControleGestao                 string // Tipo de controle de gestão
	TipoFundo                          string // Tipo de fundo
	TipoRentabFund                     string // Tipo de rentabilidade do fundo
	VencimentoCotas                    string // Vencimento das cotas
}
