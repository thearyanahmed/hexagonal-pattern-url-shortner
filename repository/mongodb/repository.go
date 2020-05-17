// package mongodb

// import (
// 	"context"
// 	"time"

// 	"github.com/pkg/errors"
// 	"go.mongodb.org/mongo-driver/bson"
// 	"	"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// 	"go.mongodb.org/mongo-driver/mongo/readpref"

// 	"github.com/thearyanahmed/url-shortner/shortener"
// )

// type mongoRepository struct {
// 	client *mongo.Client
// 	database string
// 	timeout time.Duration
// }

// func newMongoClient (url string, timeout int) (*mongo.Client, error) {

// 	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout) * time.Second)
// 	defer cancle()

// 	client , err := mongo.Connect(ctx, options.Client().ApplyURI(url))
// 	if err != nil {
// 		return nil, err
// 	}

// 	err = client.Ping(ctx, readpref.Primary())
// 	if err != nil {
// 		return nil, err
// 	}

// 	return client, nil
// }

// func NewMongoRepository(mongoURL, mongoDB string, mongoTimeout int) ( shortner.RedirectRepository, error ) {
// 	repo := &mongoRepository{
// 		timeout:  time.Duration(mongoTimeout) * time.Second,
// 		database: mongoDB,
// 	}
// 	client, err := newMongoClient(mongoURL, mongoTimeout)
// 	if err != nil {
// 		return nil, errors.Wrap(err, "repository.NewMongoRepo")
// 	}
// 	repo.client = client
// 	return repo, nil
// }

