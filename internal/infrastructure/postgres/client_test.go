package postgres_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/phoops/transitland-route-geometry-generator/internal/infrastructure/postgres"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

var testLogger *zap.SugaredLogger

const (
	dbConnectionString = "host=localhost port=5432 dbname=gtfsdb user=transit password=transit sslmode=disable timezone=UTC"
)

func init() {
	l, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}

	testLogger = l.Sugar()
}

func TestClientRouteShapesCalculationNoRows(t *testing.T) {
	pgClient := sqlx.MustConnect("postgres", dbConnectionString)

	defer func() {
		err := pgClient.Close()
		if err != nil {
			panic(err)
		}
	}()

	client := postgres.NewClient(
		testLogger,
		pgClient,
	)

	rows, err := client.CalculateRouteShapesFromTrips(context.TODO(), 1, []int{55})
	assert.Error(t, err, fmt.Errorf(
		"no route shapes calculated, zero result from the query, route ids: %v",
		[]int{55},
	))
	assert.Nil(t, rows)
}

func TestClientRouteShapesCalculationSuccess(t *testing.T) {
	pgClient := sqlx.MustConnect("postgres", dbConnectionString)

	defer func() {
		err := pgClient.Close()
		if err != nil {
			panic(err)
		}
	}()

	client := postgres.NewClient(
		testLogger,
		pgClient,
	)

	rows, err := client.CalculateRouteShapesFromTrips(context.TODO(), 1, nil)

	assert.NoError(t, err)

	// separate assertions for better readability

	assert.Len(t, rows, 4)

	assert.EqualValues(t, rows[0].RouteID, 1)
	assert.EqualValues(t, rows[0].DirectionID, 0)
	assert.EqualValues(t, rows[0].LongestShapeID, 4)

	assert.EqualValues(t, rows[1].RouteID, 1)
	assert.EqualValues(t, rows[1].DirectionID, 1)
	assert.EqualValues(t, rows[1].LongestShapeID, 2)

	assert.EqualValues(t, rows[2].RouteID, 2)
	assert.EqualValues(t, rows[2].DirectionID, 0)
	assert.EqualValues(t, rows[2].LongestShapeID, 1)

	assert.EqualValues(t, rows[3].RouteID, 2)
	assert.EqualValues(t, rows[3].DirectionID, 1)
	assert.EqualValues(t, rows[3].LongestShapeID, 8)
}
