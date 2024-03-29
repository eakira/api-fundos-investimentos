package env

import (
	"os"
	"strconv"
)

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
	return 200
}

func GetNumParticoes() int32 {

	value := os.Getenv("NUM_PARTITIONS")
	v, _ := strconv.ParseInt(value, 10, 32)
	return int32(v)
}

func GetTopics() string {
	return os.Getenv("TOPICS")
}

func GetPartitions() string {
	return os.Getenv("PARTITIONS")
}
