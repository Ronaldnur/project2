package handler

import (
	"project2/dto"
	"project2/entity"
	"project2/pkg/errs"
	"project2/pkg/helpers"
	"project2/service"

	"github.com/gin-gonic/gin"
)

type commentHandler struct {
	commentService service.CommentService
}

func NewCommentHandler(commentService service.CommentService) commentHandler {
	return commentHandler{
		commentService: commentService,
	}
}

// MakeComment godoc
// @Tags comments
// @Summary Make a new comment
// @Description Make a new comment with the provided details
// @Accept json
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param RequestBody body dto.NewCommentRequest true "Comment create request"
// @Success 201 {object} dto.CommentResponse
// @Router /comments [post]
func (ch *commentHandler) MakeComment(ctx *gin.Context) {

	var CommentRequest dto.NewCommentRequest

	if err := ctx.ShouldBindJSON(&CommentRequest); err != nil {
		errBindJson := errs.NewUnprocessibleEntityError("invalid request body")

		ctx.JSON(errBindJson.Status(), errBindJson)
		return
	}
	comment := ctx.MustGet("userData").(entity.User)

	result, err := ch.commentService.CreateNewComment(CommentRequest, comment.Id)

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}
	ctx.JSON(result.StatusCode, result)
}

// GetComments godoc
// @Tags comments
// @Summary Get comments of the authenticated user
// @Description Get comments made by the authenticated user
// @Accept json
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Success 200 {object} dto.GetCommentsResponse
// @Router /comments [get]
func (ch *commentHandler) GetComments(ctx *gin.Context) {
	user := ctx.MustGet("userData").(entity.User)

	result, err := ch.commentService.GetComments(user.Id)

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}
	ctx.JSON(result.StatusCode, result)
}

// UpdateComment godoc
// @Tags comments
// @Summary Update comment details
// @Description Update the details the user's comment
// @Accept json
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param commentId path int true "Comment ID"
// @Param RequestBody body dto.MakeCommentUpdate true "Comment update request"
// @Success 200 {object} dto.UpdateResponse
// @Router /comments/{commentId} [put]
func (ch *commentHandler) UpdateComment(ctx *gin.Context) {
	var updateRequest dto.MakeCommentUpdate

	if err := ctx.ShouldBindJSON(&updateRequest); err != nil {
		errBindJson := errs.NewUnprocessibleEntityError("invalid request body")

		ctx.JSON(errBindJson.Status(), errBindJson)
		return
	}
	commentId, err := helpers.GetParamId(ctx, "commentId")
	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	result, err := ch.commentService.UpdateComment(commentId, updateRequest)
	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}
	ctx.JSON(result.StatusCode, result)
}

// DeleteComment godoc
// @Tags comments
// @Summary Delete a comment
// @Description Delete a comment user
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param commentId path int true "Comment ID"
// @Success 200 {object} dto.DeleteCommentResponse "Successfully deleted"
// @Router /comments/{commentId} [delete]
func (ch *commentHandler) DeleteComment(ctx *gin.Context) {
	commentId, err := helpers.GetParamId(ctx, "commentId")

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	result, err := ch.commentService.DeleteComment(commentId)
	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}
	ctx.JSON(result.StatusCode, result)
}
