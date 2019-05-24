package service

import (
	"context"
	"time"

	"go.uber.org/zap"

	"github.com/cortezaproject/corteza-server/compose/internal/repository"
	"github.com/cortezaproject/corteza-server/internal/permissions"
	"github.com/cortezaproject/corteza-server/internal/store"
	"github.com/cortezaproject/corteza-server/pkg/logger"
)

type (
	permissionServicer interface {
		accessControlPermissionServicer
		Watch(ctx context.Context)
	}
)

var (
	DefaultPermissions permissionServicer

	DefaultLogger *zap.Logger

	DefaultAccessControl *accessControl

	DefaultRecord       RecordService
	DefaultModule       ModuleService
	DefaultTrigger      TriggerService
	DefaultChart        ChartService
	DefaultPage         PageService
	DefaultNotification NotificationService
	DefaultAttachment   AttachmentService
	DefaultNamespace    NamespaceService
)

func Init(ctx context.Context) error {
	DefaultLogger = logger.Default().Named("compose.service")

	fs, err := store.New("var/store")
	if err != nil {
		return err
	}

	DefaultPermissions = permissions.Service(
		ctx,
		DefaultLogger,
		permissions.Repository(repository.DB(ctx), "compose_permission_rules"))

	DefaultAccessControl = AccessControl(DefaultPermissions)

	DefaultRecord = Record()
	DefaultModule = Module()
	DefaultTrigger = Trigger()
	DefaultPage = Page()
	DefaultChart = Chart()
	DefaultNotification = Notification()
	DefaultAttachment = Attachment(fs)
	DefaultNamespace = Namespace()

	return nil
}

// Data is stale when new date does not match updatedAt or createdAt (before first update)
func isStale(new *time.Time, updatedAt *time.Time, createdAt time.Time) bool {
	if new == nil {
		// Change to true for stale-data-check
		return false
	}

	if updatedAt != nil {
		return !new.Equal(*updatedAt)
	}

	return new.Equal(createdAt)
}
