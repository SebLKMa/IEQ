package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"

	db "github.com/seblkma/ieq/db/postgres"
	"github.com/seblkma/ieq/gqlgen-ieq/graph/generated"
	"github.com/seblkma/ieq/gqlgen-ieq/graph/model"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.MetricScores, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Latestmetrics(ctx context.Context, deviceID string) (*model.Metrics, error) {
	result := &model.Metrics{}

	dbmetrics, err := db.ReadMetrics(deviceID, 1)
	if err != nil {
		return result, err
	}
	if len(dbmetrics) == 0 {
		return result, errors.New("no record found")
	}
	record := dbmetrics[0]
	result.DeviceID = record.DeviceID
	result.Timestamp = record.CreatedOn
	result.Temperature = record.Temperature
	result.Humidity = record.Humidity
	result.Co2 = record.CO2
	result.Voc = record.VOC
	result.Pm25 = record.PM25
	result.Noise = record.Noise
	result.Lighting = record.Lighting

	return result, nil
}

func (r *queryResolver) Latestmetricscores(ctx context.Context, deviceID string) (*model.MetricScores, error) {
	result := &model.MetricScores{}
	record, err := db.ReadLatestMetricScores(deviceID)
	if err != nil {
		return result, err
	}

	result.DeviceID = record.DeviceID
	result.Timestamp = record.CreatedOn
	result.Temperature = record.Temperature
	result.Humidity = record.Humidity
	result.Co2 = record.CO2
	result.Voc = record.VOC
	result.Pm25 = record.PM25
	result.Noise = record.Noise
	result.Lighting = record.Lighting

	return result, nil
}

func (r *queryResolver) Latestieqscores(ctx context.Context, deviceID string) (*model.IeqScores, error) {
	result := &model.IeqScores{}
	record, err := db.ReadLatestIeqScores(deviceID)
	if err != nil {
		return result, err
	}

	result.DeviceID = record.DeviceID
	result.Timestamp = record.CreatedOn
	result.Scheme = record.Scheme
	result.Thermal = record.Thermal
	result.Thermalweighting = record.ThermalWeighting
	result.Iaq = record.IAQ
	result.Lighting = record.Lighting
	result.Lightingweighting = record.LightingWeighting
	result.Noise = record.Noise
	result.Noiseweighting = record.NoiseWeighting
	result.Overall = record.Overall

	return result, err
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
