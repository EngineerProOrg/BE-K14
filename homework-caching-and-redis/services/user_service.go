package services

import (
	"errors"
	"fmt"
	"homework-caching-and-redis/models"
	"homework-caching-and-redis/repositories"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateANewUser(user *models.User) error {
	return repositories.CreateANewUser(user)
}

func Login(user *models.User) error {
	return repositories.Login(user)
}

func GenerateSessionId(ginContext *gin.Context, email string) (string, error) {
	sessionID := uuid.NewString()

	err := repositories.RedisClient.Set(ginContext, "session_id:"+sessionID, email, time.Hour).Err()
	log.Printf("✅ New session created: %s for %s", sessionID, email)

	if err != nil {
		return "", errors.New("could not store session in Redis")
	}
	return sessionID, nil
}

func Ping(ginContext *gin.Context) (int, error) {
	// Get session from cookie
	sessionId, err := ginContext.Cookie("session_id")
	if err != nil {
		return http.StatusUnauthorized, errors.New("🔒 No session found. Please login first.")
	}

	// Lock ping for part 2
	ok, err := repositories.RedisClient.SetNX(ginContext, "lock_session_id", sessionId, 5*time.Second).Result()
	if err != nil {
		return http.StatusInternalServerError, err
	}

	// Increase counter redis for part 3
	_, err = repositories.RedisClient.Incr(ginContext, "count_ping_session_id:"+sessionId).Result()
	if err != nil {
		fmt.Printf("❌ Failed to increment ping counter: %v", err)
	}
	if !ok {
		//if someone else has already locked
		return http.StatusTooManyRequests, errors.New("server is busy, please come back later")
	}

	// Rate limit for part 4
	rateLimitSessionIdKey := "rate_limit_session_id:" + sessionId
	countLimit, err := repositories.RedisClient.Incr(ginContext, rateLimitSessionIdKey).Result()
	if err != nil {
		return http.StatusInternalServerError, err
	}

	if countLimit == 1 {
		_ = repositories.RedisClient.Expire(ginContext, rateLimitSessionIdKey, 60*time.Second)
	}

	if countLimit > 2 {
		log.Printf("❌ User %s exceeded ping (%d)", sessionId, countLimit)
		return http.StatusTooManyRequests, errors.New("you exceeded ping in 60 seconds")
	}

	// Add new key ping_leaderboard for part 5
	pingLeaderBoardKey := "ping_leaderboard"
	_ = repositories.RedisClient.ZIncrBy(ginContext, pingLeaderBoardKey, 1, sessionId).Err()

	hyperLogKey := "ping_hyperlog"
	_ = repositories.RedisClient.PFAdd(ginContext, hyperLogKey, sessionId).Err()

	log.Printf("🟢 Ping lock granted for user %s", sessionId)
	time.Sleep(5 * time.Second)

	return http.StatusOK, nil
}

func GetTopPingers(ginContext *gin.Context) ([]models.LeaderboardEntry, error) {
	leaders, err := repositories.RedisClient.ZRevRangeWithScores(ginContext, "ping_leaderboard", 0, 9).Result()
	if err != nil {
		return nil, fmt.Errorf("Failed to get leaderboard: %w", err)
	}

	leaderBoardEntries := make([]models.LeaderboardEntry, 0, len(leaders))

	for _, member := range leaders {
		sessionId := fmt.Sprintf("%v", member.Member)

		// get email from session_id:
		email, _ := repositories.RedisClient.Get(ginContext, "session_id:"+sessionId).Result()
		leaderBoardEntries = append(leaderBoardEntries, models.LeaderboardEntry{
			SessionID: sessionId,
			Email:     email,
			PingCount: int(member.Score),
		})
	}

	return leaderBoardEntries, nil
}

func GetPingUserCount(ctx *gin.Context) (int64, error) {
	count, err := repositories.RedisClient.PFCount(ctx, "ping_hyperlog").Result()
	if err != nil {
		return 0, fmt.Errorf("failed to get HyperLogLog count: %w", err)
	}
	return count, nil
}
