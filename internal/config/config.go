package config

import (
	"flag"
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
	"time"
)

type Config struct {
	Env         string        `yaml:"env" env-default:"local"`
	StoragePath string        `yaml:"storage_path" env-required:"true"`
	TokenTTL    time.Duration `yaml:"token_ttl" env-required:"true"`
	GRPC        GRPCConfig    `yaml:"grpc"`
}

type GRPCConfig struct {
	Port    int           `yaml:"port"`
	Timeout time.Duration `yaml:"timeout"`
}

func (c *Config) Log() {
	log.Printf("%+v", c)
}
func MustLoad() *Config {
	path := fetchConfigPath()
	if path == "" {
		panic("config path is empty")
	}
	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic("config file does not exist: " + path)
	}

	var cfg Config
	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		panic("failed to read config: " + err.Error())
	}

	return &cfg
}

// Priority flag > env > default
func fetchConfigPath() string {
	var res string
	// --config="path/to/config.yaml"
	flag.StringVar(&res, "config", "", "path to config file")
	flag.Parse()

	if res == "" {
		res = os.Getenv("CONFIG_PATH")
	}

	return res
}

//type Config struct {
//	Logging logging
//	Storage storage
//	Service service
//}

//type logging struct {
//	Level string `env:"LOG_LEVEL"`
//}
//type storage struct {
//	DBPath string `env:"DB_PATH"`
//}
//type service struct {
//	AddressGRPC string `env:"ADDRESS_GRPC"`
//	TimeoutGRPC uint8  `env:"TIMEOUT_GRPC"`
//	TokenTTL    uint8  `env:"TOKEN_TTL"`
//}
//
//func (c *Config) Log() {
//	log.Printf("%+v", c)
//}
//
//func New() (*Config, error) {
//	cfg := &Config{}
//	if err := cleanenv.ReadEnv(cfg); err != nil {
//		return nil, err
//	}
//	return cfg, nil
//}
