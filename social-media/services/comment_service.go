package services

import (
	"social-media/models"
	"social-media/models/sharedmodels"
	"social-media/repositories"
	"social-media/utils"
)

func CreateComment(commentRequest *models.CommentRequestViewModel) (*models.CommentResponseViewModel, error) {
	comment := models.MapCommentRequestViewModelToCommentDbModel(commentRequest)
	c, err := repositories.CreateComment(comment)
	if err != nil {
		return nil, err
	}
	author, err := GetCachedUserInfoByUsername(commentRequest.Username)
	commentResponseVm := c.MapCommentEntityAndCommentResponseViewModel(author)
	return commentResponseVm, err
}

func GetCommentsByPostId(postId int64) ([]models.CommentResponseViewModel, error) {
	comments, err := repositories.GetCommentsByPostId(postId)
	if err != nil {
		return nil, err
	}

	var commentVMs []models.CommentResponseViewModel
	for _, c := range comments {
		vm := models.CommentResponseViewModel{
			Id:        int64(c.Id),
			Content:   c.Content,
			CreatedAt: c.CreatedAt,
			Author: sharedmodels.UserBaseViewModel{
				UserId:    c.UserId,
				FirstName: c.User.FirstName,
				LastName:  c.User.LastName,
				Name:      c.User.Name,
				Birthday:  c.User.Birthday,
				Email:     c.User.Email,
				Avatar:    c.User.Avatar,
			},
		}
		commentVMs = append(commentVMs, vm)
	}

	return commentVMs, nil
}

func UpdateComment(commentRequest *models.CommentRequestViewModel) (*models.CommentResponseViewModel, error) {
	author, err := GetCachedUserInfoByUsername(commentRequest.Username)
	if err != nil {
		return nil, err
	}

	currentComment, err := repositories.GetCommentById(commentRequest.CommentId)
	if err != nil {
		return nil, err
	}
	if currentComment.UserId != commentRequest.UserId {
		return nil, utils.ErrCannotEditComment
	}
	if currentComment.PostId != commentRequest.PostId {
		return nil, utils.ErrCommentNotInPost
	}

	updatedComment := models.MapCommentRequestViewModelToCommentDbModel(commentRequest)
	updatedComment.CreatedAt = currentComment.CreatedAt
	updatedComment, err = repositories.UpdateComment(updatedComment)
	if err != nil {
		return nil, err
	}

	commentResponse := updatedComment.MapCommentEntityAndCommentResponseViewModel(author)
	return commentResponse, err
}
