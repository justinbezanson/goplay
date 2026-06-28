package models

import "github.com/jmoiron/sqlx"

type Organization struct {
	ID       int64  `db:"id"`
	Name     string `db:"name"`
	Email    string `db:"email"`
	Address  string `db:"address"`
	Address2 *string `db:"address2"`
	City     string `db:"city"`
	State    string `db:"state"`
	Zip      string `db:"zip"`
	Country  string `db:"country"`
	CreateAt string `db:"created_at"`
}

func CreateOrg(db *sqlx.DB, o *Organization) error {
	result, err := db.NamedExec(`
		INSERT INTO organizations (name, email, address, address2, city, state, zip, country)
		VALUES (:name, :email, :address, :address2, :city, :state, :zip, :country)
	`, o)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	o.ID = id
	return nil
}

func GetOrg(db *sqlx.DB, id int64) (*Organization, error) {
	var org Organization
	err := db.Get(&org, "SELECT * FROM organizations WHERE id = ?", id)
	if err != nil {
		return nil, err
	}
	return &org, nil
}

func UpdateOrg(db *sqlx.DB, o *Organization) error {
	_, err := db.NamedExec(`
		UPDATE organizations SET name=:name, email=:email, address=:address,
		address2=:address2, city=:city, state=:state, zip=:zip, country=:country
		WHERE id=:id
	`, o)
	return err
}
