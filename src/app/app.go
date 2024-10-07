package app

import (
	"fmt"
	"os"
	"qbit-exp/logger"
	"strconv"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

var (
	QBittorrentTimeout time.Duration
	Addr               string
	ShouldShowError    bool = true
	DisableTracker     bool
	LogLevel           string
	BaseUrl            string
	Cookie             string
	Username           string
	Password           string
	AuthPassword       string
)

func LoadEnv() {
	_, err := os.Stat(".env")
	if err == nil || !os.IsNotExist(err) {
		err := godotenv.Load(".env")
		if err != nil {
			errormessage := "Error loading .env file:" + err.Error()
			panic(errormessage)
		}
	}

	LogLevel = logger.SetLogLevel(getEnv(defaultLogLevel))

	timeoutDuration, errTimeoutDuration := strconv.Atoi(getEnv(defaultTimeout))
	if errTimeoutDuration != nil {
		panic(fmt.Sprintf("%s must be an integer", defaultTimeout.Key))
	}
	if timeoutDuration < 0 {
		panic(fmt.Sprintf("%s must be > 0", defaultTimeout.Key))
	}

	Addr = getEnv(defaultAddr)
	DisableTracker = strings.ToLower(getEnv(defaultDisableTracker)) == "true"
	BaseUrl = strings.TrimSuffix(getEnv(defaultBaseUrl), "/")
	Username = getEnv(defaultUsername)
	Password = getEnv(defaultPassword)
	QBittorrentTimeout = time.Duration(timeoutDuration)
	AuthPassword = getEnv(defaultAuthPassword)
}

func GetPasswordMasked() string {
	return strings.Repeat("*", len(Password))
}
