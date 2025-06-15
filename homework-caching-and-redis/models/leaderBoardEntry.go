package models

type LeaderboardEntry struct {
	SessionID string `json:"session_id"`
	Email     string `json:"email"`
	PingCount int    `json:"ping_count"`
}
