package service

import (
	"fmt"
	"project2/entity"
	"project2/pkg/errs"
	"project2/pkg/helpers"
	"project2/repository/comments_repository"
	"project2/repository/photo_repository"
	"project2/repository/socialmedia_repository"
	"project2/repository/user_repository"

	"github.com/gin-gonic/gin"
)

type Authservice interface {
	Authentitaction() gin.HandlerFunc
	AuthorizationPhoto() gin.HandlerFunc
	AuthorizationComment() gin.HandlerFunc
	AuthorizationSocialMedia() gin.HandlerFunc
}

type authService struct {
	userRepo    user_repository.Repository
	photoRepo   photo_repository.Repository
	commentRepo comments_repository.Repository
	socialRepo  socialmedia_repository.Repository
}

func NewAuthService(userRepo user_repository.Repository, photoRepo photo_repository.Repository, commentRepo comments_repository.Repository, socialRepo socialmedia_repository.Repository) Authservice {
	return &authService{
		userRepo:    userRepo,
		photoRepo:   photoRepo,
		commentRepo: commentRepo,
		socialRepo:  socialRepo,
	}
}

func (a *authService) AuthorizationPhoto() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user := ctx.MustGet("userData").(entity.User)

		photoId, err := helpers.GetParamId(ctx, "photoId")

		if err != nil {
			ctx.AbortWithStatusJSON(err.Status(), err)
			return
		}
		photo, err := a.photoRepo.GetPhotoById(photoId)

		if err != nil {
			ctx.AbortWithStatusJSON(err.Status(), err)
			return
		}
		if photo.User_id != user.Id {
			unauthorizedErr := errs.NewUnauthorizedError("you are not authorized to modify the photo data")
			ctx.AbortWithStatusJSON(unauthorizedErr.Status(), unauthorizedErr)
			return
		}
		ctx.Next()
	}
}
func (a *authService) AuthorizationComment() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user := ctx.MustGet("userData").(entity.User)

		commentId, err := helpers.GetParamId(ctx, "commentId")

		if err != nil {
			ctx.AbortWithStatusJSON(err.Status(), err)
			return
		}
		comment, err := a.commentRepo.GetCommentById(commentId)

		if err != nil {
			ctx.AbortWithStatusJSON(err.Status(), err)
			return
		}
		if comment.User_id != user.Id {
			unauthorizedErr := errs.NewUnauthorizedError("you are not authorized to modify the comment data")
			ctx.AbortWithStatusJSON(unauthorizedErr.Status(), unauthorizedErr)
			return
		}
		ctx.Next()
	}
}
func (a *authService) AuthorizationSocialMedia() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user := ctx.MustGet("userData").(entity.User)

		socialMediaId, err := helpers.GetParamId(ctx, "socialMediaId")
		fmt.Println(err)
		if err != nil {
			fmt.Println(err)
			ctx.AbortWithStatusJSON(err.Status(), err)
			return
		}
		socialMedia, err := a.socialRepo.GetSocialMediaById(socialMediaId)
		fmt.Println(err)
		if err != nil {
			fmt.Println(err)
			ctx.AbortWithStatusJSON(err.Status(), err)
			return
		}
		if socialMedia.User_id != user.Id {
			unauthorizedErr := errs.NewUnauthorizedError("you are not authorized to modify the social media data")
			ctx.AbortWithStatusJSON(unauthorizedErr.Status(), unauthorizedErr)
			return
		}
		ctx.Next()
	}
}
func (a *authService) Authentitaction() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var invalidTokenErr = errs.NewUnauthenticatedError("invalid token")

		bearerToken := ctx.GetHeader("Authorization")

		var user entity.User

		err := user.ValidateToken(bearerToken)

		if err != nil {
			ctx.AbortWithStatusJSON(err.Status(), err)
			return
		}
		result, err := a.userRepo.GetUserByEmail(user.Email)

		if err != nil {
			ctx.AbortWithStatusJSON(invalidTokenErr.Status(), invalidTokenErr)
			return
		}

		_ = result

		ctx.Set("userData", user)

		ctx.Next()
	}
}
