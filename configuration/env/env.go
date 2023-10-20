package env

import "os"

func GetKafkaHost() string {
	return os.Getenv("KAFKA_HOST")
}

func GetKafkaTopicSincronizar() string {
	return "sincronizar"
}
