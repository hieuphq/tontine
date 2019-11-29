package repo

import (
	"context"

	"github.com/hieuphq/tontine/src/interfaces/store"
	"github.com/hieuphq/tontine/src/model"
)

// ActivityLog activity log persistent interface
type ActivityLog interface {
	LogGroup(ctx context.Context, store store.Store, g model.GroupLog) (*model.GroupLog, error)
	LogGroupItem(ctx context.Context, store store.Store, g model.GroupLogItem) (*model.GroupLogItem, error)
	LogActivity(ctx context.Context, store store.Store, g model.ActivityLog) (*model.ActivityLog, error)
	GetGroupLogs(ctx context.Context, store store.Store, g model.Group) ([]model.GroupLog, error)
	GetActivityLogs(ctx context.Context, store store.Store, invt model.Investor, g *model.Group) ([]model.ActivityLog, error)
}
