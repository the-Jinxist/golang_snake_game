package internal

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"
)

var _ ScoreService = &ScoreServiceImpol{}

type Score struct {
	ID        int       `db:"id"`
	User      string    `db:"user"`
	Session   string    `db:"session"`
	Value     int       `db:"value"`
	CreatedAt time.Time `db:"created_at"`
}

type ScoreService interface {
	GetHighScore(ctx context.Context) (Score, error)
	GetScores(ctx context.Context) ([]Score, error)
	SetCurrentScore(ctx context.Context, value int) error
	GetCurrentScore(ctx context.Context) (int, error)
}

func NewScoreService(user string, sessionMgr SessionManager, db *sql.DB) ScoreService {
	return &ScoreServiceImpol{
		db:          db,
		CurrentUser: user,
		Session:     sessionMgr,
	}
}

type ScoreServiceImpol struct {
	CurrentUser string
	db          *sql.DB
	Session     SessionManager
}

// GetHighScore implements ScoreService.
func (s *ScoreServiceImpol) GetHighScore(ctx context.Context) (Score, error) {
	var score Score

	err := s.db.QueryRowContext(ctx, `select * from scores order by value desc limit 1`).Scan(
		&score.ID,
		&score.User,
		&score.Session,
		&score.Value,
		&score.CreatedAt,
	)

	if err != nil {
		fmt.Printf("big error: %s", err)
		if err == sql.ErrNoRows {
			session, _ := s.Session.GetCurrentSession()
			return Score{
				User:      s.CurrentUser,
				Session:   session,
				Value:     0,
				CreatedAt: time.Now(),
			}, nil
		}
		return score, nil
	}

	return score, nil
}

// GetScores implements ScoreService.
func (s *ScoreServiceImpol) GetScores(ctx context.Context) ([]Score, error) {

	scores := make([]Score, 0, 5)

	rows, err := s.db.QueryContext(ctx, `select * from scores order by value desc limit 5`)
	if err != nil {
		return scores, nil
	}

	defer rows.Close()
	for rows.Next() {
		var score Score
		err = rows.Scan(
			&score.ID,
			&score.User,
			&score.Session,
			&score.Value,
			&score.CreatedAt,
		)
		if err != nil {
			log.Printf("issue with scanning each rows: %s", err)
			return scores, err
		}

		scores = append(scores, score)
	}

	if err := rows.Err(); err != nil {
		log.Printf("Error encountered while iterating over rows: %s", err)
	}

	return scores, nil

}

func (s *ScoreServiceImpol) GetCurrentScore(ctx context.Context) (int, error) {
	session, _ := s.Session.GetCurrentSession()
	var score int

	err := s.db.QueryRowContext(ctx,
		`select value from scores where session = ?`, session,
	).Scan(
		&score,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return score, nil
		}
		return score, err
	}

	return score, nil

}

// SetScore implements ScoreService.
func (s *ScoreServiceImpol) SetCurrentScore(ctx context.Context, value int) error {

	fmt.Printf("setting score: %d", value)
	session, _ := s.Session.GetCurrentSession()
	_, err := s.db.ExecContext(ctx,
		`insert into scores ("user", session, value) 
	 values (?, ?, ?)
	 on conflict(session) do update set
		value = excluded.value
	where scores.session = excluded.session and scores."user" = excluded."user";
	 `, s.CurrentUser, session, value)
	if err != nil {
		fmt.Printf("error while setting score: %s", err)
		return err
	}

	fmt.Printf("seeting score -- success: %d", value)
	return nil
}
