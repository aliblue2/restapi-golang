package router

import (
	"net/http"

	"azno-space.com/azno/models"
	"azno-space.com/azno/utils"
	"github.com/gin-gonic/gin"
)

func SignupUserHandler(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "an error occured during to binding"})
		return
	}

	hashedPassword, err := utils.HashedPasword(user.Password)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "cant hashing password"})
		return
	}

	id, err := models.SignupUser(user.Email, user.Name, hashedPassword)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "cant create user account. email has already taken.!"})
		return
	}

	user.Id = id
	context.JSON(http.StatusCreated, gin.H{"message": "Account successfully created.!"})

}

func LoginUserHandler(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "cant bind to user obj "})
		return
	}

	hashedPassword, err := user.ValidateCreadintials(user.Email)

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "email or password id invalid"})
		return
	}

	result := utils.CompareHashedPass(hashedPassword, user.Password)

	if !result {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "password is invalid.!"})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.Id)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "cant create token"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"token": token})

}
