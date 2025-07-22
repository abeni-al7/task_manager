package data

import (
	"context"
	"errors"
	"time"

	"github.com/abeni-al7/task_manager/Domain"
	"github.com/abeni-al7/task_manager/Infrastructure"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func GetUsersService() ([]domain.User, error) {
	var users []domain.User

	cur, err := domain.UserCollection.Find(context.TODO(), bson.D{{}})
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

func GetUserService(id primitive.ObjectID) (domain.User, error) {
	var user domain.User

	filter := bson.D{{Key: "_id", Value: id}}

	err := domain.UserCollection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		return domain.User{}, errors.New("user not found")
	}

	return user, nil
}

func UpdateUserService(id primitive.ObjectID, updatedUser domain.User) (domain.User, error) {
	var user domain.User
	filter := bson.D{{Key: "_id", Value: id}}

	fields := bson.D{}
	if updatedUser.Email != "" {
		fields = append(fields, bson.E{Key: "email", Value: updatedUser.Email})
	}
	fields = append(fields, bson.E{Key: "updated_at", Value: time.Now()})

	update := bson.D{{Key: "$set", Value: fields}}

	_, err := domain.UserCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return domain.User{}, errors.New(err.Error())
	}
	
	err = domain.UserCollection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		return domain.User{}, errors.New(err.Error())
	}
	return user, nil
}

func ChangePasswordService(id primitive.ObjectID, prevPassword string, newPassword string) error {
	var user domain.User

	filter := bson.D{{Key: "_id", Value: id}}

	err := domain.UserCollection.FindOne(context.TODO(), filter).Decode(&user)
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

	_, err = domain.UserCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return errors.New(err.Error())
	}

	return nil
}

func PromoteUserService(id primitive.ObjectID) (domain.User, error) {
	var user domain.User
	filter := bson.D{{Key: "_id", Value: id}}

	update := bson.D{{Key: "$set", Value: bson.D{
		{Key: "role", Value: "admin"},
	}}}

	_, err := domain.UserCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return domain.User{}, errors.New(err.Error())
	}
	
	err = domain.UserCollection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		return domain.User{}, errors.New("user not found")
	}
	return user, nil
}

func RemoveUserService(id primitive.ObjectID) error {
	var user domain.User

	filter := bson.D{{Key: "_id", Value: id}}

	err := domain.UserCollection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		return errors.New("user not found")
	}

	if user.Role == "admin" {
		return errors.New("admin cannot be deleted")
	}
	_, err = domain.UserCollection.DeleteOne(context.TODO(), filter)

	if err != nil {
		return errors.New("user not found")
	}

	return nil
}

func RegisterUserService(newUser domain.User) (domain.User, error) {
	var existingUser domain.User

	err := domain.UserCollection.FindOne(context.TODO(), bson.D{{Key: "username", Value: newUser.Username}}).Decode(&existingUser)

	if err == nil {
		return domain.User{}, errors.New("user already exists")
	}

	newUser.ID = primitive.NewObjectID()
	newUser.CreatedAt = time.Now()
	newUser.UpdatedAt = time.Now()

	hashedPassword, err := infrastructure.HashPassword(newUser.Password)
	if err != nil {
		return domain.User{}, errors.New(err.Error())
	}
	newUser.Password = hashedPassword

	userCount, err := domain.UserCollection.CountDocuments(context.TODO(), bson.D{{}})
	if err != nil {
		return domain.User{}, errors.New("unable to register user")
	}

	if userCount == 0 {
		newUser.Role = "admin"
	} else {
		newUser.Role = "regular"
	}

	_, err = domain.UserCollection.InsertOne(context.TODO(), newUser)
	if err != nil {
		return domain.User{}, errors.New(err.Error())
	}
	return newUser, nil
}

func LoginUserService(username string, password string) (string, error) {
	var user domain.User

	err := domain.UserCollection.FindOne(context.TODO(), bson.D{{Key: "username", Value: username}}).Decode(&user)
	if err != nil {
		return "", errors.New("invalid username or password")
	}

	jwtToken, err := infrastructure.GenerateJwtToken(&user, password)
	if err != nil {
		return "", err
	}

	return jwtToken, nil
}