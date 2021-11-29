package logger

import (
	"context"
	"fmt"

	"go.uber.org/zap"
)

const RootLoggerKey = "logger"

func LoggerForSubCommandFromContext(
	ctx context.Context,
	commandName string,
	verbose bool,
) *zap.SugaredLogger {
	logger, ok := ctx.Value(RootLoggerKey).(*zap.SugaredLogger)
	if !ok {
		panic(fmt.Errorf("could not extract logger from context at key: %s", RootLoggerKey))
	}

	l := logger.With("command", commandName).With("verbose", verbose)
	if !verbose {
		l = l.Desugar().
			WithOptions(
				zap.IncreaseLevel(zap.InfoLevel),
			).
			Sugar()

	}
	return l
}
