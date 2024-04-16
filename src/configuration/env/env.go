package env

import (
	"os"
	"strconv"
)

func GetPort() string {
	return os.Getenv("PORT")
}

func GetAppName() string {
	return os.Getenv("APP_NAME")
}

func GetAmbienteDev() bool {
	return os.Getenv("ENV") == "local"
}

func GetTimeShutdown() string {
	return os.Getenv("TIME_SHUTDOWN")
}

func GetMongoUrl() string {
	return os.Getenv("MONGODB_URL")
}

func GetMongoDataBase() string {
	return os.Getenv("MONGODB_DATABASE")
}

// Se a persistencia vai ser executado na hora ou se vai para fila
func GetPersistenciaLocal() bool {
	return os.Getenv("PERSISTENCIA") == "local"
}

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
func GetCvmUrl() string {
	return os.Getenv("CVM_URL")
}

func GetPathArquivosCvm() string {
	return os.Getenv("STORAGE")
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

func GetLogOutup() string {
	return os.Getenv("LOG_OUTPUT")
}

func GetLogInfoLevel() string {
	return os.Getenv("LOG_INFO_PATH")
}

func GetLogErrorPath() string {
	return os.Getenv("LOG_ERROR_PATH")
}

func GetLogLevel() string {
	return os.Getenv("LOG_LEVEL")
}
