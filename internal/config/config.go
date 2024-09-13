package config

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	TokenTTL  time.Duration `yaml:"token_ttl" env:"TOKEN_TTL" env-default:"1h"`
	SecretKey string        `yaml:"secret_key" env:"SECRET_KEY" env-required:"true"`
	Address   string        `yaml:"address" env-default:":4000"`
	Port      string        `yaml:"port" env:"PORT" env-default:"5432"`
	Host      string        `yaml:"host" env:"HOST" env-default:"localhost"`
	Name      string        `yaml:"name" env:"NAME" env-default:"postgres"`
	DBName    string        `yaml:"db_name" env:"DB_NAME" env-default:"users"`
	User      string        `yaml:"user" env:"USER" env-default:"user"`
	Password  string        `yaml:"password" env:"PASSWORD" env-required:"true"`
	SSLMode   string        `yaml:"sslmode" env:"SSLMODE" env-default:"disable"`
}

func NewConfig() *Config {
	viper.AddConfigPath("config")
	viper.SetConfigName("dev")
	viper.ReadInConfig()

	if err := godotenv.Load(); err != nil {
		log.Fatalf("ошибка загрузки переменных окружения: %s", err.Error())
	}

	return &Config{
		TokenTTL:  viper.GetDuration("token_ttl"),
		SecretKey: os.Getenv("SECRET_KEY"),
		Address:   viper.GetString("address"),
		Port:      viper.GetString("port"),
		Host:      viper.GetString("host"),
		Name:      viper.GetString("name"),
		DBName:    viper.GetString("db_name"),
		User:      viper.GetString("user"),
		Password:  os.Getenv("PASSWORD"),
		SSLMode:   viper.GetString("sslmode"),
	}
}
