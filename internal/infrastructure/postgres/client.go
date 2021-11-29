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
	RouteID              int    `db:"route_id"`
	DirectionID          int    `db:"direction_id"`
	LongestShapeID       int    `db:"longest_shape_id"`
	LongestShapeGeometry string `db:"longest_shape_geometry"`
	LongestShapeCentroid string `db:"longest_shape_centroid"`
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

// func (c *Client) SetRouteShapes(
// 	ctx context.Context,
// 	routeShapes []RouteShapeRow,
// 	gtfsFeedID int,
// ) error {
// 	// we set the shapes in chunk, in a transaction
// 	tx, err := c.db.BeginTxx(ctx, nil)
// 	if err != nil {
// 		return errors.Wrap(err, "could not initiate route shapes set transaction")
// 	}
// 	return nil
// }

func (c *Client) CalculateRouteShapesFromTrips(
	ctx context.Context,
	gtfsFeedID int,
	routeIDs []int,
) ([]RouteShapeRow, error) {
	c.logger.Debugw("starting calculating shapes from trips", "route_ids", routeIDs)

	stmbt := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

	q := stmbt.Select(
		"trips_shapes.route_id",
		"trips_shapes.direction_id",
		"first_value(shapes.id) over ( partition by (route_id, direction_id ) order by ST_LENGTH(shapes.geometry) desc) as longest_shape_id",
		"first_value(st_force2d(shapes.geometry::geometry)::geography) over ( partition by (route_id, direction_id ) order by ST_LENGTH(shapes.geometry) desc) as longest_shape_geometry",
		"first_value(St_centroid(st_force2d(shapes.geometry::geometry))::geography) over ( partition by (route_id, direction_id ) order by ST_LENGTH(shapes.geometry) desc) as longest_shape_centroid",
	).
		Distinct().
		FromSelect(
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
		Where(squirrel.Eq{"shapes.feed_version_id": gtfsFeedID}).
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
