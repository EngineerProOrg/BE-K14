package databases

import (
	"log"
	"social-media/constants"
	"social-media/models"
	"social-media/utils"
	"time"

	"gorm.io/gorm"
)

func seedSampleUserData(db *gorm.DB) {
	var count int64
	if err := db.Model(&models.User{}).Count(&count).Error; err != nil {
		log.Printf("❌ Failed to count users: %v\n", err)
		return
	}

	if count > 0 {
		log.Println("✅ Users table already has data. Skipping seeding.")
		return
	}

	hash, _ := utils.HashPassword("P@ssword123")
	users := []models.User{
		{
			FirstName: "Joe",
			LastName:  "Walker",
			Name:      "Joe Walker",
			Birthday:  time.Date(1990, 1, 10, 0, 0, 0, 0, time.UTC),
			Email:     "joe.walker@gmail.com",
			Username:  utils.GetUsernameFromEmail("joe.walker@gmail.com"),
			Password:  hash,
			CreatedAt: time.Now(),
			UpdatedAt: nil,
		},
		{
			FirstName: "Alice",
			LastName:  "Nguyen",
			Name:      "Alice Nguyen",
			Birthday:  time.Date(1992, 3, 15, 0, 0, 0, 0, time.UTC),
			Email:     "alice.nguyen@gmail.com",
			Username:  utils.GetUsernameFromEmail("alice.nguyen@gmail.com"),
			Password:  hash,
			CreatedAt: time.Now(),
			UpdatedAt: nil,
		},
		{
			FirstName: "Bill",
			LastName:  "Sanders",
			Name:      "Bill Sanders",
			Birthday:  time.Date(1988, 7, 21, 0, 0, 0, 0, time.UTC),
			Email:     "bill.sanders@gmail.com",
			Username:  utils.GetUsernameFromEmail("bill.sanders@gmail.com"),
			Password:  hash,
			CreatedAt: time.Now(),
			UpdatedAt: nil,
		},
		{
			FirstName: "Mary",
			LastName:  "Tran",
			Name:      "Mary Tran",
			Birthday:  time.Date(1995, 5, 30, 0, 0, 0, 0, time.UTC),
			Email:     "mary.tran@gmail.com",
			Username:  utils.GetUsernameFromEmail("mary.tran@gmail.com"),
			Password:  hash,
			CreatedAt: time.Now(),
			UpdatedAt: nil,
		},
		{
			FirstName: "Tom",
			LastName:  "Pham",
			Name:      "Tom Pham",
			Birthday:  time.Date(1993, 12, 9, 0, 0, 0, 0, time.UTC),
			Email:     "tom.pham@gmail.com",
			Username:  utils.GetUsernameFromEmail("tom.pham@gmail.com"),
			Password:  hash,
			CreatedAt: time.Now(),
			UpdatedAt: nil,
		},
	}

	if err := db.Create(&users).Error; err != nil {
		log.Printf("❌ Failed to seed users: %v\n", err)
		return
	}

	log.Println("✅ Successfully seeded sample users.")
}

func seedSamplePostData(db *gorm.DB) {
	var count int64
	if err := db.Model(&models.Post{}).Count(&count).Error; err != nil {
		log.Printf("❌ Failed to count posts: %v\n", err)
		return
	}

	if count > 0 {
		log.Println("✅ Posts table already has data. Skipping seeding.")
		return
	}

	posts := []models.Post{
		{
			Title:     "Học Golang trong 30 ngày",
			Content:   "Golang là một ngôn ngữ tuyệt vời cho backend.",
			UserId:    1, // user 001
			CreatedAt: time.Now(),
		},
		{
			Title:     "Tại sao nên dùng GORM?",
			Content:   "GORM giúp quản lý database trong Golang dễ dàng hơn nhiều.",
			UserId:    1,
			CreatedAt: time.Now(),
		},
		{
			Title:     "Sự khác biệt giữa Gin và Echo",
			Content:   "Gin có tốc độ tốt hơn trong nhiều benchmark.",
			UserId:    2,
			CreatedAt: time.Now(),
		},
		{
			Title:     "Tạo RESTful API với Gin",
			Content:   "Xây dựng API đơn giản và nhanh chóng với Gin framework.",
			UserId:    3,
			CreatedAt: time.Now(),
		},
		{
			Title:     "Deploy Golang App lên Docker",
			Content:   "Việc deploy Golang app vào container rất tiện lợi nhờ Docker.",
			UserId:    4,
			CreatedAt: time.Now(),
		},
		{
			Title:     "Cách tối ưu GORM query",
			Content:   "Để tránh N+1 query và tăng tốc độ, nên dùng Preload hoặc Joins.",
			UserId:    5,
			CreatedAt: time.Now(),
		},
	}

	if err := db.Create(&posts).Error; err != nil {
		log.Printf("❌ Failed to seed posts: %v\n", err)
		return
	}

	log.Println("✅ Successfully seeded sample posts.")
}

