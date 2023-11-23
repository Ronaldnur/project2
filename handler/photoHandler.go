package handler

import (
	"project2/dto"
	"project2/entity"
	"project2/pkg/errs"
	"project2/pkg/helpers"
	"project2/service"

	"github.com/gin-gonic/gin"
)

type photoHandler struct {
	photoService service.PhotoService
}

func NewPhotoHandler(photoService service.PhotoService) photoHandler {
	return photoHandler{
		photoService: photoService,
	}
}

// Posting godoc
// @Tags photos
// @Summary Posting a new photo
// @Description Post a new photo
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param RequestBody body dto.NewPhotoRequest true "Photo posting request"
// @Success 201 {object} dto.NewPhotoResponse
// @Router /photos [post]
func (ph *photoHandler) Posting(ctx *gin.Context) {
	var newPhotoRequest dto.NewPhotoRequest

	if err := ctx.ShouldBindJSON(&newPhotoRequest); err != nil {
		errBindJson := errs.NewUnprocessibleEntityError("invalid request body")

		ctx.JSON(errBindJson.Status(), errBindJson)
		return
	}
	user := ctx.MustGet("userData").(entity.User)

	result, err := ph.photoService.PostPhoto(user.Id, newPhotoRequest)

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}
	ctx.JSON(result.StatusCode, result)
}

// GetPhoto godoc
// @Tags photos
// @Summary Get photos of the authenticated user
// @Description Get photos uploaded by the authenticated user
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Success 200 {object} dto.GetPhotoResponse
// @Router /photos [get]
func (ph *photoHandler) GetPhoto(ctx *gin.Context) {
	user := ctx.MustGet("userData").(entity.User)

	result, err := ph.photoService.GetPhotoUsers(user.Id)

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}
	ctx.JSON(result.StatusCode, result)
}

// UpdatePhoto godoc
// @Tags photos
// @Summary Update photo details
// @Description Update the details of the authenticated user's photo
// @Accept json
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param photoId path int true "Photo ID"
// @Param RequestBody body dto.NewPhotoRequest true "Photo update request"
// @Success 200 {object} dto.NewUpdateResponse
// @Router /photos/{photoId} [put]
func (ph *photoHandler) UpdatePhoto(ctx *gin.Context) {
	var updateRequest dto.NewPhotoRequest

	if err := ctx.ShouldBindJSON(&updateRequest); err != nil {
		errBindJson := errs.NewUnprocessibleEntityError("invalid request body")

		ctx.JSON(errBindJson.Status(), errBindJson)
		return
	}

	photoId, err := helpers.GetParamId(ctx, "photoId")

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	result, err := ph.photoService.PhotoUpdate(photoId, updateRequest)
	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	ctx.JSON(result.StatusCode, result)
}

// DeletePhoto godoc
// @Tags photos
// @Summary Delete a photo
// @Description Delete a photo uploaded by the authenticated user
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param photoId path int true "Photo ID"
// @Success 200 {object} dto.DeletePhotoResponse "Successfully deleted"
// @Router /photos/{photoId} [delete]
func (ph *photoHandler) DeletePhoto(ctx *gin.Context) {

	photoId, err := helpers.GetParamId(ctx, "photoId")

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	result, err := ph.photoService.PhotoDelete(photoId)
	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}
	ctx.JSON(result.StatusCode, result)

}
