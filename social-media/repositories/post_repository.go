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

	err := databases.GormDb.Preload("User").First(postEntity, postId).Error // First yêu cầu truyền địa chỉ vùng nhớ pointer
	if err != nil {
		return nil, fmt.Errorf("post does not exist")
	}
	return postEntity, nil
}

func GetPosts() ([]models.Post, error) {
	var posts []models.Post

	//GORM chỉ ghi vào vùng nhớ gốc, nên mình phải truyền đúng địa chỉ vùng nhớ (&) cho nó
	err := databases.GormDb.Preload("User").Order("created_at DESC").Find(&posts).Error

	if err != nil {
		return nil, err
	}
	return posts, nil
}

func GetPostsByUserId(userId int64) ([]models.Post, error) {
	var postEntities []models.Post

	err := databases.GormDb.Preload("User").Where("user_id = ?", userId).Find(&postEntities).Error
	if err != nil {
		return nil, fmt.Errorf("post does not exist")
	}
	return postEntities, nil
}
