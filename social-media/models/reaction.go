package models

type ReactionCount struct {
	ReactionType string
	Count        int64
}

type UserReactionResponseViewModel struct {
	UserId       int64   `json:"user_id"`
	UserName     string  `json:"user_name"`
	AvatarUrl    *string `json:"avatar_url"`
	ReactionType string  `json:"reaction_type"`
}
