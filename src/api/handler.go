package api

import (
	"github.com/hieuphq/tontine/src/interfaces/repo"
	"github.com/hieuphq/tontine/src/interfaces/store"
)

// Handlers http handler
type Handlers interface {
	InvestorHandler
	GroupHandler
}

type impl struct {
	store store.Store
	repo  *repo.Repo
}

// NewHandlers init handler
func NewHandlers(s store.Store, r *repo.Repo) Handlers {
	return &impl{
		store: s,
		repo:  r,
	}

}
