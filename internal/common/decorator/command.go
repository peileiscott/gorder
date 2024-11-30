package decorator

import (
	"context"

	"github.com/sirupsen/logrus"
)

// CommandHandler defines a generic type that receives a Command C,
// and returns a result R.
type CommandHandler[C, R any] interface {
	Handle(ctx context.Context, command C) (R, error)
}

func ApplyCommandDecorators[C, R any](handler CommandHandler[C, R], logger *logrus.Entry, client MetricsClient) CommandHandler[C, R] {
	return commandLoggingDecorator[C, R]{
		base: commandMetricsDecorator[C, R]{
			base:   handler,
			client: client,
		},
		logger: logger,
	}
}
