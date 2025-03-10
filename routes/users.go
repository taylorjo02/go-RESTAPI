package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/taylorjo02/go-RESTAPI/models"
)

func signup(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}
	user.ID = 1
	if err = user.Save(); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create user."})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "User created!", "user": user})

	

}