package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	ServiceName string `mapstructure:"SERVICE_NAME"`

	PGDatabaseHost     string `mapstructure:"PG_DATABASE_HOST"`
	PGDatabaseUser     string `mapstructure:"PG_DATABASE_USER"`
	PGDatabasePassword string `mapstructure:"PG_DATABASE_PASSWORD"`
	PGDatabaseDBName   string `mapstructure:"PG_DATABASE_DBNAME"`
	PGDatabasePort     string `mapstructure:"PG_DATABASE_PORT"`

	GrpcAddr string `mapstructure:"GRPC_ADDR"`
	GrpcPort string `mapstructure:"GRPC_PORT"`

	KafkaTopic                  string `mapstructure:"KAFKA_TOPIC"`
	ResultsKafkaTopic           string `mapstructure:"RESULTS_KAFKA_TOPIC"`
	KafkaConsumerGroupId        string `mapstructure:"KAFKA_CONSUMER_GROUP_ID"`
	KafkaResultsConsumerGroupId string `mapstructure:"KAFKA_RESULTS_CONSUMER_GROUP_ID"`

	PathPrivateKey string `mapstructure:"PATH_PRIVATE_KEY"`
	PathPublicKey  string `mapstructure:"PATH_PUBLIC_KEY"`
}

func NewConfig(path string, name string) (config *Config, err error) {
	cfg := &Config{}

	viper.AddConfigPath(path)
	viper.SetConfigName(name)
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	if err = viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
		return nil, err
	}

	if err = viper.Unmarshal(&cfg); err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
		return nil, err
	}

	return cfg, err
}
