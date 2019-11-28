package store

import (
	"context"

	"github.com/volatiletech/sqlboiler/boil"
)

// FinallyFunc function to finish a transaction
type FinallyFunc = func(error) error

// Store contains db context for
type Store interface {
	DB() boil.ContextExecutor
	BeginTx(ctx context.Context) (Store, FinallyFunc)
}

// NewTestDBRepo impl repo for testing
func NewTestDBRepo() Store {
	return &testRepo{}
}

type testRepo struct {
}

func (*testRepo) DB() boil.ContextExecutor {
	return nil
}

func (repo *testRepo) BeginTx(ctx context.Context) (Store, FinallyFunc) {
	return repo, func(err error) error { return err }
}
