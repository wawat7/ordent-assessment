package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"ordent-assessment/entity"
	"time"
)

type ProductRepository interface {
	FindAll(ctx context.Context) ([]entity.Product, error)
	FindById(ctx context.Context, Id string) (entity.Product, error)
	Create(ctx context.Context, product entity.Product) (string, error)
	Update(ctx context.Context, product entity.Product) error
	Delete(ctx context.Context, product entity.Product) error
}

type productRepository struct {
	collection *mongo.Collection
}

func NewProductRepository(database *mongo.Database) *productRepository {
	return &productRepository{collection: database.Collection("products")}
}

func (repo *productRepository) FindAll(ctx context.Context) ([]entity.Product, error) {
	var products []entity.Product

	results, err := repo.collection.Find(ctx, bson.M{"deleted_at": nil})
	if err != nil {
		return products, err
	}

	defer results.Close(ctx)

	for results.Next(ctx) {
		var product entity.Product
		err := results.Decode(&product)
		if err != nil {
			return products, err
		}
		products = append(products, product)
	}
	return products, nil
}

func (repo *productRepository) FindById(ctx context.Context, Id string) (entity.Product, error) {
	var product entity.Product

	objId, _ := primitive.ObjectIDFromHex(Id)
	err := repo.collection.FindOne(ctx, bson.M{"_id": objId, "deleted_at": nil}).Decode(&product)
	if err != nil {
		return product, err
	}
	return product, nil
}

func (repo *productRepository) Create(ctx context.Context, product entity.Product) (string, error) {
	product.CreatedAt = time.Now()
	product.UpdatedAt = time.Now()
	result, err := repo.collection.InsertOne(ctx, product)
	if err != nil {
		return result.InsertedID.(primitive.ObjectID).Hex(), err
	}
	return result.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (repo *productRepository) Update(ctx context.Context, product entity.Product) error {
	product.UpdatedAt = time.Now()
	_, err := repo.collection.UpdateOne(ctx, bson.M{"_id": product.Id}, bson.M{"$set": product})
	if err != nil {
		return err
	}
	return nil
}

func (repo *productRepository) Delete(ctx context.Context, product entity.Product) error {
	product.DeletedAt = time.Now()

	err := repo.Update(ctx, product)
	if err != nil {
		return err
	}
	return nil
}
