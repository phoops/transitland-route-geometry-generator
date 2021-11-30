package commands

import (
	"context"
	"fmt"
	"os"

	"github.com/phoops/transitland-route-geometry-generator/internal/infrastructure/logger"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Verbose bool

func init() {
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
}

var rootCmd = &cobra.Command{
	Use:   "transitland-route-geometry-generator",
	Short: "transitland-route-geometry-generator",
	Long: `
Generate geometries for your gtfs routes in transistland.
Uses your trips geometries in order to compute the route shape
More information at https://github.com/phoops/transitland-route-geometry-generator`,
}

func Execute() {
	var appL *zap.SugaredLogger
	ctx := context.Background()
	conf := zap.NewProductionConfig()
	conf.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	conf.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	l, err := conf.Build()
	if err != nil {
		panic(err)
	}
	appL = l.Sugar()
	rootCtx := context.WithValue(ctx, logger.RootLoggerKey, appL) //nolint
	if err := rootCmd.ExecuteContext(rootCtx); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
