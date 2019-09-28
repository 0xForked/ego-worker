package helper

import "strconv"

type Config struct {
	App          *AppConfig
	Service      *ServiceConfig
	MySQL        *MySQlConfig
	RabbitMQ     *RabbitMQConfig
	EmailDefault *EmailDefaultConfig
}

type AppConfig struct {
	Name  string
	Debug bool
}

type ServiceConfig struct {
	Delivery string
}

type MySQlConfig struct {
	Driver string
	URL    string
}

type RabbitMQConfig struct {
	URL        string
	Topic      string
	Durable    bool
	AutoDelete bool
	Exclusive  bool
	NoWait     bool
}

type EmailDefaultConfig struct {
	Host string
	Port string
	User string
	Pass string
}

func GetConfig() *Config {
	appDebug, _ :=
		strconv.ParseBool(
			GetEnv("APP_DEBUG", "false"))
	queueDurable, _ :=
		strconv.ParseBool(
			GetEnv("MQ_QUEUE_DURABLE", "false"))
	queueAutoDelete, _ :=
		strconv.ParseBool(
			GetEnv("MQ_QUEUE_AUTO_DELETE", "false"))
	queueExclusive, _ :=
		strconv.ParseBool(
			GetEnv("MQ_QUEUE_EXCLUSIVE", "false"))
	queueNoWait, _ :=
		strconv.ParseBool(
			GetEnv("MQ_QUEUE_NO_WAIT", "false"))

	return &Config{
		App: &AppConfig{
			Name:  GetEnv("APP_NAME", "mego_worker"),
			Debug: appDebug,
		},
		Service: &ServiceConfig{
			Delivery: GetEnv("SERVICE_DELIVERY", "schema"),
		},
		MySQL: &MySQlConfig{
			Driver: GetEnv("DB_DRIVER", "mysql"),
			URL:    GetEnv("DB_CONNECTION_URL", ""),
		},
		RabbitMQ: &RabbitMQConfig{
			URL:        GetEnv("MQ_CONNECTION_URL", ""),
			Topic:      GetEnv("MQ_QUEUE_TOPIC", "ego_worker"),
			Durable:    queueDurable,
			AutoDelete: queueAutoDelete,
			Exclusive:  queueExclusive,
			NoWait:     queueNoWait,
		},
		EmailDefault: &EmailDefaultConfig{
			Host: GetEnv("MAIL_HOST", "smtp.example.com"),
			Port: GetEnv("MAIL_PORT", "587"),
			User: GetEnv("MAIL_USER", "user@example.com"),
			Pass: GetEnv("MAIL_PASS", "password"),
		},
	}
}
