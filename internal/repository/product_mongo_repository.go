package repository

import (
	"context"
	"fmt"

	"github.com/vagnerlg/deliveryApi/internal/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type repo struct {
	coll *mongo.Collection
}

func New(coll *mongo.Collection) *repo {
	return &repo{coll: coll}
}

func (r *repo) Create(product *entity.Product) {
	result, _ := r.coll.InsertOne(context.TODO(), product)
	product.ID = result.InsertedID
}

func (r *repo) Update(product *entity.Product) {
	idObject, _ := primitive.ObjectIDFromHex(product.ID.(string))
	product.ID = idObject
	filter := bson.M{"_id": idObject}
	fmt.Println(product)
	result, error := r.coll.ReplaceOne(context.TODO(), filter, product)
	if error != nil {
		fmt.Println(error)
	}
	fmt.Println(result)
}

func (r *repo) List() []entity.Product {
	cursor, _ := r.coll.Find(context.TODO(), bson.D{})

	var results []entity.Product
	if err := cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	return results
}

func (r *repo) Get(id string) entity.Product {
	var product entity.Product
	idObject, _ := primitive.ObjectIDFromHex(id)

	result := r.coll.FindOne(context.TODO(), bson.M{"_id": idObject})
	if result.Err() != nil {
		panic(result)
	}

	result.Decode(&product)

	return product
}
