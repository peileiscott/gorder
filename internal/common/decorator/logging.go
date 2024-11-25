package decorator

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"
)

type queryLoggingDecorator[Q, R any] struct {
	logger  *logrus.Entry
	handler QueryHandler[Q, R]
}

func (q queryLoggingDecorator[Q, R]) Handle(ctx context.Context, query Q) (result R, err error) {
	logger := q.logger.WithFields(logrus.Fields{
		"query":      fmt.Sprintf("%T", query),
		"query_body": fmt.Sprintf("%#v", query),
	})
	logger.Debug("Executing query")
	defer func() {
		if err != nil {
			logger.Error("Query failed ", err)
		} else {
			logger.Info("Query executed successfully")
		}
	}()
	return q.handler.Handle(ctx, query)
}
