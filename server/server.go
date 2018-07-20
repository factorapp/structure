package server

import (
	"net/http"
	"strconv"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/packr"
)

type Server struct {
	hdl http.Handler
}

func (s *Server) Start(port int) error {
	return http.ListenAndServe(":"+strconv.Itoa(port), s.hdl)
}

func NewServer(templates packr.Box) http.Handler {
	app := buffalo.New(buffalo.Options{})
	app.GET("/wasm_exec.js", wasmJS)
	handler := mux.NewRouter()
	handler.HandleFunc
	src := http.Server{}
	// TODO: this thing needs to serve the wasm app, the wasm.js
	return &Server{hdl: app}
}
