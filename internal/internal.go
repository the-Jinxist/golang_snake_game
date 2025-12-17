package internal

import (
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var (
	sessionManager SessionManager
	scoreService   ScoreService
)

func GetSessionManager() SessionManager {
	return sessionManager
}

func GetScoreService() ScoreService {
	return scoreService
}

func IntializeConfigs() {
	db := CreateDB()
	if db == nil {
		log.Fatalf("sqlite db cannot be initialized")
	}

	user, err := os.Hostname()
	if err != nil {
		log.Fatalf("system host name cannot be retrieved")
	}

	sessionManager = NewSessionManager()
	scoreService = NewScoreService(user, sessionManager, db)
}
