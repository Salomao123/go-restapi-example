package controller

import (
	"net/http"

	"github.com/Salomao123/go-restapi-example/models"
	"github.com/Salomao123/go-restapi-example/service"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController(userService *service.UserService) *UserController {
	return &UserController{userService: userService}
}

func (c *UserController) CreateUser(g *gin.Context) {
	var user models.User

	err := g.ShouldBindJSON(&user)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = c.userService.Create(g.Request.Context(), user)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusCreated, gin.H{"message": "User created successfully!", "user": user})

}

func (c *UserController) FindAllUsers(g *gin.Context) {
	users, err := c.userService.GetAll(g)
	if err != nil {
		return
	}

	g.JSON(http.StatusOK, gin.H{"users": users})
}
