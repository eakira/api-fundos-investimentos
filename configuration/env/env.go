package env

import "os"

func GetKafkaHost() string {
	return os.Getenv("KAFKA_HOST")
}

func GetTopicSincronizar() string {
	return "sincronizar"
}

func GetTopicProcessarArquivos() string {
	return "processar-arquivos"
}

func GetTopicPersistenciaDados() string {
	return "persistencia-dados"
}

func GetPathArquivosCvm() string {
	return "storage/cvm/"
}

func GetLimitInsert() int {
	return 1000
}
