package models

import "github.com/jmoiron/sqlx"

type Session struct {
	ID        int64  `db:"id"`
	UserID    int64  `db:"user_id"`
	Token     string `db:"token"`
	Data      []byte `db:"data"`
	ExpiresAt string `db:"expires_at"`
	CreatedAt string `db:"created_at"`
}

func CreateSession(db *sqlx.DB, s *Session) error {
	result, err := db.NamedExec(`
		INSERT INTO sessions (user_id, token, data, expires_at)
		VALUES (:user_id, :token, :data, :expires_at)
	`, s)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	s.ID = id
	return nil
}

func GetSessionByToken(db *sqlx.DB, token string) (*Session, error) {
	var s Session
	err := db.Get(&s, "SELECT * FROM sessions WHERE token = ? AND expires_at > datetime('now')", token)
	if err != nil {
		return nil, err
	}
	return &s, nil
}

func DeleteSession(db *sqlx.DB, token string) error {
	_, err := db.Exec("DELETE FROM sessions WHERE token = ?", token)
	return err
}

func DeleteExpiredSessions(db *sqlx.DB) error {
	_, err := db.Exec("DELETE FROM sessions WHERE expires_at <= datetime('now')")
	return err
}
