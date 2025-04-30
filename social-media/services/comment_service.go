package services

import (
	"social-media/models"
	"social-media/models/sharedmodels"
	"social-media/repositories"
)

func CreateComment(comment *models.Comment) (*models.CommentResponseViewModel, error) {
	c, err := repositories.CreateComment(comment)
	if err != nil {
		return nil, err
	}

	commentResponseVm := c.CreateMappingCommentEntityAndCommentResponseViewModel()
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
			Author: sharedmodels.UserResponseViewModel{
				Name:     c.User.Name,
				Username: c.User.Username,
			},
		}
		commentVMs = append(commentVMs, vm)
	}

	return commentVMs, nil
}
