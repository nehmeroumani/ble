package gatt

import (
	"sync"

	"github.com/nehmeroumani/ble"
	"github.com/nehmeroumani/ble/linux/att"
)

// NewServer
func NewServer(defaultServices ...*ble.Service) (*Server, error) {
	return &Server{
		svcs: defaultServices,
		db:   att.NewDB(defaultServices, uint16(1)),
	}, nil
}

// Server ...
type Server struct {
	sync.Mutex
	svcs []*ble.Service
	db   *att.DB
}

// AddService ...
func (s *Server) AddService(svc *ble.Service) error {
	s.Lock()
	defer s.Unlock()
	s.svcs = append(s.svcs, svc)
	s.db = att.NewDB(s.svcs, uint16(1)) // ble attrs start at 1
	return nil
}

// DB ...
func (s *Server) DB() *att.DB {
	return s.db
}
