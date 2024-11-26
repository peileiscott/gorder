package decorator

import (
	"context"
	"fmt"
	"time"
)

type MetricsClient interface {
	Inc(key string, value int)
}

type queryMetricsDecorator[Q, R any] struct {
	client  MetricsClient
	handler QueryHandler[Q, R]
}

func (q queryMetricsDecorator[Q, R]) Handle(ctx context.Context, query Q) (result R, err error) {
	start := time.Now()
	defer func() {
		end := time.Since(start)
		q.client.Inc(fmt.Sprintf("query.%T.duration", query), int(end.Seconds()))
		if err != nil {
			q.client.Inc(fmt.Sprintf("query.%T.failure", query), 1)
		} else {
			q.client.Inc(fmt.Sprintf("query.%T.success", query), 1)
		}
	}()
	return q.handler.Handle(ctx, query)
}

type commandMetricsDecorator[C, R any] struct {
	client  MetricsClient
	handler CommandHandler[C, R]
}

func (c commandMetricsDecorator[C, R]) Handle(ctx context.Context, command C) (result R, err error) {
	start := time.Now()
	defer func() {
		end := time.Since(start)
		c.client.Inc(fmt.Sprintf("command.%T.duration", command), int(end.Seconds()))
		if err != nil {
			c.client.Inc(fmt.Sprintf("command.%T.failure", command), 1)
		} else {
			c.client.Inc(fmt.Sprintf("command.%T.success", command), 1)
		}
	}()
	return c.handler.Handle(ctx, command)
}
