package main

import (
	"fmt"
	"net"
	"net/http"
	"runtime"

	"qbit-exp/qbit"

	app "qbit-exp/app"
	logger "qbit-exp/logger"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	Version     = "dev"
	Author      = "martabal"
	ProjectName = "qbittorrent-exporter"
)

func basicAuth(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, pass, ok := r.BasicAuth()
		if app.AuthToken != "" && (!ok || pass != app.AuthToken) {
			w.Header().Set("WWW-Authenticate", `Basic realm="restricted"`)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		handler(w, r)
	}
}

func main() {
	app.LoadEnv()
	fmt.Printf("%s (version %s)\n", ProjectName, Version)
	fmt.Println("Author:", Author)
	fmt.Println("Using log level: " + fmt.Sprintf("%s%s%s", logger.ColorLogLevel[logger.LogLevels[app.LogLevel]], app.LogLevel, logger.Reset))

	logger.Log.Info("qbittorrent URL: " + app.BaseUrl)
	logger.Log.Info("username: " + app.Username)
	logger.Log.Info("password: " + app.GetPasswordMasked())
	logger.Log.Info("Started")
	isTrackerEnabled := "enabled"
	if app.DisableTracker {
		isTrackerEnabled = "disabled"
	}
	logger.Log.Debug("Trackers info is " + isTrackerEnabled)

	qbit.Auth()

	http.HandleFunc("/metrics", basicAuth(metrics))
	if app.Addr != app.DEFAULT_ADDR {
		logger.Log.Info("Listening on port " + app.Addr)
	}
	logger.Log.Info("Starting the exporter")
	err := http.ListenAndServe(app.Addr, nil)
	if err != nil {
		panic(err)
	}
}

func metrics(w http.ResponseWriter, req *http.Request) {
	ip, _, err := net.SplitHostPort(req.RemoteAddr)
	if err == nil {
		logger.Log.Trace("New request from " + ip)
	} else {
		logger.Log.Trace("New request")
	}

	registry := prometheus.NewRegistry()
	err = qbit.AllRequests(registry)
	if err != nil {
		http.Error(w, "", http.StatusServiceUnavailable)
		runtime.GC()
	} else {
		h := promhttp.HandlerFor(registry, promhttp.HandlerOpts{})
		h.ServeHTTP(w, req)
	}

}
