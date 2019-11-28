package usecase

import "github.com/hieuphq/tontine/src/model"

// GroupUC group usecase
type GroupUC interface {
	Create(g model.Group) (*model.Group, error)
	Update(g model.Group) (*model.Group, error)
	Close(g model.Group) (*model.Group, error)
}

type implGroup struct{}

// func NewGroupUC()
