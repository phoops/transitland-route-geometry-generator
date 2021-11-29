package commands

import (
	"os"
	"strconv"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/olekukonko/tablewriter"
	"github.com/phoops/transitland-route-geometry-generator/internal/infrastructure/logger"
	"github.com/phoops/transitland-route-geometry-generator/internal/infrastructure/postgres"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

// thanks cobra, dependency injection is for dummies, apri tutto

var DryRun bool
var DbConnectionString string
var RouteIDs []int

func init() {
	rootCmd.AddCommand(generateCmd)
	generateCmd.Flags().BoolVarP(&DryRun, "dry-run", "n", false, "dry run generation, without inserts")
	generateCmd.Flags().StringVarP(&DbConnectionString, "database", "d", "", "dry run generation, without inserts")
	generateCmd.Flags().IntSliceVarP(&RouteIDs, "routes", "r", []int{}, "route ids to include in generation, all by default")

	err := generateCmd.MarkFlagRequired("database")
	if err != nil {
		panic(err)
	}
}

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate routes",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		cmdLogger := logger.LoggerForSubCommandFromContext(
			cmd.Context(),
			"generate",
			Verbose,
		)
		feedVersionID, err := strconv.Atoi(args[0])
		if err != nil {
			cmdLogger.Errorw("invalid feed version id, should be a int", zap.Error(err))
			return err
		}

		cmdLogger.Infow("starting generation", "dry-run", DryRun, "feed_version_id", feedVersionID)

		db, err := sqlx.Connect("postgres", DbConnectionString)
		if err != nil {
			cmdLogger.Errorw(
				"could not connect to the postgres database",
				"db_connection_string",
				DbConnectionString,
				zap.Error(err),
			)
			return err
		}
		client := postgres.NewClient(cmdLogger, db)
		result, err := client.CalculateRouteShapesFromTrips(
			cmd.Context(),
			feedVersionID,
			RouteIDs,
		)
		if err != nil {
			cmdLogger.Errorw(
				"error during the select route shapes query",
				zap.Error(err),
			)

			return err
		}

		if DryRun {
			tableResult := tableizeDbRows(result, Verbose)
			th := []string{"Route_ID", "Direction_ID", "Longest_Shape_ID"}
			if Verbose {
				th = append(th, "LongestShapeGeometry", "LongestShapeCentroid")
			}
			table := tablewriter.NewWriter(os.Stdout)
			table.SetHeader(th)
			table.AppendBulk(tableResult)
			table.Render()

			cmdLogger.Info("dry run completed, run again without the n flag")
			return nil
		}
		return nil
	},
}

func tableizeDbRows(rows []postgres.RouteShapeRow, verbose bool) [][]string {
	var output [][]string
	for _, r := range rows {
		tr := []string{strconv.Itoa(r.RouteID), strconv.Itoa(r.DirectionID), strconv.Itoa(r.LongestShapeID)}
		if verbose {
			tr = append(tr, r.LongestShapeGeometry, r.LongestShapeCentroid)
		}
		output = append(output, tr)
	}

	return output
}
