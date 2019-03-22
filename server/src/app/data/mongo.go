package mongo

import (
	"context"
	"time"

	"github.com/mongodb/mongo-go-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)

var client, err = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
