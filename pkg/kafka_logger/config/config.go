package config

type KafkaConfig struct {
	Brokers []string
	Topic   string
}

type RedisConfig struct {
	Addr     string
	Password string
	DB       int
}

func LoadKafkaConfig() KafkaConfig {
	return KafkaConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "logs",
	}
}

func LoadRedisConfig() RedisConfig {
	return RedisConfig{
		Addr:     "localhost:6379",
		Password: "test_password",
		DB:       0,
	}
}
