package configurations

import (
	"os"
	"log"
	"strconv"
)

type Config struct {
	Port     int
	Protocol string
}

func GetConfigurations() (config *Config) {
	rawPort, present := os.LookupEnv("PORT")
	if !present {
		log.Fatalf("Environment variable \"PORT\" is required but was not provided!\n")
	}
	port, err := strconv.Atoi(rawPort)
	if err != nil {
		log.Fatalf("Could not convert provided value for environment variable \"PORT\" into an int. \"Atoi\" error: \"%v\".\n", err.Error())
	}
	if port < 0 || port > 65535 {
		log.Fatalf("Value provided for environment variable \"PORT\" (\"%v\") is invalid. Must be greater than or equal to 0 and less than or equal to 65535.\n", port)
	}

	rawProtocol, present := os.LookupEnv("PROTOCOL")
	if !present {
		log.Fatalf("Environment variable \"PROTOCOL\" is required but was not provided!\n")
	}
	protocol := rawProtocol
	if protocol != "tcp" && protocol != "udp" {
		log.Fatalf("Value provided for environment variable \"PROTOCOL\" (\"%v\") is unsupported! Must be either \"tcp\" or \"udp\".\n", protocol)
	}

	return &Config{
		Port:     port,
		Protocol: protocol,
	}
}
