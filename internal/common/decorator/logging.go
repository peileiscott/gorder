package decorator

import (
	"context"
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"
)

type commandLoggingDecorator[C, R any] struct {
	base   CommandHandler[C, R]
	logger *logrus.Entry
}

func (c commandLoggingDecorator[C, R]) Handle(ctx context.Context, command C) (result R, err error) {
	logger := c.logger.WithFields(logrus.Fields{
		"command":      generateActionName(command),
		"command_body": fmt.Sprintf("%#v", command),
	})
	logger.Debug("Executing command")
	defer func() {
		if err != nil {
			logger.Error("Failed to execute command: ", err)
		} else {
			logger.Info("Command executed successfully")
		}
	}()
	return c.base.Handle(ctx, command)
}

type queryLoggingDecorator[Q, R any] struct {
	base   QueryHandler[Q, R]
	logger *logrus.Entry
}

func (q queryLoggingDecorator[Q, R]) Handle(ctx context.Context, query Q) (result R, err error) {
	logger := q.logger.WithFields(logrus.Fields{
		"query":      generateActionName(query),
		"query_body": fmt.Sprintf("%#v", query),
	})
	logger.Debug("Executing query")
	defer func() {
		if err != nil {
			logger.Error("Failed to execute query: ", err)
		} else {
			logger.Info("Query executed successfully")
		}
	}()
	return q.base.Handle(ctx, query)
}

func generateActionName(action any) string {
	return strings.Split(fmt.Sprintf("%T", action), ".")[1]
}
