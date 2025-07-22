package repositories

import (
	"context"
	"errors"
	"time"

	"github.com/abeni-al7/task_manager/Domain"
	infrastructure "github.com/abeni-al7/task_manager/Infrastructure"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type UserRepositoryInterface interface {
	Register(user *domain.User) (domain.User, error)
	Login(username string, password string) (string, error)
	Promote(id primitive.ObjectID) (domain.User, error)
	FetchAll() ([]domain.User, error)
	Fetch(id primitive.ObjectID) (domain.User, error)
	Update(id primitive.ObjectID, updatedUser domain.User) (domain.User, error)
	ChangePassword(id primitive.ObjectID, prevPassword string, newPassword string) error
	Remove(id primitive.ObjectID) error
}

type UserRepository struct {
	database mongo.Database
	collection string
}

func (ur *UserRepository) Register(user *domain.User) (domain.User, error) {
	var existingUser domain.User

	err := ur.database.Collection(ur.collection).FindOne(context.TODO(), bson.D{{Key: "username", Value: user.Username}}).Decode(&existingUser)

	if err == nil {
		return domain.User{}, errors.New("user already exists")
	}

	user.ID = primitive.NewObjectID()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	hashedPassword, err := infrastructure.HashPassword(user.Password)
	if err != nil {
		return domain.User{}, errors.New(err.Error())
	}
	user.Password = hashedPassword

	userCount, err := ur.database.Collection(ur.collection).CountDocuments(context.TODO(), bson.D{{}})
	if err != nil {
		return domain.User{}, errors.New("unable to register user")
	}

	if userCount == 0 {
		user.Role = "admin"
	} else {
		user.Role = "regular"
	}

	_, err = ur.database.Collection(ur.collection).InsertOne(context.TODO(), user)
	if err != nil {
		return domain.User{}, errors.New(err.Error())
	}
	return *user, nil
}

func (ur *UserRepository) Login(username string, password string) (string, error) {
	var user domain.User

	err := ur.database.Collection(ur.collection).FindOne(context.TODO(), bson.D{{Key: "username", Value: username}}).Decode(&user)
	if err != nil {
		return "", errors.New("invalid username or password")
	}

	jwtToken, err := infrastructure.GenerateJwtToken(&user, password)
	if err != nil {
		return "", err
	}

	return jwtToken, nil
}

func (ur *UserRepository) Promote(id primitive.ObjectID) (domain.User, error) {
	var user domain.User
	filter := bson.D{{Key: "_id", Value: id}}

	update := bson.D{{Key: "$set", Value: bson.D{
		{Key: "role", Value: "admin"},
	}}}

	_, err := ur.database.Collection(ur.collection).UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return domain.User{}, errors.New(err.Error())
	}
	
	err = ur.database.Collection(ur.collection).FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		return domain.User{}, errors.New("user not found")
	}
	return user, nil
}

func (ur *UserRepository) FetchAll() ([]domain.User, error) {
	var users []domain.User

	cur, err := ur.database.Collection(ur.collection).Find(context.TODO(), bson.D{{}})
	if err != nil {
		return []domain.User{}, errors.New("cannot retrieve users")
	}

	for cur.Next(context.TODO()) {
		var user domain.User
		err := cur.Decode(&user)
		if err != nil {
			return []domain.User{}, errors.New("cannot retrieve users")
		}
		users = append(users, user)
	}

	if err := cur.Err(); err != nil {
		return []domain.User{}, errors.New("cannot retrieve users")
	}

	cur.Close(context.TODO())

	return users, nil
}

func (ur *UserRepository) Fetch(id primitive.ObjectID) (domain.User, error) {
	var user domain.User

	filter := bson.D{{Key: "_id", Value: id}}

	err := ur.database.Collection(ur.collection).FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		return domain.User{}, errors.New("user not found")
	}

	return user, nil
}

func (ur *UserRepository) Update(id primitive.ObjectID, updatedUser domain.User) (domain.User, error) {
	var user domain.User
	filter := bson.D{{Key: "_id", Value: id}}

	fields := bson.D{}
	if updatedUser.Email != "" {
		fields = append(fields, bson.E{Key: "email", Value: updatedUser.Email})
	}
	fields = append(fields, bson.E{Key: "updated_at", Value: time.Now()})

	update := bson.D{{Key: "$set", Value: fields}}

	_, err := ur.database.Collection(ur.collection).UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return domain.User{}, errors.New(err.Error())
	}
	
	err = ur.database.Collection(ur.collection).FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		return domain.User{}, errors.New(err.Error())
	}
	return user, nil
}

func (ur *UserRepository) ChangePassword(id primitive.ObjectID, prevPassword string, newPassword string) error {
	var user domain.User

	filter := bson.D{{Key: "_id", Value: id}}

	err := ur.database.Collection(ur.collection).FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		return errors.New(err.Error())
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(prevPassword)) != nil {
		return errors.New("incorrect password")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return errors.New(err.Error())
	}

	update := bson.D{{Key: "$set", Value: bson.D{
		{Key: "password", Value: string(hashedPassword)},
	}}}

	_, err = ur.database.Collection(ur.collection).UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return errors.New(err.Error())
	}

	return nil
}

func (ur *UserRepository) Remove(id primitive.ObjectID) error {
	var user domain.User

	filter := bson.D{{Key: "_id", Value: id}}

	err := ur.database.Collection(ur.collection).FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		return errors.New("user not found")
	}

	if user.Role == "admin" {
		return errors.New("admin cannot be deleted")
	}
	_, err = ur.database.Collection(ur.collection).DeleteOne(context.TODO(), filter)

	if err != nil {
		return errors.New("user not found")
	}

	return nil
}