package env

import (
	"log"

	"github.com/spf13/viper"
)

type Env struct {
	Kafka kafka `mapstructur"e="KAFKA"`
	Mail  mail  `mapstructure="MAIL"`
}

type kafka struct {
	Url   string `mapstructure="URL"`
	Topic string `mapstructure="TOPIC"`
}
type mail struct {
	From     string   `mapstructure="FROM"`
	To       []string `mapstructure="TO"`
	Password string   `mapstructure="PASSWORD"`
	Smtphost string   `mapstructure="SMTPHOST"`
	Smtpport string   `mapstructure="SMTPPORT"`
}

func LoadEnv(path string) *Env {
	viper.AddConfigPath(path)
	viper.SetConfigName("dev")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()

	if err != nil {
		log.Fatal("Error opening config file")
	}

	conf := &Env{}
	err = viper.Unmarshal(&conf)

	if err != nil {
		log.Fatal("Error unmarshal config file")
	}
	return conf

}
