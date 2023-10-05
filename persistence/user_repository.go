package persistence

import (
	"context"
	"errors"
	"github.com/sokangho-wex/ps-tag-onboarding-go/models"
	"github.com/sokangho-wex/ps-tag-onboarding-go/models/errs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
			return models.User{}, errs.NewNotFoundError()
		} else {
			return models.User{}, errs.NewUnexpectedError(err)
		}
	}

	return result, nil
}

func (r *UserRepository) SaveUser(user models.User) error {
	filter := bson.D{{"_id", user.ID}}
	update := bson.D{{"$set", bson.D{{"first_name", user.FirstName}, {"last_name", user.LastName}, {"email", user.Email}, {"age", user.Age}}}}
	opts := options.Update().SetUpsert(true)

	_, err := r.db.Collection(userCollection).UpdateOne(context.TODO(), filter, update, opts)

	if err != nil {
		return errs.NewUnexpectedError(err)
	}

	return nil
}

func (r *UserRepository) ExistsByFirstNameAndLastName(firstName, lastName string) (bool, error) {
	filter := bson.D{{"first_name", firstName}, {"last_name", lastName}}

	count, err := r.db.Collection(userCollection).CountDocuments(context.TODO(), filter)
	if err != nil {
		return false, errs.NewUnexpectedError(err)
	}

	if count > 0 {
		return true, nil
	}

	return false, nil
}
