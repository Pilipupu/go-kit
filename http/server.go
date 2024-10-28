package http

import (
	"fmt"
	"github.com/Pilipupu/go-kit/framwork/component"
	"github.com/Pilipupu/go-kit/log"
	"net/http"
)

var logger = log.New(log.NewDefaultFileWriter(), log.NewDefaultOptions())

var components []component.Component

func NewHttpServer(components []component.Component) error {
	server := http.Server{Addr: "0.0.0.0:8080", Handler: http.DefaultServeMux}
	err := server.ListenAndServe()
	if err != nil {
		return fmt.Errorf("http server error: %w", err)
	}

	return nil
}
