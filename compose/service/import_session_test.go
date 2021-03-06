package service

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

var ctx context.Context = context.WithValue(context.Background(), "testing", true)

func TestFindRecordByID(t *testing.T) {
	DefaultLogger = zap.New(nil, nil)
	svc := ImportSession()
	ss, _ := svc.SetRecordByID(ctx, 1, 0, 0, nil, nil, nil)
	sid := ss.SessionID

	t.Run("Found", func(t *testing.T) {
		s, err := svc.FindRecordByID(ctx, sid)
		require.True(t,
			s != nil,
			"Session should be found",
		)

		require.True(t,
			err == nil,
			"Returned with error",
		)
	})

	t.Run("Not found", func(t *testing.T) {
		s, err := svc.FindRecordByID(ctx, sid+1)
		require.True(t,
			s == nil,
			"Session should not be found",
		)

		require.True(t,
			err != nil,
			"Error should not be nil",
		)
	})
}

func TestSetRecordByID(t *testing.T) {
	DefaultLogger = zap.New(nil, nil)
	svc := ImportSession()

	t.Run("New", func(t *testing.T) {
		ss, err := svc.SetRecordByID(ctx, 1, 0, 0, nil, nil, nil)
		require.True(t,
			len(svc.records) == 1 && ss != nil,
			"Session should be created",
		)

		require.True(t,
			err == nil,
			"Returned with error",
		)
	})

	t.Run("Existing", func(t *testing.T) {
		svc := ImportSession()
		ss, err := svc.SetRecordByID(ctx, 1, 0, 0, nil, nil, nil)
		ns, err := svc.SetRecordByID(ctx, ss.SessionID, 0, 0, nil, nil, nil)
		require.True(t,
			len(svc.records) == 1 && ns != nil && ss.SessionID == ns.SessionID,
			"Existing session should be edited",
		)

		require.True(t,
			err == nil,
			"Returned with error",
		)
	})
}

func TestDeleteRecordByID(t *testing.T) {
	DefaultLogger = zap.New(nil, nil)
	svc := ImportSession()
	ss, _ := svc.SetRecordByID(ctx, 1, 0, 0, nil, nil, nil)

	t.Run("Delete existing", func(t *testing.T) {
		err := svc.DeleteRecordByID(ctx, ss.SessionID)
		require.True(t,
			len(svc.records) == 0,
			"Session should be deleted",
		)

		require.True(t,
			err == nil,
			"Returned with error",
		)
	})

	t.Run("Session not found", func(t *testing.T) {
		ss, _ := svc.SetRecordByID(ctx, 1, 0, 0, nil, nil, nil)
		err := svc.DeleteRecordByID(ctx, ss.SessionID+1)
		require.True(t,
			len(svc.records) == 1,
			"Session should not deleted",
		)

		require.True(t,
			err == nil,
			"Returned with error",
		)
	})
}
