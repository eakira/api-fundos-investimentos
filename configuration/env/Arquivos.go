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
	tipo        string
	Ano         int
	Mes         int
	Dia         int
}

func GetConfigCvmBalancete() ArquivosCVM {
	return ArquivosCVM{
		Folder:      "FI/DOC/BALANCETE/DADOS/balancete_fi_",
		Extension:   ".zip",
		DataInicial: "2015-01-01",
		DataFinal:   time.Now().AddDate(0, -2, 0).Format("2006-01-02"),
		Formato:     "200601",
		Ano:         0,
		Mes:         1,
		Dia:         0,
		tipo:        "balancete",
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
		tipo:        "balancete",
	}
}

func GetConfigCvmCda() ArquivosCVM {
	return ArquivosCVM{
		Folder:      "FI/DOC/CDA/DADOS/cda_fi_",
		Extension:   ".zip",
		DataInicial: "2021-01-01",
		DataFinal:   time.Now().AddDate(0, -2, 0).Format("2006-01-02"),
		Formato:     "200601",
		Ano:         0,
		Mes:         1,
		Dia:         0,
		tipo:        "cda",
	}
}

func GetConfigCvmCdaHist() ArquivosCVM {
	return ArquivosCVM{
		Folder:      "FI/DOC/CDA/DADOS/HIST/cda_fi_",
		Extension:   ".zip",
		DataInicial: "2005-01-01",
		DataFinal:   "2020-12-01",
		Formato:     "2006",
		Ano:         1,
		Mes:         0,
		Dia:         0,
		tipo:        "cda",
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
		tipo:        "perfil-mensal",
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
		tipo:        "informacoes-complementares",
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
		tipo:        "extrato",
	}
}

func GetConfigInformacaoDiaria() ArquivosCVM {
	return ArquivosCVM{
		Folder:      "FI/DOC/INF_DIARIO/DADOS/inf_diario_fi_",
		Extension:   ".zip",
		DataInicial: "2021-01-01",
		DataFinal:   time.Now().AddDate(0, -1, 0).Format("2006-01-02"),
		Formato:     "200601",
		Ano:         0,
		Mes:         1,
		Dia:         0,
		tipo:        "informacao-diaria",
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
		tipo:        "informacao-diaria",
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
		tipo:        "lamina",
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
		tipo:        "lamina",
	}
}

func GetConfigCvmArquivosCadastros() []string {
	return []string{
		"FI/CAD/DADOS/cad_fi.csv",
		"FI/CAD/DADOS/cad_fi_hist.zip",
		"cadastros",
	}
}
