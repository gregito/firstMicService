package homepage

import (
	"log"
	"net/http"
	"time"
)

const helloMessage = "Welcome"

type Handlers struct {
	logger *log.Logger
}

func (h *Handlers) Home(writer http.ResponseWriter, request *http.Request) {
	h.logger.Printf("request coming from: %s\n", request.Host)
	writer.Header().Set("Content-Type", "text/plain; charset=utf-8")
	writer.WriteHeader(http.StatusOK)
	writer.Write([]byte(helloMessage))
}

func (h *Handlers) Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		defer h.logger.Printf("requet processed in %s\n", time.Now().Sub(startTime))
		next(w, r)
	}
}
func (h *Handlers) SetupRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", h.Logger(h.Home))
}

func NewHandlers(loggerImpl *log.Logger) *Handlers {
	return &Handlers{
		logger: loggerImpl,
	}
}
