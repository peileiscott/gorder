package decorator

import (
	"context"

	"github.com/sirupsen/logrus"
)

type CommandHandler[C any] interface {
	Handle(ctx context.Context, command C) error
}

func ApplyCommandDecorators[C any](handler CommandHandler[C], logger *logrus.Entry, client MetricsClient) CommandHandler[C] {
	return commandLoggingDecorator[C]{
		base: commandMetricsDecorator[C]{
			base:   handler,
			client: client,
		},
		logger: logger,
	}
}
