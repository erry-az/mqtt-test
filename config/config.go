package config

import (
	"errors"
	"os"
	"strings"
)

const mqttHostEnv = "MQTT_HOST"

func GetMQTTHost() (string, error) {
	host := os.Getenv(mqttHostEnv)
	if len(strings.TrimSpace(host)) < 1 {
		return "", errors.New("host not found please set env var for " + mqttHostEnv)
	}

	return host, nil
}
