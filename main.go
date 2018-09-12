package main

import (
	"firstMicS/homepage"
	"log"
	"net/http"
	"os"
	"time"

	"firstMicS/server"
)

var (
	CertFile       = os.Getenv("SERVER_CERT_FILE")
	KeyFile        = os.Getenv("SERVER_KEY_FILE")
	ServiceAddress = os.Getenv("SERVICE_ADDRESS")
)

func main() {
	logger := log.New(os.Stdout, getTimeStamp(), log.LstdFlags|log.Lshortfile)

	h := homepage.NewHandlers(logger)

	mux := http.NewServeMux()
	h.SetupRoutes(mux)

	srv := server.New(server.NewTlsConfig(), mux, ServiceAddress)

	logger.Println("Server starting...")
	if err := srv.ListenAndServeTLS(CertFile, KeyFile); err != nil {
		logger.Fatalf("server failed to start <%v>", err)
	}
}

func getTimeStamp() string {
	return time.Unix(unixMilli(time.Now())/1e3, (unixMilli(time.Now())%1e3)*int64(time.Millisecond)/int64(time.Nanosecond)).String()
}

func unixMilli(t time.Time) int64 {
	return t.Round(time.Millisecond).UnixNano() / (int64(time.Millisecond) / int64(time.Nanosecond))
}
