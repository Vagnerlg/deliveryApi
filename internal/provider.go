package provider

import (
	"context"

	"github.com/vagnerlg/deliveryApi/config"
	"github.com/vagnerlg/deliveryApi/internal/entity"
	"github.com/vagnerlg/deliveryApi/internal/repository"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Repository() entity.ProductRespository {
	configDB := config.DB()
	client, err := mongo.Connect(context.TODO(), options.Client().
		ApplyURI(configDB.Uri))
	if err != nil {
		panic(err)
	}

	coll := client.Database(configDB.Name).Collection("product")

	return repository.New(coll)
}
