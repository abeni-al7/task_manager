package controllers

import (
	"net/http"

	domain "github.com/abeni-al7/task_manager/Domain"
	usecases "github.com/abeni-al7/task_manager/Usecases"
	"github.com/gin-gonic/gin"
)

type UserInput struct {
	Username string `json:"username"`
	Email string `json:"email"`
	Password string `json:"password"`
}

type PasswordInput struct {
	PrevPassword string `json:"prev_password"`
	NewPassword string `json:"new_password"`
}

type UserController struct {
	UserUsecase usecases.UserUsecase
}


func (uc *UserController) Register(ctx *gin.Context) {
	var newUser UserInput
	
	if err := ctx.ShouldBindJSON(&newUser); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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
		return
	}
	ctx.JSON(http.StatusCreated, user)
}

func (uc *UserController) Login(ctx *gin.Context) {
	var user UserInput

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := uc.UserUsecase.Login(user.Username, user.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"token": token})
}

func (uc *UserController) Promote(ctx *gin.Context) {
	id := ctx.Param("id")

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
	id := ctx.Param("id")

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

	if err := ctx.ShouldBindJSON(&updatedUser); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid fields"})
		return
	}

	user, err := uc.UserUsecase.Update(idStr, updatedUser)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (uc *UserController) ChangePassword(ctx *gin.Context) {
	var passwordInput PasswordInput

	id := ctx.Param("id")

	if err := ctx.ShouldBindJSON(&passwordInput); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := uc.UserUsecase.ChangePassword(id, passwordInput.PrevPassword, passwordInput.NewPassword)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "password updated successfully"})
}

func (uc *UserController) Remove(ctx *gin.Context) {
	id := ctx.Param("id")

	err := uc.UserUsecase.Remove(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}