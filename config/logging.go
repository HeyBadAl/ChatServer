package config

import "log"

func ConfigureLogging(logLevel string) {
	log.SetPrefix("[ChatServer] ")
	switch logLevel {
	case "debug":
		log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Lshortfile)
	case "info":
		log.SetFlags(log.LstdFlags)
	default:
		log.Fatal("Invalid log level: ", logLevel)
	}
}
