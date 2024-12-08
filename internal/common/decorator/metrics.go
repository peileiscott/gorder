package decorator

import (
	"context"
	"fmt"
	"time"
)

type MetricsClient interface {
	Inc(key string, value int)
}

type commandMetricsDecorator[C any] struct {
	base   CommandHandler[C]
	client MetricsClient
}

func (c commandMetricsDecorator[C]) Handle(ctx context.Context, command C) (err error) {
	start := time.Now()
	actionName := generateActionName(command)

	defer func() {
		end := time.Since(start)
		c.client.Inc(fmt.Sprintf("command.%s.duration", actionName), int(end.Seconds()))
		if err != nil {
			c.client.Inc(fmt.Sprintf("command.%s.failure", actionName), 1)
		} else {
			c.client.Inc(fmt.Sprintf("command.%s.success", actionName), 1)
		}
	}()

	return c.base.Handle(ctx, command)
}

type queryMetricsDecorator[Q, R any] struct {
	base   QueryHandler[Q, R]
	client MetricsClient
}

func (q queryMetricsDecorator[Q, R]) Handle(ctx context.Context, query Q) (result R, err error) {
	start := time.Now()
	actionName := generateActionName(query)

	defer func() {
		end := time.Since(start)
		q.client.Inc(fmt.Sprintf("query.%s.duration", actionName), int(end.Seconds()))
		if err != nil {
			q.client.Inc(fmt.Sprintf("query.%s.failure", actionName), 1)
		} else {
			q.client.Inc(fmt.Sprintf("query.%s.success", actionName), 1)
		}
	}()

	return q.base.Handle(ctx, query)
}
