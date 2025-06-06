package utils

import "errors"

// Auth & User
var ErrInvalidLogin = errors.New("email or password is incorrect")
var ErrUserDoesNotExist = errors.New("user does not exist")

// Comment
var ErrCommentDoesNotExist = errors.New("comment does not exist")
var ErrCannotEditComment = errors.New("you are not allowed to update this comment")
var ErrCommentNotInPost = errors.New("comment does not belong to this post")

// Post
var ErrPostDoesNotExist = errors.New("post does not exist")

// Reaction
var ErrReactionDoesNotExist = errors.New("reaction does not exist")

// Generic permission
var ErrPermissionDenied = errors.New("you do not have permission to perform this action")
