package mongodb

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/options"
	"go.mongodb.org/mongo-driver/readpref"

	"github.com/thearyanahmed/url-shortner/shortner"
)

type mongoRepository struct {
	client *mongo.Client
	database string
	timeout time.Duration
}