package config

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)


func(server *Server) DBConection() {
ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
server.Client, _ = mongo.Connect(ctx, clientOptions)
}
