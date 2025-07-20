package data

import (
	"context"
	"errors"
	"os"
	"time"

	"github.com/abeni-al7/task_manager/models"
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func GetUsersService() ([]models.User, error) {
	var users []models.User

	cur, err := UserCollection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		return []models.User{}, errors.New("cannot retrieve users")
	}

	for cur.Next(context.TODO()) {
		var user models.User
		err := cur.Decode(&user)
		if err != nil {
			return []models.User{}, errors.New("cannot retrieve users")
		}
		users = append(users, user)
	}

	if err := cur.Err(); err != nil {
		return []models.User{}, errors.New("cannot retrieve users")
	}

	cur.Close(context.TODO())

	return users, nil
}

func GetUserService(id primitive.ObjectID) (models.User, error) {
	var user models.User

	filter := bson.D{{Key: "_id", Value: id}}

	err := UserCollection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		return models.User{}, errors.New("user not found")
	}

	return user, nil
}

func UpdateUserService(id primitive.ObjectID, updatedUser models.User) (models.User, error) {
	var user models.User
	filter := bson.D{{Key: "_id", Value: id}}

	fields := bson.D{}
	if updatedUser.Username != "" {
		fields = append(fields, bson.E{Key: "username", Value: updatedUser.Username})
	}
	if updatedUser.Email != "" {
		fields = append(fields, bson.E{Key: "email", Value: updatedUser.Email})
	}
	fields = append(fields, bson.E{Key: "updated_at", Value: time.Now()})

	update := bson.D{{Key: "$set", Value: fields}}

	_, err := UserCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return models.User{}, errors.New(err.Error())
	}
	
	err = UserCollection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		return models.User{}, errors.New(err.Error())
	}
	return user, nil
}

func ChangePasswordService(id primitive.ObjectID, prevPassword string, newPassword string) error {
	var user models.User

	filter := bson.D{{Key: "_id", Value: id}}

	err := UserCollection.FindOne(context.TODO(), filter).Decode(&user)
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

	_, err = UserCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return errors.New(err.Error())
	}

	return nil
}

func PromoteUserService(id primitive.ObjectID) (models.User, error) {
	var user models.User
	filter := bson.D{{Key: "_id", Value: id}}

	update := bson.D{{Key: "$set", Value: bson.D{
		{Key: "role", Value: "admin"},
	}}}

	_, err := UserCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return models.User{}, errors.New(err.Error())
	}
	
	err = UserCollection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		return models.User{}, errors.New("user not found")
	}
	return user, nil
}

func RemoveUserService(id primitive.ObjectID) error {
	var user models.User

	filter := bson.D{{Key: "_id", Value: id}}

	err := UserCollection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		return errors.New("user not found")
	}

	if user.Role == "admin" {
		return errors.New("admin cannot be deleted")
	}
	_, err = UserCollection.DeleteOne(context.TODO(), filter)

	if err != nil {
		return errors.New("user not found")
	}

	return nil
}

func RegisterUserService(newUser models.User) (models.User, error) {
	var existingUser models.User

	err := UserCollection.FindOne(context.TODO(), bson.D{{Key: "email", Value: newUser.Email}}).Decode(&existingUser)

	if err == nil {
		return models.User{}, errors.New("user already exists")
	}

	newUser.ID = primitive.NewObjectID()
	newUser.CreatedAt = time.Now()
	newUser.UpdatedAt = time.Now()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		return models.User{}, errors.New(err.Error())
	}
	newUser.Password = string(hashedPassword)

	userCount, err := UserCollection.CountDocuments(context.TODO(), bson.D{{}})
	if err != nil {
		return models.User{}, errors.New("unable to register user")
	}

	switch userCount{
	case 0:
		newUser.Role = "admin"
	default:
		newUser.Role = "regular"
	}

	_, err = UserCollection.InsertOne(context.TODO(), newUser)
	if err != nil {
		return models.User{}, errors.New(err.Error())
	}
	return newUser, nil
}

func LoginUserService(email string, password string) (string, error) {
	var user models.User

	err := UserCollection.FindOne(context.TODO(), bson.D{{Key: "email", Value: email}}).Decode(&user)
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
		return "", errors.New("invalid email or password")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"email": user.Email,
		"username": user.Username,
		"role": user.Role,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	jwtToken, err := token.SignedString(os.Getenv("JWT_SECRET"))
	if err != nil {
		return "", errors.New("unable to generate token")
	}

	return jwtToken, nil
}