package server

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type httpServer struct {
	Log *Log
}

func newHTTPServer() *httpServer {
	return &httpServer{
		Log: NewLog(),
	}
}

type ProduceRequest struct {
	Record `json:"record"`
}

type ProduceResponse struct {
	Offset uint64 `json:"offset"`
}

type ConsumeRequest struct {
	Offset uint64 `json:"offset"`
}

type ConsumeResponse struct {
	Record `json:"record"`
}

func NewHTTPServer(addr string) *http.Server {
	srv := newHTTPServer()
	r := mux.NewRouter()

	r.HandleFunc("/", srv.handleProduce).Methods("POST")
	r.HandleFunc("/", srv.handleConsume).Methods("GET")

	// From here user just need to call ListenAndServe()
	return &http.Server{
		Addr:    addr,
		Handler: r,
	}
}

func (s *httpServer) handleProduce(w http.ResponseWriter, r *http.Request) {
	var req ProduceRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w http.ResponseWriter, error string, code int)
	}
}