func seedSampleCommentData(db *gorm.DB) {
	var count int64
	if err := db.Model(&models.Comment{}).Count(&count).Error; err != nil {
		log.Printf("❌ Failed to count comments: %v\n", err)
		return
	}
	if count > 0 {
		log.Println("✅ Comments table already has data. Skipping seeding.")
		return
	}

	comments := []models.Comment{
		{Content: "Bài viết rất hay!", UserId: 1, PostId: 1},
		{Content: "Cảm ơn bạn đã chia sẻ.", UserId: 2, PostId: 1},
		{Content: "Mình cũng đang học Golang đây!", UserId: 3, PostId: 2},
		{Content: "Dễ hiểu quá 👍", UserId: 4, PostId: 2},
		{Content: "Bài này hữu ích ghê.", UserId: 5, PostId: 3},
	}

	if err := db.Create(&comments).Error; err != nil {
		log.Printf("❌ Failed to seed comments: %v\n", err)
		return
	}

	log.Println("✅ Successfully seeded sample comments.")
}

func seedSampleReactionData(db *gorm.DB) {
	var count int64
	if err := db.Model(&models.Reaction{}).Count(&count).Error; err != nil {
		log.Printf("❌ Failed to count reactions: %v\n", err)
		return
	}

	if count > 0 {
		log.Println("✅ Reactions table already has data. Skipping seeding.")
		return
	}

	now := time.Now()
	reactions := []models.Reaction{
		// 👍 Likes on posts (TargetType = "post")
		{UserId: 2, TargetId: 1, TargetType: "post", ReactionType: constants.ReactionLike, CreatedAt: now},
		{UserId: 3, TargetId: 1, TargetType: "post", ReactionType: constants.ReactionLove, CreatedAt: now},
		{UserId: 1, TargetId: 2, TargetType: "post", ReactionType: constants.ReactionHaha, CreatedAt: now},
		{UserId: 4, TargetId: 3, TargetType: "post", ReactionType: constants.ReactionSad, CreatedAt: now},
		{UserId: 5, TargetId: 3, TargetType: "post", ReactionType: constants.ReactionFire, CreatedAt: now},

		// 💬 Likes on comments (TargetType = "comment")
		{UserId: 1, TargetId: 2, TargetType: "comment", ReactionType: constants.ReactionLike, CreatedAt: now},
		{UserId: 3, TargetId: 3, TargetType: "comment", ReactionType: constants.ReactionLike, CreatedAt: now},
		{UserId: 4, TargetId: 3, TargetType: "comment", ReactionType: constants.ReactionHaha, CreatedAt: now},
		{UserId: 2, TargetId: 5, TargetType: "comment", ReactionType: constants.ReactionHaha, CreatedAt: now},
	}

	if err := db.Create(&reactions).Error; err != nil {
		log.Printf("❌ Failed to seed reactions: %v\n", err)
		return
	}

	log.Println("✅ Successfully seeded sample reactions.")
}

func ClearAllData(db *gorm.DB) {
	log.Println("🧹 Clearing all tables: Comments → Posts → Users")

	// delete comment first
	if err := db.Exec("DELETE FROM comments").Error; err != nil {
		log.Printf("❌ Failed to clear comments: %v\n", err)
	} else {
		log.Println("✅ Cleared comments")
	}

	// delete post
	if err := db.Exec("DELETE FROM posts").Error; err != nil {
		log.Printf("❌ Failed to clear posts: %v\n", err)
	} else {
		log.Println("✅ Cleared posts")
	}

	// delete user
	if err := db.Exec("DELETE FROM users").Error; err != nil {
		log.Printf("❌ Failed to clear users: %v\n", err)
	} else {
		log.Println("✅ Cleared users")
	}
}
