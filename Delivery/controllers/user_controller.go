package controllers

import (
	"net/http"

	"github.com/abeni-al7/task_manager/Domain"
	"github.com/abeni-al7/task_manager/Usecases"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserController struct {
	UserUsecase usecases.UserUsecase
}


func (uc *UserController) Register(ctx *gin.Context) {
	var newUser domain.RegisterUserInput
	
	if err := ctx.ShouldBindJSON(&newUser); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	if newUser.Username == "" || newUser.Email == "" || newUser.Password == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Missing required fields"})
		return
	}

	userToRegister := domain.User{
		Username: newUser.Username,
		Password: newUser.Password,
		Email: newUser.Email,
	}
	
	user, err := uc.UserUsecase.Register(&userToRegister)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
	}
	ctx.JSON(http.StatusCreated, user)
}

func (uc *UserController) Login(ctx *gin.Context) {
	var body map[string]interface{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	username, ok := body["username"].(string)
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "username is required"})
		return
	}

	password, ok := body["password"].(string)
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "password is required"})
		return
	}

	token, err := uc.UserUsecase.Login(username, password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"token": token})
}

func (uc *UserController) Promote(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	user, err := uc.UserUsecase.Promote(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (uc *UserController) FetchAll(ctx *gin.Context) {
	users, err := uc.UserUsecase.FetchAll()
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"users": users})
}

func (uc *UserController) Fetch(ctx *gin.Context) {
	idStr := ctx.Param("id")

	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := uc.UserUsecase.Fetch(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (uc *UserController) Update(ctx *gin.Context) {
	var updatedUser domain.User

	idStr := ctx.Param("id")
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ctx.ShouldBindJSON(&updatedUser); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := uc.UserUsecase.Update(id, updatedUser)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (uc *UserController) ChangePassword(ctx *gin.Context) {
	var body map[string]interface{}

	idStr := ctx.Param("id")
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	prevPassword, ok := body["prev_password"].(string)
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "previous password missing"})
		return
	}

	newPassword, ok := body["new_password"].(string)
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "new password missing"})
		return
	}

	err = uc.UserUsecase.ChangePassword(id, prevPassword, newPassword)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "password updated successfully"})
}

func (uc *UserController) Remove(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = uc.UserUsecase.Remove(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}