package promock

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-kit/log"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/prometheus/config"
	"github.com/prometheus/prometheus/storage"
	"github.com/prometheus/prometheus/storage/remote"
)

type Server struct {
	aliases []*Series

	remoteReadHandler http.Handler
}

func NewServer(logger log.Logger) *Server {
	s := &Server{}

	registry := prometheus.NewRegistry()

	config := func() config.Config {
		return config.Config{}
	}
	s.remoteReadHandler = remote.NewReadHandler(logger, registry, s, config, 10000, 10, 1048576)
	return s
}

func (s *Server) Serve(address string) {
	http.ListenAndServe(address, s.remoteReadHandler)
}

func (s *Server) ChunkQuerier(ctx context.Context, mint, maxt int64) (storage.ChunkQuerier, error) {
	return nil, fmt.Errorf("unimplemented")
}

func (s *Server) Querier(ctx context.Context, mint, maxt int64) (storage.Querier, error) {
	return nil, fmt.Errorf("unimplemented")
}
