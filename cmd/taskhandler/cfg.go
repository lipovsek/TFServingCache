package main

import (
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func SetConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")
	setDefaults()
	viper.SetEnvPrefix("tfsc")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Info("No config file found. Reading from env vars")
		} else {
			log.WithError(err).Panic("Could not read config file")
		}
	}

	logLevel := viper.GetString("logging.level")
	logFormat := viper.GetString("logging.format")
	switch logFormat {
	case "json":
		log.SetFormatter(&log.JSONFormatter{})
		break
	case "text":
		log.SetFormatter(&log.TextFormatter{})
		break
	default:
		log.SetFormatter(&log.TextFormatter{})
	}

	log.Infof("Log Level: %v", logLevel)

	// Set log level
	switch logLevel {
	case "panic":
		log.SetLevel(log.PanicLevel)
		break
	case "fatal":
		log.SetLevel(log.FatalLevel)
		break
	case "warning":
		log.SetLevel(log.WarnLevel)
		break
	case "debug":
		log.SetLevel(log.DebugLevel)
		break
	case "info":
	default:
		log.SetLevel(log.InfoLevel)
	}

}

func setDefaults() {
	viper.SetDefault("healthprobe.modelName", "__TFSERVINGCACHE_PROBE_CHECK__")
}
