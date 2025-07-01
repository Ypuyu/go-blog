package util

import (
	"os"
	"strings"
)

func GetServerName() string {
	return strings.ToLower(os.Getenv("server"))
}

func GetEnv() string {
	return strings.ToLower(os.Getenv("ENV"))
}

func IsLocal() bool {
	return strings.ToLower(os.Getenv("ENV")) == "local"
}

func IsTest() bool {
	return strings.ToLower(os.Getenv("ENV")) == "test"
}

func IsDev() bool {
	return strings.ToLower(os.Getenv("ENV")) == "dev"
}

func IsLive() bool {
	return strings.ToLower(os.Getenv("ENV")) == "live" ||
		strings.ToLower(os.Getenv("ENV")) == "ecs_live"
}
