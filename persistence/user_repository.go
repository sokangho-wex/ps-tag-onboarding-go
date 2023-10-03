package persistence

import (
	"context"
	"errors"
	"github.com/sokangho-wex/ps-tag-onboarding-go/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// TODO: Need to create an interface for UserRepo

const userCollection = "user"

type UserRepository struct {
	db *mongo.Database
}

func NewUserRepo(db *mongo.Database) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) FindById(id string) models.User {
	filter := bson.D{{"_id", id}}

	var result models.User
	err := r.db.Collection(userCollection).FindOne(context.TODO(), filter).Decode(&result)

	// TODO: Handle error properly
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			panic(err)
		} else {
			panic(err)
		}
	}

	return result
}

func (r *UserRepository) Insert(user models.User) {
	_, err := r.db.Collection(userCollection).InsertOne(context.TODO(), user)

	// TODO: Handle error properly
	if err != nil {
		panic(err)
	}
}
