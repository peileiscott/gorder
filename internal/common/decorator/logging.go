package decorator

import (
	"context"
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"
)

type commandLoggingDecorator[C any] struct {
	base   CommandHandler[C]
	logger *logrus.Entry
}

func (c commandLoggingDecorator[C]) Handle(ctx context.Context, command C) (err error) {
	c.logger.WithFields(logrus.Fields{
		"command":      generateActionName(command),
		"command_body": fmt.Sprintf("%#v", command),
	}).Debug("Executing command")

	defer func() {
		if err != nil {
			c.logger.WithError(err).Error("Failed to execute command")
		} else {
			c.logger.Debug("Command executed successfully")
		}
	}()

	return c.base.Handle(ctx, command)
}

type queryLoggingDecorator[Q, R any] struct {
	base   QueryHandler[Q, R]
	logger *logrus.Entry
}

func (q queryLoggingDecorator[Q, R]) Handle(ctx context.Context, query Q) (result R, err error) {
	q.logger.WithFields(logrus.Fields{
		"query":      generateActionName(query),
		"query_body": fmt.Sprintf("%#v", query),
	}).Debug("Executing query")

	defer func() {
		if err != nil {
			q.logger.WithError(err).Error("Failed to execute query")
		} else {
			q.logger.Debug("Query executed successfully")
		}
	}()

	return q.base.Handle(ctx, query)
}

func generateActionName(action any) string {
	return strings.Split(fmt.Sprintf("%T", action), ".")[1]
}
