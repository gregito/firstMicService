package main

import (
	"log"
	"net/http"
	"os"

	"firstMicS/server"
)

const helloMessage = "Welcome"

var (
	CertFile       = os.Getenv("SERVER_CERT_FILE")
	KeyFile        = os.Getenv("SERVER_KEY_FILE")
	ServiceAddress = os.Getenv("SERVICE_ADDRESS")
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "text/plain; charset=utf-8")
		writer.WriteHeader(http.StatusOK)
		writer.Write([]byte(helloMessage))
	})

	srv := server.New(server.NewTlsConfig(), mux, ServiceAddress)

	if err := srv.ListenAndServeTLS(CertFile, KeyFile); err != nil {
		log.Fatalf("server failed to start <%v>", err)
	}
}
