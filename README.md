# Transitland route geometry generator

[![Go Report Card](https://goreportcard.com/badge/github.com/phoops/transitland-route-geometry-generator)](https://goreportcard.com/report/github.com/phoops/transitland-route-geometry-generator)
[![Contributor Covenant](https://img.shields.io/badge/Contributor%20Covenant-2.1-4baaaa.svg)](code_of_conduct.md) 
[![Dockerhub Badge](https://img.shields.io/docker/pulls/phoops/transitland-route-geometry-generator.svg)]("https://img.shields.io/docker/pulls/phoops/transitland-route-geometry-generator.svg")

![Main branch](https://github.com/phoops/transitland-route-geometry-generator/actions/workflows/github-actions-on-push.yaml/badge.svg)
![Main branch docker](https://github.com/phoops/transitland-route-geometry-generator/actions/workflows/github-actions-on-tag.yaml/badge.svg)
## Generate your transitland route shapes from gtfs trips

This project aims to generate [transitland]("https://github.com/interline-io/transitland-server") `gtfs routes` shapes from `gtfs trips` shapes.

## Installation

### Go binary

`go install github.com/phoops/transitland-route-geometry-generator/...@latest`

### Github release - recommended

You can grab the latest release on this repository [releases]("https://github.com/phoops/transitland-route-geometry-generator/releases")

### Docker image

`phoops/transitland-route-geometry-generator`



`docker run run --rm -ti phoops/transitland-route-geometry-generator 1 -d postgres://transit:transit@db/gtfsdb?sslmode=disable`
## How it works?

On `gtfs` spec, we don't have a direct association between `shapes` and `routes`, in `transitland` domain is possibile to associate a `geometry` to a `route`.

In order to associate a `geometry` to a `transitland route`, we need to process the `route trips` and the `shapes` associated to the `trips`.

In order to keep thing simple, we choose the `longest geometry` of all the `trips` and promote that geometry to the geometry of the `route`.

This project depends on `postgresql` and `postgis` and we interact directly with the `transitland` database, so `transitland` schema is needed.

### Steps

- Fetch trips for a specific `gtfs` feed and `routes` (Both can be provided as input to the CLI)
- Choose the `longest` shape associated to the `trips` of the routes
- Take the `geometry` of chosen shape and promote that to `route` geometry
- Persit the `geometry` in `tl_route_geometries table

**The shapes generation is idempotent, so you can use the CLI freely in your workflow, without worring of breaking things, when the shapes geometry are updated by a new feed, they will be imported and the route shape is computed again**

## Usage

```bash
    Generate geometries for your gtfs routes in transistland.
    Uses your trips geometries in order to compute the route shape
    More information at https://github.com/phoops/transitland-route-geometry-generator

    Usage:
    transitland-route-geometry-generator [command]

    Available Commands:
    completion  generate the autocompletion script for the specified shell
    generate    Generate routes
    help        Help about any command
    version     Print the version

    Flags:
    -h, --help      help for transitland-route-geometry-generator
    -v, --verbose   verbose output
```

### Generate routes

```
Usage:
  transitland-route-geometry-generator generate [flags]

Flags:
  -d, --database string   postgres database connection string
  -n, --dry-run           dry run generation, without inserts
  -h, --help              help for generate
  -r, --routes ints       route ids to include in generation, all by default
```

The `-v` flag will output more informations about the process, also the raw `postgresql` queries made by the cli.

The `generate` command wants a single argument, representing the `id` of the `gtfs` feed, this id can be obtained querying `transitland`, or in the
database table `current_feeds`.

The `-d` flag is also required, you should specify a `postgresql` connection string, like `postgres://transit:transit@db/gtfsdb?sslmode=disable`.

The `-r` flag will restrict the `route` shapes generation only to specific route `ids`, routes should be separated by comma, ex:  `-d 1,2,3,4`.

**tl;dr:**



`transitland-route-geometry-generator generate 1 -d "postgres://transit:transit@localhost/gtfsdb?sslmode=disable"`
### Dry run

If you want to only get the computed `shapes` for your routes, you can use the `-n` flag, this will skip the database write of `route shapes`.
The output will be a table like that

```bash
{"level":"info","ts":"2021-11-30T12:56:58.274+0100","caller":"commands/generator.go:50","msg":"starting generation","command":"generate","verbose":false,"dry-run":true,"feed_version_id":1}
+----------+--------------+------------------+
| ROUTE ID | DIRECTION ID | LONGEST SHAPE ID |
+----------+--------------+------------------+
|        1 |            0 |                4 |
|        1 |            1 |                2 |
|        2 |            0 |                1 |
|        2 |            1 |                8 |
+----------+--------------+------------------+
{"level":"info","ts":"2021-11-30T12:56:58.350+0100","caller":"commands/generator.go:88","msg":"dry run completed, run again without the n flag","command":"generate","verbose":false}
```

If you pass the `-v` flag, the output will contain also the raw `geometry` and the `centroid` of shape.

**tl;dr:**


`transitland-route-geometry-generator generate 1 -n -d "postgres://transit:transit@localhost/gtfsdb?sslmode=disable"`

