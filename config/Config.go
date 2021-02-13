package config

type Config struct {
	Server struct {
		BindAddress string `yaml:"bindAddress" envconfig:"SERVER__BIND_ADDRESS"`
		Cors        struct {
			AllowOrigins []string `yaml:"allowOrigins" envconfig:"SERVER__CORS__ALLOW_ORIGINS`
		} `yaml:"cors"`
	} `yaml:"server"`
	Database struct {
		Driver  string `yaml:"driver" envconfig:"DATABASE__DRIVER"`
		Address string `yaml:"address" envconfig:"DATABASE__ADDRESS"`
		Debug   bool   `yaml:"debug" envconfig:"DATABASE__DEBUG"`
	} `yaml:"database"`
}
