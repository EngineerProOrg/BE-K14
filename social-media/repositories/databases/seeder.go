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
			FirstName: "User 001",
			LastName:  "Global InfoTrack",
			Name:      "User 001 - Global InfoTrack",
			Birthday:  time.Date(2000, 8, 15, 0, 0, 0, 0, time.UTC),
			Email:     "user001@infotrack.com.au",
			Username:  utils.GetUsernameFromEmail("user001@infotrack.com.au"),
			Password:  hash,
			CreatedAt: time.Now(),
			UpdatedAt: nil,
		},
		{
			FirstName: "User 002",
			LastName:  "Global InfoTrack",
			Name:      "User 002 - Global InfoTrack",
			Birthday:  time.Date(1998, 9, 11, 0, 0, 0, 0, time.UTC),
			Email:     "user002@infotrack.com.au",
			Username:  utils.GetUsernameFromEmail("user002@infotrack.com.au"),
			Password:  hash,
			CreatedAt: time.Now(),
			UpdatedAt: nil,
		},
		{
			FirstName: "User 003",
			LastName:  "Global InfoTrack",
			Name:      "User 003 - Global InfoTrack",
			Birthday:  time.Date(1997, 11, 20, 0, 0, 0, 0, time.UTC),
			Email:     "user003@infotrack.com.au",
			Username:  utils.GetUsernameFromEmail("user003@infotrack.com.au"),
			Password:  hash,
			CreatedAt: time.Now(),
			UpdatedAt: nil,
		},
		{
			FirstName: "User 004",
			LastName:  "Global InfoTrack",
			Name:      "User 004 - Global InfoTrack",
			Birthday:  time.Date(1999, 12, 24, 0, 0, 0, 0, time.UTC),
			Email:     "user004@infotrack.com.au",
			Username:  utils.GetUsernameFromEmail("user004@infotrack.com.au"),
			Password:  hash,
			CreatedAt: time.Now(),
			UpdatedAt: nil,
		},
		{
			FirstName: "User 005",
			LastName:  "Global InfoTrack",
			Name:      "User 005 - Global InfoTrack",
			Birthday:  time.Date(1996, 10, 17, 0, 0, 0, 0, time.UTC),
			Email:     "user005@infotrack.com.au",
			Username:  utils.GetUsernameFromEmail("user005@infotrack.com.au"),
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
		log.Println("✅ Likes table already has data. Skipping seeding.")
		return
	}

	reactions := []models.Reaction{
		// Likes on posts (PostId != nil, CommentId = nil)
		{UserId: 2, PostId: 1, ReactionType: constants.ReactionLike, CreatedAt: time.Now()},
		{UserId: 3, PostId: 1, ReactionType: constants.ReactionLove, CreatedAt: time.Now()},
		{UserId: 1, PostId: 2, ReactionType: constants.ReactionHaha, CreatedAt: time.Now()},
		{UserId: 4, PostId: 3, ReactionType: constants.ReactionSad, CreatedAt: time.Now()},
		{UserId: 5, PostId: 3, ReactionType: constants.ReactionFire, CreatedAt: time.Now()},

		// Likes on comments (must include both PostId and CommentId)
		{UserId: 1, PostId: 1, CommentId: ptrInt(2), ReactionType: constants.ReactionLike, CreatedAt: time.Now()},
		{UserId: 3, PostId: 2, CommentId: ptrInt(3), ReactionType: constants.ReactionLike, CreatedAt: time.Now()},
		{UserId: 4, PostId: 2, CommentId: ptrInt(3), ReactionType: constants.ReactionHaha, CreatedAt: time.Now()},
		{UserId: 2, PostId: 3, CommentId: ptrInt(5), ReactionType: constants.ReactionHaha, CreatedAt: time.Now()},
	}

	if err := db.Create(&reactions).Error; err != nil {
		log.Printf("❌ Failed to seed reactions: %v\n", err)
		return
	}

	log.Println("✅ Successfully seeded sample reactions.")
}

func ptrInt(i int64) *int64 {
	return &i
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
