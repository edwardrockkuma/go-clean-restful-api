package handler

import (
	services "edward/cleanarc/pkg/usecase/interface"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

type UserHandler struct {
	userService services.UserService
}

type Response struct {
	Id   uint   `copier:"must"`
	Name string `copier:"must"`
	Mail string `copier:"must"`
}

func NewUserHandler(usecase services.UserService) *UserHandler {
	return &UserHandler{
		userService: usecase,
	}
}

func (u *UserHandler) FindAll(c *gin.Context) {
	users, err := u.userService.FindAll(c.Request.Context())

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		response := []Response{}
		copier.Copy(&response, &users)

		c.JSON(http.StatusOK, response)
	}
}

func (u *UserHandler) FindById(c *gin.Context) {
	paramsId := c.Param("id")
	id, err := strconv.Atoi(paramsId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "id format error",
		})
		return
	}

	user, err := u.userService.FindById(c.Request.Context(), uint(id))

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		response := Response{}
		copier.Copy(&response, &user)

		c.JSON(http.StatusOK, response)
	}

}
