package app

import (
	"os"
	"qbit-exp/logger"
	"strconv"
)

const DEFAULT_ADDR = ":8090"
const DEFAULT_TIMEOUT = 30

type Env struct {
	Key          string
	DefaultValue string
	Help         string
}

var defaultLogLevel = Env{
	Key:          "LOG_LEVEL",
	DefaultValue: "INFO",
	Help:         "",
}

var defaultAddr = Env{
	Key:          "EXPORTER_ADDRESS",
	DefaultValue: DEFAULT_ADDR,
	Help:         "Exporter address is not set. Using default exporter address",
}

var defaultTimeout = Env{
	Key:          "DEFAULT_TIMEOUT",
	DefaultValue: strconv.Itoa(DEFAULT_TIMEOUT),
	Help:         "",
}

var defaultUsername = Env{
	Key:          "QBITTORRENT_USERNAME",
	DefaultValue: "admin",
	Help:         "Qbittorrent username is not set. Using default username",
}

var defaultPassword = Env{
	Key:          "QBITTORRENT_PASSWORD",
	DefaultValue: "adminadmin",
	Help:         "Qbittorrent password is not set. Using default password",
}

var defaultBaseUrl = Env{
	Key:          "QBITTORRENT_BASE_URL",
	DefaultValue: "http://localhost:8080",
	Help:         "Qbittorrent base_url is not set. Using default base_url",
}

var defaultDisableTracker = Env{
	Key:          "DISABLE_TRACKER",
	DefaultValue: "false",
	Help:         "",
}

var defaultAuthPassword = Env{
	Key:          "AUTH_PASSWORD",
	DefaultValue: "",
	Help:         "Basic Auth password for access to this instance",
}

func getEnv(env Env) string {
	value, ok := os.LookupEnv(env.Key)
	if !ok || value == "" {
		if env.Help != "" {
			logger.Log.Warn(env.Help)
		}
		return env.DefaultValue
	}
	return value
}
