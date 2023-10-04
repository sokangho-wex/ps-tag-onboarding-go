package persistence

import (
	"context"
	"errors"
	"github.com/sokangho-wex/ps-tag-onboarding-go/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const userCollection = "user"

type UserRepository struct {
	db *mongo.Database
}

func NewUserRepo(db *mongo.Database) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) FindByID(id string) (models.User, error) {
	filter := bson.D{{"_id", id}}

	var result models.User
	err := r.db.Collection(userCollection).FindOne(context.TODO(), filter).Decode(&result)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return models.User{}, models.UserNotFoundError
		} else {
			return models.User{}, models.UnexpectedError
		}
	}

	return result, nil
}

func (r *UserRepository) AddUser(user models.User) error {
	_, err := r.db.Collection(userCollection).InsertOne(context.TODO(), user)

	if err != nil {
		return models.UnexpectedError
	}

	return nil
}
