package config

import "github.com/spf13/viper"

// HttpConfig contains env config
type HttpConfig struct {
	Host string `env:"HOST"`
	Port int    `env:"PORT"`
}

// LoadHttpConfig loads http config from env vars
func LoadHttpConfig() (*HttpConfig, error) {
	viper.AutomaticEnv()

	httpConfig := &HttpConfig{}
	err := viper.Unmarshal(httpConfig)
	if err != nil {
		return nil, err
	}
	return httpConfig, nil
}
