package mongo

import (
	"context"
	"errors"
	"github.com/sokangho-wex/ps-tag-onboarding-go/internal/domain/onboardingerrors"
	"github.com/sokangho-wex/ps-tag-onboarding-go/internal/domain/users"
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

func (r *UserRepository) FindByID(ctx context.Context, id string) (users.User, error) {
	filter := bson.D{{"_id", id}}

	var result users.User
	err := r.db.Collection(userCollection).FindOne(ctx, filter).Decode(&result)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return users.User{}, onboardingerrors.NewNotFoundError()
		} else {
			return users.User{}, onboardingerrors.NewUnexpectedError(err)
		}
	}

	return result, nil
}

func (r *UserRepository) SaveUser(ctx context.Context, user users.User) error {
	filter := bson.D{{"_id", user.ID}}
	update := bson.D{{"$set", bson.D{{"firstname", user.FirstName}, {"lastname", user.LastName}, {"email", user.Email}, {"age", user.Age}}}}
	opts := options.Update().SetUpsert(true)

	_, err := r.db.Collection(userCollection).UpdateOne(ctx, filter, update, opts)

	if err != nil {
		return onboardingerrors.NewUnexpectedError(err)
	}

	return nil
}

func (r *UserRepository) ExistsByFirstNameAndLastName(ctx context.Context, firstName, lastName string) (bool, error) {
	filter := bson.D{{"firstname", firstName}, {"lastname", lastName}}

	count, err := r.db.Collection(userCollection).CountDocuments(ctx, filter)
	if err != nil {
		return false, onboardingerrors.NewUnexpectedError(err)
	}

	if count > 0 {
		return true, nil
	}

	return false, nil
}
