import { Box, Paper, Typography } from "@mui/material";
import PostCard from "./PostCard";
import HttpClient from "../../apis/HttpClient";
import { useEffect, useState } from "react";
import { PostWithAuthorViewModel } from "../../models/post";

function PostCardList() {
  const [posts, setPosts] = useState<PostWithAuthorViewModel[]>([]);

  useEffect(() => {
    const getPosts = async () => {
      const result = await HttpClient.Post.GetAllPosts();
      setPosts(result.posts);
    };
    getPosts();
  }, []);

  return (
    <Box sx={{ p: 2 }}>
      <Typography variant="h5">News Feed</Typography>
      {posts.map((post, index) => {
        return (
          <Paper sx={{ p: 2, mt: 2 }}>
            <PostCard postWithAuthor={post} />
          </Paper>
        );
      })}
    </Box>
  );
}

export default PostCardList;
