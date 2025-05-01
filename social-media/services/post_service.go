package services

import (
	"social-media/models"
	"social-media/repositories"
)

func CreatePost(post *models.Post) (*models.PostResponseViewModel, error) {
	postModel, err := repositories.CreatePost(post)
	if err != nil {
		return nil, err
	}
	postResponseVm := postModel.MapPostDbModelToPostResponseViewModel()
	return postResponseVm, nil
}

func GetPostById(postId int64) (*models.PostResponseViewModel, error) {
	postModel, err := repositories.GetPostById(postId)
	if err != nil {
		return nil, err
	}

	postResponseVm := postModel.MapPostDbModelToPostResponseViewModel()
	return postResponseVm, nil
}

func GetPosts() ([]*models.PostResponseViewModel, error) {
	postModels, err := repositories.GetPosts()
	if err != nil {
		return nil, err
	}

	postResponses := make([]*models.PostResponseViewModel, 0, len(postModels))
	for _, post := range postModels {
		postResponses = append(postResponses, post.MapPostDbModelToPostResponseViewModel())
	}

	return postResponses, nil
}

func GetPostsByUserId(userId int64) ([]*models.PostResponseViewModel, error) {
	_, err := repositories.CheckUserExist(userId)
	if err != nil {
		return nil, err
	}

	postEntites, err := repositories.GetPostsByUserId(userId)
	if err != nil {
		return nil, err
	}

	postResponseVm := make([]*models.PostResponseViewModel, 0, len(postEntites))
	for _, post := range postEntites {
		postResponseVm = append(postResponseVm, post.MapPostDbModelToPostResponseViewModel())
	}
	return postResponseVm, nil
}
