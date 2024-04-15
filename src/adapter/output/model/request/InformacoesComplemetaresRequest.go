package request

type InformacoesFundoResquest struct {
	AgenciaRisco               string `json:"AG_RISCO,omitempty"`                  // Nome da agência de classificação de risco (100 caracteres)
	ApresentacaoAdministrador  string `json:"APRES_ADMIN,omitempty"`               // Apresentação do administrador
	ApresentacaoGestor         string `json:"APRES_GESTOR,omitempty"`              // Apresentação do gestor de recursos
	CnpjAgenciaRisco           Cnpj   `json:"CNPJ_AG_RISCO,omitempty"`             // CNPJ da agência de classificação de risco (20 caracteres)
	FundoCnpj                  Cnpj   `json:"CNPJ_FUNDO,omitempty"`                // CNPJ do fundo (20 caracteres)
	DisclaimerAgenciaRisco     string `json:"DISCLAIMER_AG_RISCO,omitempty"`       // Padronização do disclaimer relativo à advertência sobre a manutenção do serviço de classificação de risco (100 caracteres)
	DistribuidorLigado         string `json:"DISTRIB_LIGADO,omitempty"`            // Indica se o distribuidor oferta para o público alvo do fundo, preponderantemente, fundos geridos por um único gestor ou por gestoras ligadas a um mesmo grupo econômico (1 caractere)
	DataCompetencia            Date   `json:"DT_COMPTC,omitempty"`                 // Data de competência do documento (AAAA-MM-DD)
	GrauRisco                  string `json:"GRAU_RISCO,omitempty"`                // Grau de risco atribuído (50 caracteres)
	InformacoesAutorregulacao  string `json:"INF_AUTORREGULACAO_ANBIMA,omitempty"` // Informações sobre autorregulação ANBIMA
	OutrasInformacoes          string `json:"OUTRO_INF,omitempty"`                 // Demais informações relevantes
	PeriodoMinimoDivulgacao    string `json:"PERIODO_MIN_DIVULG_CART,omitempty"`   // Periodicidade mínima para divulgação da composição da carteira do fundo (250 caracteres)
	PoliticaAdministracaoRisco string `json:"POLIT_ADM_RISCO,omitempty"`           // Política de administração de risco, em especial dos métodos utilizados pelo administrador para gerenciar os riscos a que o fundo se encontra sujeito, inclusive risco de liquidez
	PoliticaDistribuicaoCotas  string `json:"POLIT_DISTRIB,omitempty"`             // Política de distribuição de cotas
	PoliticaVoto               string `json:"POLIT_VOTO,omitempty"`                // Política relativa ao exercício de direito do voto
	RiscoCarteira              string `json:"RISCO_CART,omitempty"`                // Exposição, em ordem de relevância, dos fatores de riscos inerentes à composição da carteira do fundo
	ExistenciaAgenciaRisco     string `json:"ST_AG_RISCO,omitempty"`               // Indica se existe ou não agência de classificação de rating (1 caractere)
	TipoFundo                  string `json:"TP_FUNDO,omitempty"`                  // Tipo de fundo (15 caracteres)
	TributacaoFundoCotista     string `json:"TRIBUT_FUNDO_COTST,omitempty"`        // Descrição da tributação aplicável ao fundo e a seus cotistas, contemplando a política a ser adotada pelo administrador quanto ao tratamento tributário perseguido
	VotoGestor                 string `json:"VOTO_GESTOR,omitempty"`               // Indica se o gestor vota ou não em assembleias dos ativos que compõem a carteira (1 caractere)
}

type InformacoesDivulgacaoRequest struct {
	FundoCnpj        Cnpj   `json:"CNPJ_FUNDO,omitempty"` // CNPJ do fundo (20 caracteres)
	DataCompetencia  Date   `json:"DT_COMPTC,omitempty"`  // Data de competência do documento (AAAA-MM-DD)
	FormaInformacoes string `json:"FORMA_INF,omitempty"`  // Forma de divulgação das informações (100 caracteres)
	LocalInformacoes string `json:"LOCAL_INF,omitempty"`  // Local de divulgação das informações (300 caracteres)
	MeioInformacoes  string `json:"MEIO_INF,omitempty"`   // Meio de divulgação das informações (10 caracteres)
	TipoFundo        string `json:"TP_FUNDO,omitempty"`   // Tipo de fundo (15 caracteres)
}

type InformacoesCotistaResquest struct {
	FundoCnpj            Cnpj   `json:"CNPJ_FUNDO,omitempty"`        // CNPJ do fundo (20 caracteres)
	DescricaoRespCotista string `json:"DS_RESP_INF_COTST,omitempty"` // Descrição do responsável referente ao local, meio e forma de solicitação de informações pelo cotista (250 caracteres)
	DataCompetencia      Date   `json:"DT_COMPTC,omitempty"`         // Data de competência do documento (AAAA-MM-DD)
	FormaInfCotista      string `json:"FORMA_INF_COTST,omitempty"`   // Forma de solicitação de informações pelo cotista (100 caracteres)
	LocalInfCotista      string `json:"LOCAL_INF_COTST,omitempty"`   // Local de solicitação de informações pelo cotista (100 caracteres)
	MeioInfCotista       string `json:"MEIO_INF_COTST,omitempty"`    // Meio de solicitação de informações pelo cotista (10 caracteres)
	TipoFundo            string `json:"TP_FUNDO,omitempty"`          // Tipo de fundo (15 caracteres)
}

type ServicoPrestadoResquest struct {
	FundoCnpj        Cnpj   `json:"CNPJ_FUNDO,omitempty"`    // CNPJ do fundo (20 caracteres)
	DescricaoServico string `json:"DS_SERV_PREST,omitempty"` // Descrição do serviço prestado (100 caracteres)
	DataCompetencia  Date   `json:"DT_COMPTC,omitempty"`     // Data de competência do documento (AAAA-MM-DD)
	NomePrestador    string `json:"NM_PREST_SERV,omitempty"` // Nome do prestador do serviço (100 caracteres)
	TipoFundo        string `json:"TP_FUNDO,omitempty"`      // Tipo de fundo (15 caracteres)
}
