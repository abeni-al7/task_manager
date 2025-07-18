package controllers

import (
	"net/http"

	"github.com/abeni-al7/task_manager/data"
	"github.com/abeni-al7/task_manager/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


func GetUsersController(ctx *gin.Context) {
	users, err := data.GetUsersService()
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"users": users})
}

func GetUserController(ctx *gin.Context) {
	idStr := ctx.Param("id")

	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := data.GetUserService(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func UpdateUserController(ctx *gin.Context) {
	var updatedUser models.User

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

	user, err := data.UpdateUserService(id, updatedUser)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func ChangePasswordController(ctx *gin.Context) {
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

	err = data.ChangePasswordService(id, prevPassword, newPassword)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "password updated successfully"})
}

func PromoteUserController(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	user, err := data.PromoteUserService(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func RemoveUserController(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = data.RemoveUserService(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

func RegisterUserController(ctx *gin.Context) {
	var newUser models.User
	
	if err := ctx.ShouldBindJSON(&newUser); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user"})
		return
	}
	
	if newUser.FirstName == "" || newUser.LastName == "" || 
	newUser.Email == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Missing required fields"})
		return
	}
	
	user, err := data.RegisterUserService(newUser)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
	}
	ctx.JSON(http.StatusCreated, user)
}

func LoginUserController(ctx *gin.Context) {
	var body map[string]interface{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	email, ok := body["email"].(string)
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "email is required"})
		return
	}

	password, ok := body["password"].(string)
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "password is required"})
		return
	}

	token, err := data.LoginUserService(email, password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"token": token})
}