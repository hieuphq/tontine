package sqlite

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/volatiletech/sqlboiler/boil"

	"github.com/hieuphq/tontine/src/interfaces/repo"
	"github.com/hieuphq/tontine/src/interfaces/store"
	"github.com/hieuphq/tontine/src/model/errors"
	"github.com/hieuphq/tontine/src/util"
)

// impl is implimentation of repository
type impl struct {
	db *sql.DB
	tx *sql.Tx
}

// NewRepo new default Repo
func NewRepo() *repo.Repo {
	return &repo.Repo{
		Group:       newGroupRepo(),
		ActivityLog: newActivityLogRepo(),
		Investor:    newInvestorRepo(),
	}
}

// NewStore new a sqlite imple for Repo
func NewStore(filePath string) store.Store {
	db, err := sql.Open("sqlite3", filePath)
	if err != nil {
		panic(err)
	}

	return &impl{
		db: db,
		tx: nil,
	}
}

// DB database connection
func (r *impl) DB() boil.ContextExecutor {
	if r.tx != nil {
		return r.tx
	}

	return r.db
}

// BeginTx for database connection
func (r *impl) BeginTx(ctx context.Context) (newRepo store.Store, finallyFn store.FinallyFunc) {
	newTx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, nil
	}

	finallyFn = func(err error) error {
		if err != nil {
			nErr := newTx.Rollback()
			if nErr != nil {
				return errors.NewStringError(nErr.Error(), http.StatusInternalServerError)
			}
			return errors.NewStringError(err.Error(), util.ParseErrorCode(err))
		}

		cErr := newTx.Commit()
		if cErr != nil {
			return errors.NewStringError(cErr.Error(), http.StatusInternalServerError)
		}
		return nil
	}

	return &impl{
		db: r.db,
		tx: newTx,
	}, finallyFn
}
