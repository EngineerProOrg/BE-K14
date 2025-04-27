package repositories

import (
	"fmt"
	"social-media/models"
	"social-media/repositories/databases"
)

func CreatePost(post *models.Post) (*models.Post, error) {
	err := databases.GormDb.Create(post).Error
	if err != nil {
		return nil, err
	}

	err = databases.GormDb.Preload("User").First(post, post.Id).Error
	if err != nil {
		return nil, err
	}
	return post, nil
}

func GetPostById(postId int64) (*models.Post, error) {
	// *models.Post nghĩa là con trỏ trỏ tới struct Post.
	// & trong Go nghĩa là lấy địa chỉ vùng nhớ của biến.
	// &post	Là toán tử lấy địa chỉ của biến post
	// var post *models.Post	Một con trỏ Post chưa trỏ đến đâu cả (giá trị ban đầu là nil)
	// post := &models.Post{}	Một con trỏ Post đã trỏ đến vùng nhớ hợp lệ

	//var postEntity *models.Post // khai báo postEntity là *model.Post nhưng chưa cấp phát bộ nhớ
	// Trong GORM, First(&postEntity, postId) kỳ vọng bạn phải đưa vào 1 struct đã được khởi tạo (hoặc địa chỉ của nó), chứ không phải một con trỏ nil.

	postEntity := &models.Post{}

	err := databases.GormDb.First(postEntity, postId)
	if err != nil {
		return nil, fmt.Errorf("post does not exist")
	}
	return postEntity, nil
}
