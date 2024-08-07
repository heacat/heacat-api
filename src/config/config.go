package config

import (
	"flag"
	"os"

	"github.com/heacat/heacat-api/src/logger"
	"github.com/spf13/viper"
)

var Config Config_t

type Config_t struct {
	Alarm struct {
		ServerNickName string
		Telegram       struct {
			Enabled         bool
			Token           string
			ChatID          string
			MessageThreadID int
		}
		Slack struct {
			Enabled    bool
			WebHookURL string
		}
	}
	API struct {
		Host string
		Port int
	}
	Disk struct {
		FileSystems   []string
		PartUseLimit  int
		CheckInterval int
		Unit          string
	}
	Cpu struct {
		LoadLimit     float32
		CheckInterval int
	}
	Memory struct {
		UseLimit      int
		CheckInterval int
		Unit          string
	}
	Network struct {
		Interfaces    []string
		CheckInterval int
	}
}

func InitConfig() {
	config_file := "config.yaml"

	version := flag.Bool("version", false, "Print version information")
	config_path := flag.String("config", "config.yaml", "Path to the config file")
	flag.Parse()

	if *version {
		logger.Log.Infof("Version: 1.0.0")
		os.Exit(0)
	}

	if *config_path != "config.yaml" {
		config_file = *config_path
	}

	if _, err := os.Stat(config_file); os.IsNotExist(err) {
		logger.Log.Fatalf("Configuration file: %s does not exist, %v\n", config_file, err)
	}

	viper.SetConfigFile(config_file)
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		logger.Log.Fatalf("Error reading config file, %s\n", err)
	}

	err := viper.Unmarshal(&Config)
	if err != nil {
		logger.Log.Fatalf("Unable to decode into struct, %v\n", err)
	}
	logger.Log.Infof("Config file loaded successfully: %s\n", config_file)
}
