package service

import (
	"context"

	"github.com/AliShan302/Task_Schedular/db"
)

// Service initializes our database instance
type Service struct {
	db    db.DataStore
	Store db.DataStore
}

// NewService creates a connection to our database
func NewService(ds db.DataStore) *Service {
	return &Service{db: ds}
}
func (s *Service) Close(ctx context.Context) error {
	return s.db.Disconnect(ctx)
}
