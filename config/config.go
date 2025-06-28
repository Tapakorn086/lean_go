package config

type Config struct {
	AppName string `mapstructure:"APP_NAME"`
	Port    int    `mapstructure:"PORT"`
	DBURI   string   `mapstructure:"DBURI"`
}
