package checkers

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Mongo implements the ICheckable interface
type Mongo struct {
	config *MongoConfig
}

// MongoPinger an interface that allows direct pinging of the database
type MongoPinger interface {
	Ping(ctx context.Context, rp *readpref.ReadPref) error
}

// MongoConfig configuration for MongoDB database check
type MongoConfig struct {
	Pinger MongoPinger
}

// Status checks MongoDB connection status
func (mc *Mongo) Status() (interface{}, error) {
	err := mc.config.Pinger.Ping(context.Background(), readpref.Primary())
	return nil, err
}

// NewMongo creates a new MongoDB database checker
func NewMongo(config *MongoConfig) (*Mongo, error) {
	if config == nil {
		return nil, errors.New("undefined config")
	}

	if config.Pinger == nil {
		return nil, errors.New("undefined Pinger")
	}

	return &Mongo{config: config}, nil
}
