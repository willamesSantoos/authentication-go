package configs

import "github.com/spf13/viper"

var cfg *config

type config struct {
	Server ServerConfig
	DB     DBConfig
}

type ServerConfig struct {
	Port string `json:"port"`
}

type DBConfig struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Database string `json:"database"`
}

func init() {
	viper.SetDefault("server.port", "8000")
	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("database.port", "5432")
}

func LoadConfig() error {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()

	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return err
		}
	}

	cfg = new(config)

	cfg.Server = ServerConfig{
		Port: viper.GetString("server.port"),
	}

	cfg.DB = DBConfig{
		Host:     viper.GetString("database.host"),
		Port:     viper.GetString("database.port"),
		User:     viper.GetString("database.user"),
		Password: viper.GetString("database.pass"),
		Database: viper.GetString("database.database"),
	}

	return nil
}

func GetDBConfig() DBConfig {
	return cfg.DB
}

func GetServerPort() string {
	return cfg.Server.Port
}
