package graph

//go:generate go run github.com/99designs/gqlgen

import "github.com/seblkma/ieq/gqlgen-ieq/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// Resolver supports the queries defined in grahql schema
type Resolver struct {
	latestmetrics      *model.Metrics
	latestmetricscores *model.MetricScores
	latestieqscores    *model.IeqScores
}
