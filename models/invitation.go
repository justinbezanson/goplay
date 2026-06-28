package models

import "github.com/jmoiron/sqlx"

type Invitation struct {
	ID        int64  `db:"id"`
	OrgID     int64  `db:"org_id"`
	Email     string `db:"email"`
	Token     string `db:"token"`
	ExpiresAt string `db:"expires_at"`
	Used      bool   `db:"used"`
	CreatedBy int64  `db:"created_by"`
	CreatedAt string `db:"created_at"`
}

func CreateInvitation(db *sqlx.DB, inv *Invitation) error {
	result, err := db.NamedExec(`
		INSERT INTO invitations (org_id, email, token, expires_at, used, created_by)
		VALUES (:org_id, :email, :token, :expires_at, :used, :created_by)
	`, inv)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	inv.ID = id
	return nil
}

func GetInvitationByToken(db *sqlx.DB, token string) (*Invitation, error) {
	var inv Invitation
	err := db.Get(&inv, "SELECT * FROM invitations WHERE token = ?", token)
	if err != nil {
		return nil, err
	}
	return &inv, nil
}

func MarkInvitationUsed(db *sqlx.DB, id int64) error {
	_, err := db.Exec("UPDATE invitations SET used = 1 WHERE id = ?", id)
	return err
}

func ListPendingInvitations(db *sqlx.DB, orgID int64) ([]Invitation, error) {
	var invs []Invitation
	err := db.Select(&invs, "SELECT * FROM invitations WHERE org_id = ? AND used = 0 AND expires_at > datetime('now') ORDER BY created_at", orgID)
	if err != nil {
		return nil, err
	}
	return invs, nil
}
