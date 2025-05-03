✅ Project Progress Checklist
🔐 User Authentication
- [x] Signup API with password hashing
  POST /users/signup
- [x] Signin API with JWT token generation
  POST /users/signin
- [x] JWT Authorization middleware to protect routes

📝 Post Management
- [x] Create Post API
  POST /posts
- [x] Get Post by ID API
  GET /posts/:postId
- [x] Get Posts by User ID API
  GET /users/:userId/posts
- [x] Get All Posts API
  GET /posts

👤 User Profile
- [x] Get current user profile API
  GET /users/profile/:userId
- [x] Update user profile API
  PUT /users/profile/:userId

💬 Comment System
- [X] Create Comment API
  POST /posts/:postId/comments
- [X] Get Comments by Post ID API
  GET /posts/:postId/comments

❤️ Like Feature (Inprogress)
- [] Like Post API
  POST /posts/:postId/reaction
- [] Unlike Post API
  DELETE /posts/:postId/reaction
- [] Like Comment API
  POST /posts/:postId/comments/:commentId/reaction
- [] Unlike comment API
  DELETE /posts/:postId/comments/:commentId/reaction
- [] Count Likes Post API
  GET /posts/:postId/reactions
- [] Count Likes Cmt API
  GET /posts/:postId/comments/:commentId/reactions
- [] Check Like Status
  GET /posts/:postId/reaction-status
- [] Check Like Comment
  GET /posts/:postId/comments/:commentId/reaction-status

🔗 Follow System (Coming soon)
- [] Follow User API
- [] Unfollow User API

📰 Newsfeed (Coming soon)
- [] Get Newsfeed from followed users API

⚙️ Other Features
- [x] Data Seeder for testing
- [x] Redis client setup
- [] Redis caching for performance (optional)
- [] Pagination for listing posts and comments
- [] Input validation for requests
- [] Consistent error handling and responses