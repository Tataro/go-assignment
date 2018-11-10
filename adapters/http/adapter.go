package http

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gitlab.com/upaphong/go-assignment/engine"
)

type HTTPAdapter struct {
	srv *http.Server
}

func (adapter *HTTPAdapter) Start() {
	go func() {
		log.Println("server start on port 3000")
		log.Fatal(adapter.srv.ListenAndServe())
	}()
}

func (adapter *HTTPAdapter) Stop() {
	if err := adapter.srv.Shutdown(nil); err != nil {
		panic(err)
	}
}

func NewHTTPAdapter(e engine.Engine) *HTTPAdapter {
	r := mux.NewRouter()
	RegisterRoutes(e, r)

	srv := &http.Server{
		Handler: r,
		Addr:    ":3000",
	}
	return &HTTPAdapter{srv: srv}
}
