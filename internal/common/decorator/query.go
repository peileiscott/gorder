package decorator

import (
	"context"

	"github.com/sirupsen/logrus"
)

// QueryHandler defines a generic type that receives a Query Q,
// and returns a result R.
type QueryHandler[Q, R any] interface {
	Handle(ctx context.Context, query Q) (R, error)
}

func ApplyQueryDecorators[Q, R any](handler QueryHandler[Q, R], logger *logrus.Entry, client MetricsClient) QueryHandler[Q, R] {
	return queryLoggingDecorator[Q, R]{
		base: queryMetricsDecorator[Q, R]{
			base:   handler,
			client: client,
		},
		logger: logger,
	}
}
