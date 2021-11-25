package postgres

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type RouteShapeRow struct {
	RouteID     string `db:"route_id"`
	DirectionID string `db:"direction_id"`
	Geometry    string `db:"geometry"`
	Centroid    string `db:"centroid"`
}

type Client struct {
	logger *zap.SugaredLogger
	db     *sqlx.DB
}

func NewClient(
	logger *zap.SugaredLogger,
	db *sqlx.DB,
) *Client {
	l := logger.With("component", "postgres-client")

	return &Client{
		logger: l,
		db:     db,
	}
}

func (c *Client) CalculateRouteShapesFromTrips(
	ctx context.Context,
	routeIDs []int,
) ([]RouteShapeRow, error) {
	c.logger.Debugw("starting calculating shapes from trips", "route_ids", routeIDs)

	stmbt := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

	q := stmbt.Select(
		"trips_shapes.route_id",
		"trips_shapes.direction_id",
		"ST_ASTEXT(ST_UNION(ST_FORCE2D(shapes.geometry::geometry))::geography) as geometry",
		"ST_ASTEXT(ST_CENTROID(ST_UNION(ST_FORCE2D(shapes.geometry::geometry)))::geography) as centroid",
	).FromSelect(
		stmbt.Select(
			"route_id",
			"direction_id",
			"gt.shape_id",
		).
			LeftJoin(
				"gtfs_shapes gs on gs.id = gt.shape_id",
			).
			From("gtfs_trips gt").
			GroupBy("route_id", "direction_id", "gt.shape_id"),
		"trips_shapes",
	).
		Join("gtfs_shapes shapes on trips_shapes.shape_id = shapes.id").
		GroupBy("route_id", "direction_id").
		OrderBy("route_id")

	if routeIDs != nil {
		q = q.Where(squirrel.Eq{"route_id": routeIDs})
	}

	query, args, err := q.ToSql()
	if err != nil {
		return nil, errors.Wrapf(
			err,
			"could not build the route sapes calculation query for routes: %v",
			routeIDs,
		)
	}

	c.logger.Debugw(
		"route shapes calculation query built",
		"query",
		query,
		"args",
		args,
	)

	var result []RouteShapeRow

	err = c.db.SelectContext(ctx, &result, query, args...)
	if err != nil {
		return nil, errors.Wrapf(
			err,
			"could not perform the route shapes calculation query for routes: %v",
			routeIDs,
		)
	}

	if len(result) == 0 {
		return nil, fmt.Errorf("no route shapes calculated, zero result from the query, route ids: %v", routeIDs)
	}

	c.logger.Debugw(
		"calculated shapes for routes",
		"route_shapes_row",
		result,
	)

	return result, nil
}
