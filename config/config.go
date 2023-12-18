package config

import "flag"

type Config struct {
	Port       string
	LogLevel   string
	MaxClients int
}

const DefaultPort = "8080"

func ReadConfig() *Config {

	port := flag.String("port", DefaultPort, "The server port")
	logLevel := flag.String("logLevel", "info", "The server log level")
	maxClients := flag.Int("maxClients", 100, "The server max clients")

	flag.Parse()

	ConfigureLogging(*logLevel)

	return &Config{
		Port:       *port,
		LogLevel:   *logLevel,
		MaxClients: *maxClients,
	}
}
