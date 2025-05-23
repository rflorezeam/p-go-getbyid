package repositories

import (
	"context"

	"github.com/rflorezeam/libro-read-one/config"
	"github.com/rflorezeam/libro-read-one/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type LibroRepository interface {
	ObtenerLibroPorID(id string) (*models.Libro, error)
}

type libroRepository struct {
	collection *mongo.Collection
}

func NewLibroRepository() LibroRepository {
	return &libroRepository{
		collection: config.GetCollection(),
	}
}

func (r *libroRepository) ObtenerLibroPorID(id string) (*models.Libro, error) {
	var libro models.Libro
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	err = r.collection.FindOne(context.TODO(), bson.M{"_id": objectID}).Decode(&libro)
	if err != nil {
		return nil, err
	}

	return &libro, nil
} 