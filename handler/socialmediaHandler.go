package handler

import (
	"project2/dto"
	"project2/entity"
	"project2/pkg/errs"
	"project2/pkg/helpers"
	"project2/service"

	"github.com/gin-gonic/gin"
)

type socialMediaHandler struct {
	socialMediaService service.SocialMediaService
}

func NewSocialMediaHandler(socialMediaService service.SocialMediaService) socialMediaHandler {
	return socialMediaHandler{
		socialMediaService: socialMediaService,
	}
}

// PostSocialMedia godoc
// @Tags socialMedia
// @Summary Post a new social media
// @Description Post a new social media
// @Accept json
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param RequestBody body dto.NewSocialMediaRequest true "Social media create request"
// @Success 201 {object} dto.SocialMediaResponse
// @Router /socialmedias [post]
func (smh *socialMediaHandler) PostSocialMedia(ctx *gin.Context) {
	var newSocialMedia dto.NewSocialMediaRequest

	if err := ctx.ShouldBindJSON(&newSocialMedia); err != nil {
		errBindJson := errs.NewUnprocessibleEntityError("invalid request body")

		ctx.JSON(errBindJson.Status(), errBindJson)
		return
	}

	user := ctx.MustGet("userData").(entity.User)

	result, err := smh.socialMediaService.MakeSocialMedia(user.Id, newSocialMedia)
	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	ctx.JSON(result.StatusCode, result)
}

// GetSocialMedia godoc
// @Tags socialMedia
// @Summary Get social media from the user
// @Description Get social media posted by the user
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Success 200 {object} dto.GetSocialMediaResponse
// @Router /socialmedias [get]
func (smh *socialMediaHandler) GetSocialMedia(ctx *gin.Context) {

	user := ctx.MustGet("userData").(entity.User)

	result, err := smh.socialMediaService.GetSocialMedia(user.Id)

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	ctx.JSON(result.StatusCode, result)
}

// UpdateSocialMedia godoc
// @Tags socialMedia
// @Summary Update social media details
// @Description Update the details of the authenticated user's social media
// @Accept json
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param socialMediaId path int true "Social Media ID"
// @Param RequestBody body dto.NewSocialMediaRequest true "Social media update request"
// @Success 200 {object} dto.UpdateSocialMediaResponse
// @Router /socialmedias/{socialMediaId} [put]
func (smh *socialMediaHandler) UpdateSocialMedia(ctx *gin.Context) {
	var newUpdate dto.NewSocialMediaRequest

	if err := ctx.ShouldBindJSON(&newUpdate); err != nil {
		errBindJson := errs.NewUnprocessibleEntityError("invalid request body")

		ctx.JSON(errBindJson.Status(), errBindJson)
		return
	}

	socialMediaId, err := helpers.GetParamId(ctx, "socialMediaId")

	if err != nil {

		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	result, err := smh.socialMediaService.UpdateSocialMedia(socialMediaId, newUpdate)

	if err != nil {

		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	ctx.JSON(result.StatusCode, result)
}

// DeleteSocialMedia godoc
// @Tags socialMedia
// @Summary Delete a social media
// @Description Delete a social media posted by the user
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param socialMediaId path int true "Social Media ID"
// @Success 200 {object} dto.DeleteResponseSocialMedia
// @Router /socialmedias/{socialMediaId} [delete]
func (smh *socialMediaHandler) DeleteSocialMedia(ctx *gin.Context) {
	socialMediaId, err := helpers.GetParamId(ctx, "socialMediaId")
	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}
	result, err := smh.socialMediaService.DeleteSocialMedia(socialMediaId)
	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	ctx.JSON(result.StatusCode, result)
}
