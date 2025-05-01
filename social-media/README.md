✅ Project Progress Checklist
🔐 User Authentication
- [x] Signup API with password hashing
  Endpoint: POST /users/signup
- [x] Signin API with JWT token generation
  Endpoint: POST /users/signin
- [x] JWT Authorization middleware to protect routes

📝 Post Management
- [x] Create Post API
  Endpoint: POST /posts
- [x] Get Post by ID API
  Endpoint: GET /posts/:postId
- [x] Get Posts by User ID API
  Endpoint: GET /users/:userId/posts
- [x] Get All Posts API
  Endpoint: GET /posts

👤 User Profile
- [x] Get current user profile API
  Endpoint: GET /users/profile/:userId
- [x] Update user profile API
  Endpoint: PUT /users/profile/:userId

💬 Comment System
- [X] Create Comment API
  Endpoint: POST /posts/:postId/comments
- [X] Get Comments by Post ID API
  Endpoint: GET /posts/:postId/comments

❤️ Like Feature
- [] Like Post API                                
  Endpoint: POST /posts/:id/like
- [] Unlike Post API                              
  Endpoint: DELETE /posts/:id/like
- [] Like Comment API                             
  Endpoint: /comments/:id/like
- [] Unlike comment API                           
  Endpoint: /comments/:id/like
- [] Count Like Numbers of Post/Comment API       
  Endpoint: /posts/:id/likes or GET /comments/:id/likes
- [] Check Like Stauts                            
  Endpoint: /posts/:id/like-status

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