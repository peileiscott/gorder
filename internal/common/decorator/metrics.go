package decorator

import (
	"context"
	"fmt"
	"time"
)

type MetricsClient interface {
	Inc(key string, value int)
}

type commandMetricsDecorator[C, R any] struct {
	base   CommandHandler[C, R]
	client MetricsClient
}

func (c commandMetricsDecorator[C, R]) Handle(ctx context.Context, command C) (result R, err error) {
	commandName := generateActionName(command)
	start := time.Now()
	defer func() {
		end := time.Since(start)
		c.client.Inc(fmt.Sprintf("command.%s.duration", commandName), int(end.Seconds()))
		if err != nil {
			c.client.Inc(fmt.Sprintf("command.%s.failure", commandName), 1)
		} else {
			c.client.Inc(fmt.Sprintf("command.%s.success", commandName), 1)
		}
	}()
	return c.base.Handle(ctx, command)
}

type queryMetricsDecorator[Q, R any] struct {
	base   QueryHandler[Q, R]
	client MetricsClient
}

func (q queryMetricsDecorator[Q, R]) Handle(ctx context.Context, query Q) (result R, err error) {
	queryName := generateActionName(query)
	start := time.Now()
	defer func() {
		end := time.Since(start)
		q.client.Inc(fmt.Sprintf("query.%s.duration", queryName), int(end.Seconds()))
		if err != nil {
			q.client.Inc(fmt.Sprintf("query.%s.failure", queryName), 1)
		} else {
			q.client.Inc(fmt.Sprintf("query.%s.success", queryName), 1)
		}
	}()
	return q.base.Handle(ctx, query)
}
