package handler

import (
	"project2/dto"
	"project2/entity"
	"project2/pkg/errs"
	"project2/service"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) userHandler {
	return userHandler{
		userService: userService,
	}
}

// Register godoc
// @Tags users
// @Summary Register a new user
// @Description Register a new user with the provided details
// @Accept json
// @Produce json
// @Param RequestBody body dto.NewUserRequest true "User registration request"
// @Success 201 {object} dto.NewUserResponse
// @Router /users/register [post]
func (u *userHandler) Register(ctx *gin.Context) {
	var newUserRequsest dto.NewUserRequest

	if err := ctx.ShouldBindJSON(&newUserRequsest); err != nil {
		errBindJson := errs.NewUnprocessibleEntityError("invalid request body")

		ctx.JSON(errBindJson.Status(), errBindJson)
		return
	}

	result, err := u.userService.CreateNewUser(newUserRequsest)

	if err != nil {
		ctx.JSON(err.Status(), err)
		return
	}
	ctx.JSON(result.StatusCode, result)

}

// Login godoc
// @Tags users
// @Summary Log in a user
// @Description Log in a user with the provided credentials
// @Accept json
// @Produce json
// @Param RequestBody body dto.NewUserLogin true "User login request"
// @Success 200 {object} dto.LoginResponse
// @Router /users/login [post]
func (u *userHandler) Login(ctx *gin.Context) {
	var newUserRequsest dto.NewUserLogin

	if err := ctx.ShouldBindJSON(&newUserRequsest); err != nil {
		errBindJson := errs.NewUnprocessibleEntityError("invalid request body")

		ctx.JSON(errBindJson.Status(), errBindJson)
		return
	}
	result, err := u.userService.Login(newUserRequsest)

	if err != nil {
		ctx.JSON(err.Status(), err)
		return
	}

	ctx.JSON(result.StatusCode, result)
}

// UpdateUser godoc
// @Tags users
// @Summary Update user details
// @Description Update the details of the authenticated user
// @Accept json
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param RequestBody body dto.NewUserUpdate true "User update request"
// @Success 200 {object} dto.UserUpdateResponse
// @Router /users [put]
func (u *userHandler) UpdateUser(ctx *gin.Context) {
	var newUserUpdate dto.NewUserUpdate

	if err := ctx.ShouldBindJSON(&newUserUpdate); err != nil {
		errBindJson := errs.NewUnprocessibleEntityError("invalid request body")

		ctx.JSON(errBindJson.Status(), errBindJson)
		return
	}
	user := ctx.MustGet("userData").(entity.User)

	result, err := u.userService.UpdateUser(user.Id, newUserUpdate)

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	ctx.JSON(result.StatusCode, result)
}

// Delete godoc
// @Tags users
// @Summary Delete the authenticated user
// @Description Delete the account of the authenticated user
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Produce json
// @Success 200 {object} dto.DeleteResponse "Successfully deleted"
// @Router /users [delete]
func (u *userHandler) Delete(ctx *gin.Context) {
	user := ctx.MustGet("userData").(entity.User)
	result, err := u.userService.DeleteUser(user.Id)

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	ctx.JSON(result.StatusCode, result)

}
