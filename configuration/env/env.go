package env

import "os"

func GetKafkaHost() string {
	return os.Getenv("KAFKA_HOST")
}

func GetTopicSincronizar() string {
	return "sincronizar"
}

func GetPathArquivosCvm() string {
	return "storage/cvm/"
}
