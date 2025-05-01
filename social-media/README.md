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

❤️ Like Feature
- [] Like Post API
  POST /posts/:id/like
- [] Unlike Post API
  DELETE /posts/:id/like
- [] Like Comment API
  /comments/:id/like
- [] Unlike comment API
  /comments/:id/like
- [] Count Like Numbers of Post/Comment API
  /posts/:id/likes or GET /comments/:id/likes
- [] Check Like Status
  /posts/:id/like-status

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