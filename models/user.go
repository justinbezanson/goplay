package models

import "github.com/jmoiron/sqlx"

type User struct {
	ID           int64  `db:"id"`
	OrgID        int64  `db:"org_id"`
	Email        string `db:"email"`
	PasswordHash string `db:"password_hash"`
	Name         string `db:"name"`
	Role         string `db:"role"`
	CreatedAt    string `db:"created_at"`
}

func CreateUser(db *sqlx.DB, u *User) error {
	result, err := db.NamedExec(`
		INSERT INTO users (org_id, email, password_hash, name, role)
		VALUES (:org_id, :email, :password_hash, :name, :role)
	`, u)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	u.ID = id
	return nil
}

func GetUserByID(db *sqlx.DB, id int64) (*User, error) {
	var u User
	err := db.Get(&u, "SELECT * FROM users WHERE id = ?", id)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func GetUserByEmail(db *sqlx.DB, email string) (*User, error) {
	var u User
	err := db.Get(&u, "SELECT * FROM users WHERE email = ?", email)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func ListUsersByOrg(db *sqlx.DB, orgID int64) ([]User, error) {
	var users []User
	err := db.Select(&users, "SELECT * FROM users WHERE org_id = ? ORDER BY created_at", orgID)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func UpdateUserPassword(db *sqlx.DB, id int64, hash string) error {
	_, err := db.Exec("UPDATE users SET password_hash = ? WHERE id = ?", hash, id)
	return err
}
