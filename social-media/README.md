# 🧠 Mini Social Media App (Golang Backend)

A RESTful API built in Go for a simplified social media platform, including user authentication, post creation, commenting, reactions, and more.

---

## 🚀 How to Run

```bash
go run main.go

✅ Project Progress Checklist
🔐 Authentication & User
[X] User Signup
[X] User Signin (JWT generation)
[X] JWT Middleware for protected routes
[X] Get current user profile
[X] Update user profile

📝 Posts
[X] Create Post
[X] Get Post by ID
[X] Get Posts by UserID
[X] Update Post

💬 Comments
[X] Create Comment
[X] Get Comments by PostID
[X] Update Comment
[X] Delete Comment (optional)

👍 Reactions (Like/Heart/Other)
[X] Create or Update Reaction
[X] Get Reactions by Target (post)
[X] Count Reactions by Type

➕ Social Features
 Follow / Unfollow User

 Get Newsfeed (posts from followed users)

🛠️ Other Features
[X] Seeder for sample data
[X] Redis Client Connection
[X] Custom error handling (e.g. ErrUserDoesNotExist, ErrCommentNotInPost)

 Input validation (e.g. required fields, min/max length)

 Pagination for post & comment listing

 Unit tests (optional)

🚧 Advanced Optimization (To be done later)
These are performance/scale-related tasks planned for later phase:

 Redis Caching for posts/newsfeed

 Async job processing (goroutines, channels)

 Database query optimization (indexing, joins)

 Pub/Sub architecture or CQRS pattern