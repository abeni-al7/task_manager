package repositories

import (
	"context"
	"errors"
	"time"

	domain "github.com/abeni-al7/task_manager/Domain"
	infrastructure "github.com/abeni-al7/task_manager/Infrastructure"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(collection *mongo.Collection) *UserRepository {
	return &UserRepository{
		collection: collection,
	}
}

func (ur *UserRepository) FetchByUsername(username string) (domain.User, error) {
	var existingUser domain.User

	err := ur.collection.FindOne(context.TODO(), bson.D{{Key: "username", Value: username}}).Decode(&existingUser)

	if err != nil {
		return domain.User{}, errors.New("user does not exists")
	}
	return existingUser, nil
}

func (ur *UserRepository) CountUsers() (int, error) {
	userCount, err := ur.collection.CountDocuments(context.TODO(), bson.D{{}})
	if err != nil {
		return 0, errors.New("unable to register user")
	}
	return int(userCount), nil
}

func (ur *UserRepository) Register(user *domain.User) (domain.User, error) {
	user.ID = primitive.NewObjectID()

	_, err := ur.collection.InsertOne(context.TODO(), user)
	if err != nil {
		return domain.User{}, errors.New(err.Error())
	}

	return *user, nil
}

func (ur *UserRepository) Promote(user *domain.User) (domain.User, error) {
	var updatedUser domain.User

	filter := bson.D{{Key: "_id", Value: user.ID}}

	update := bson.D{{Key: "$set", Value: bson.D{
		{Key: "role", Value: "admin"},
	}}}

	_, err := ur.collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return domain.User{}, errors.New(err.Error())
	}
	
	err = ur.collection.FindOne(context.TODO(), filter).Decode(&updatedUser)
	if err != nil {
		return domain.User{}, errors.New("user not found")
	}
	return updatedUser, nil
}

func (ur *UserRepository) FetchAll() ([]domain.User, error) {
	var users []domain.User

	cur, err := ur.collection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		return []domain.User{}, errors.New("could not fetch users")
	}

	err = cur.All(context.TODO(), &users)
	if err != nil {
		return []domain.User{}, errors.New("could not fetch users")
	}

	cur.Close(context.TODO())

	return users, nil
}

func (ur *UserRepository) Fetch(idStr string) (domain.User, error) {
	var user domain.User

	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		return domain.User{}, errors.New("invalid id")
	}

	filter := bson.D{{Key: "_id", Value: id}}

	err = ur.collection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		return domain.User{}, errors.New("user not found")
	}

	return user, nil
}

func (ur *UserRepository) Update(idStr string, updatedUser domain.User) (domain.User, error) {
	var user domain.User

	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		return domain.User{}, errors.New("invalid id")
	}

	filter := bson.D{{Key: "_id", Value: id}}

	fields := bson.D{}
	if updatedUser.Email != "" {
		fields = append(fields, bson.E{Key: "email", Value: updatedUser.Email})
	}
	fields = append(fields, bson.E{Key: "updated_at", Value: time.Now()})

	update := bson.D{{Key: "$set", Value: fields}}

	_, err = ur.collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return domain.User{}, errors.New(err.Error())
	}
	
	err = ur.collection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		return domain.User{}, errors.New(err.Error())
	}
	return user, nil
}

func (ur *UserRepository) ChangePassword(idStr string, prevPassword string, newPassword string) error {
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		return errors.New("invalid id")
	}
	filter := bson.D{{Key: "_id", Value: id}}

	hashedPassword, err := new(infrastructure.Infrastructure).HashPassword(newPassword)
	if err != nil {
		return errors.New("system could not hash the password")
	}

	update := bson.D{{Key: "$set", Value: bson.D{
		{Key: "password", Value: string(hashedPassword)},
	}}}

	_, err = ur.collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return errors.New("system could not update user")
	}

	return nil
}

func (ur *UserRepository) Remove(idStr string) error {
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		return errors.New("invalid id")
	}

	filter := bson.D{{Key: "_id", Value: id}}

	_, err = ur.collection.DeleteOne(context.TODO(), filter)

	if err != nil {
		return errors.New("user not found")
	}

	return nil
}