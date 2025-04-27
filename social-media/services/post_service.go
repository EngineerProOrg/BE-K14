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
	postResponseVm := postModel.CreateMappingPostEntityToPostResponseViewModel()
	return postResponseVm, nil
}

func GetPostById(postId int64) (*models.PostResponseViewModel, error) {
	postModel, err := repositories.GetPostById(postId)
	if err != nil {
		return nil, err
	}

	postResponseVm := postModel.CreateMappingPostEntityToPostResponseViewModel()
	return postResponseVm, nil
}

func GetPosts() ([]*models.PostResponseViewModel, error) {
	postModels, err := repositories.GetPosts()
	if err != nil {
		return nil, err
	}

	postResponses := make([]*models.PostResponseViewModel, 0, len(postModels))
	for _, post := range postModels {
		postResponses = append(postResponses, post.CreateMappingPostEntityToPostResponseViewModel())
	}

	return postResponses, nil
}
