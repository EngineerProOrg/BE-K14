package services

import (
	"social-media/models"
	"social-media/repositories"
	"time"
)

func CreatePost(username string, post *models.Post) (*models.CreatedOrUpdatedPostResponseViewModel, error) {
	postModel, err := repositories.CreatePost(post)
	if err != nil {
		return nil, err
	}

	author, err := GetCachedUserInfoByUsername(username)
	if err != nil {
		return nil, err
	}

	postResponseVm := postModel.MapPostDbModelToCreatedPostResponseViewModel(author)
	return postResponseVm, nil
}

func GetPostById(postId int64) (*models.CreatedOrUpdatedPostResponseViewModel, error) {
	postModel, err := repositories.GetPostById(postId)
	if err != nil {
		return nil, err
	}

	author := postModel.User.MapUserDbModelToUserProfileResponseViewModel()

	postResponseVm := postModel.MapPostDbModelToCreatedPostResponseViewModel(author)
	return postResponseVm, nil
}

func GetPosts() (*models.PostsWithAuthorResponse, error) {
	postModels, err := repositories.GetPosts()
	if err != nil {
		return nil, err
	}

	postResponses := make([]models.PostWithAuthorViewModel, 0, len(postModels))
	for _, post := range postModels {
		vm := post.MapPostDbModelToPostWithAuthorViewModel()
		postResponses = append(postResponses, *vm)
	}

	return &models.PostsWithAuthorResponse{
		Posts: postResponses,
	}, nil
}

func GetPostsByUserId(userId int64, username string) (*models.PostUserResponseViewModel, error) {
	postEntities, err := repositories.GetPostsByUserId(userId)
	if err != nil {
		return nil, err
	}

	author, err := GetCachedUserInfoByUsername(username)
	if err != nil {
		return nil, err
	}

	postResponseVm := make([]models.PostResponseViewModel, 0, len(postEntities))
	for _, post := range postEntities {
		postVmPtr := post.MapPostDbModelToPostResponseViewModel()
		postResponseVm = append(postResponseVm, *postVmPtr)
	}

	return &models.PostUserResponseViewModel{
		Author: *author,
		Posts:  postResponseVm,
	}, nil
}

func UpdatePost(postId int64, userId int64, postRequestViewModel *models.PostRequestViewModel) error {
	postDbModel := models.MapPostRequestViewModelToPostDbModel(postRequestViewModel)
	now := time.Now()
	postDbModel.UpdatedAt = &now
	return repositories.UpdatePost(postId, userId, postDbModel)
}
