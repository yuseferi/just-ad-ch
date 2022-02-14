package log

import (
	"context"

	"go.uber.org/zap"
)

// Get zap.Logger from context
func Get(ctx context.Context) *zap.Logger {
	return zap.L()
}
