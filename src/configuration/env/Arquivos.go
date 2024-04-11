package env

import (
	"time"
)

type ArquivosCVM struct {
	Folder      string
	Extension   string
	DataInicial string
	DataFinal   string
	Formato     string
	Outro       string
	Tipo        string
	Ano         int
	Mes         int
	Dia         int
}

func GetConfigCvmBalancete() ArquivosCVM {
	return ArquivosCVM{
		Folder:      "FI/DOC/BALANCETE/DADOS/balancete_fi_",
		Extension:   ".zip",
		DataInicial: "2015-01-01",
		DataFinal:   time.Now().AddDate(0, -3, 0).Format("2006-01-02"),
		Formato:     "200601",
		Ano:         0,
		Mes:         1,
		Dia:         0,
		Tipo:        "balancete",
	}
}

func GetConfigCvmBalanceteHist() ArquivosCVM {
	return ArquivosCVM{
		Folder:      "FI/DOC/BALANCETE/DADOS/HIST/balancete_fi_",
		Extension:   ".zip",
		DataInicial: "2005-05-01",
		DataFinal:   "2014-12-01",
		Formato:     "200601",
		Ano:         0,
		Mes:         1,
		Dia:         0,
		Tipo:        "balancete",
	}
}

func GetConfigCvmCda() ArquivosCVM {
	return ArquivosCVM{
		Folder:      "FI/DOC/CDA/DADOS/cda_fi_",
		Extension:   ".zip",
		DataInicial: "2023-01-01",
		DataFinal:   time.Now().AddDate(0, -3, 0).Format("2006-01-02"),
		Formato:     "200601",
		Ano:         0,
		Mes:         1,
		Dia:         0,
		Tipo:        "cda",
	}
}

func GetConfigCvmCdaHist() ArquivosCVM {
	return ArquivosCVM{
		Folder:      "FI/DOC/CDA/DADOS/HIST/cda_fi_",
		Extension:   ".zip",
		DataInicial: "2005-01-01",
		DataFinal:   "2021-12-01",
		Formato:     "2006",
		Ano:         1,
		Mes:         0,
		Dia:         0,
		Tipo:        "cda",
	}
}

func GetMapCda() map[string]string {
	return map[string]string{
		"BLC_1":       "cda-blc-1",
		"BLC_2":       "cda-blc-2",
		"BLC_3":       "cda-blc-3",
		"BLC_4":       "cda-blc-4",
		"BLC_5":       "cda-blc-5",
		"BLC_6":       "cda-blc-6",
		"BLC_7":       "cda-blc-7",
		"BLC_8":       "cda-blc-8",
		"fi_CONFID_":  "cda-confidencial",
		"fiim":        "cda-fiim",
		"fiim_CONFID": "cda-fiim-confidencial",
		"fi_PL":       "cda-patrimonio-liquido",
	}
}

func GetMapInformacaoComplementar() map[string]string {
	return map[string]string{
		"compl_fi_20":           "informacao-complementar",
		"compl_fi_inf_20":       "informacao-complementar-divulgacao",
		"compl_fi_inf_cotst_20": "informacao-complementar-cotista",
		"compl_fi_prest_20":     "informacao-complementar-servico-prestado",
	}
}

func GetConfigCvmArquivosPerfilMensal() ArquivosCVM {
	return ArquivosCVM{
		Folder:      "FI/DOC/PERFIL_MENSAL/DADOS/perfil_mensal_fi",
		Extension:   ".csv",
		DataInicial: "2019-01-01",
		DataFinal:   time.Now().AddDate(0, -2, 0).Format("2006-01-02"),
		Formato:     "200601",
		Ano:         0,
		Mes:         1,
		Dia:         0,
		Tipo:        "perfil-mensal",
	}
}

func GetConfigCvmInformacoesComplementares() ArquivosCVM {
	return ArquivosCVM{
		Folder:      "FI/DOC/COMPL/DADOS/compl_fi_",
		Extension:   ".zip",
		DataInicial: "2018-01-01",
		DataFinal:   "2019-04-01",
		Formato:     "200601",
		Ano:         0,
		Mes:         1,
		Dia:         0,
		Tipo:        "informacoes-complementares",
	}
}

func GetConfigCvmExtrato() ArquivosCVM {
	return ArquivosCVM{
		Outro:       "FI/DOC/EXTRATO/DADOS/extrato_fi.csv",
		Folder:      "FI/DOC/EXTRATO/DADOS/extrato_fi_",
		Extension:   ".csv",
		DataInicial: "2015-01-01",
		DataFinal:   "2022-12-31",
		Formato:     "2006",
		Ano:         1,
		Mes:         0,
		Dia:         0,
		Tipo:        "extrato",
	}
}

func GetConfigInformacaoDiaria() ArquivosCVM {
	return ArquivosCVM{
		Folder:      "FI/DOC/INF_DIARIO/DADOS/inf_diario_fi_",
		Extension:   ".zip",
		DataInicial: "2021-01-01",
		DataFinal:   time.Now().AddDate(0, -3, 0).Format("2006-01-02"),
		Formato:     "200601",
		Ano:         0,
		Mes:         1,
		Dia:         0,
		Tipo:        "informacao-diaria",
	}
}

func GetConfigInformacaoDiariaHist() ArquivosCVM {
	return ArquivosCVM{
		Folder:      "FI/DOC/INF_DIARIO/DADOS/HIST/inf_diario_fi_",
		Extension:   ".zip",
		DataInicial: "2000-01-01",
		DataFinal:   "2019-12-01",
		Formato:     "2006",
		Ano:         1,
		Mes:         0,
		Dia:         0,
		Tipo:        "informacao-diaria",
	}
}

func GetConfigCvmArquivosLaminas() ArquivosCVM {
	return ArquivosCVM{
		Folder:      "FI/DOC/LAMINA/DADOS/lamina_fi_",
		Extension:   ".zip",
		DataInicial: "2019-01-01",
		DataFinal:   time.Now().AddDate(0, -2, 0).Format("2006-01-02"),
		Formato:     "200601",
		Ano:         0,
		Mes:         1,
		Dia:         0,
		Tipo:        "lamina",
	}
}

func GetConfigCvmArquivosLaminasHist() ArquivosCVM {
	return ArquivosCVM{
		Folder:      "FI/DOC/LAMINA/DADOS/HIST/lamina_fi_",
		Extension:   ".zip",
		DataInicial: "2014-01-01",
		DataFinal:   "2018-11-01",
		Formato:     "200601",
		Ano:         0,
		Mes:         1,
		Dia:         0,
		Tipo:        "lamina",
	}
}

func GetConfigCvmArquivosCadastros() []string {
	return []string{
		"FI/CAD/DADOS/cad_fi.csv",
		"FI/CAD/DADOS/cad_fi_hist.zip",
		"cadastros",
	}
}
