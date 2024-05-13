package models

import (
	"listings.com/db"
)

type Listing struct {
	ID			int64
	Slug		string	`binding:"required"`
	Address		string	`binding:"required"`
	Price 		int64	`binding:"required"`
}

func (l *Listing) Save() error {
	query := `
	INSERT INTO listings(slug, address, price)
	VALUES (?, ?, ?)`
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	// close the stmt connection after function finishes execution
	defer stmt.Close()

	result, err := stmt.Exec(l.Slug, l.Address, l.Price)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	l.ID = id

	return err
}

func GetAllListings() ([]Listing, error) {
	query := "SELECT * FROM listings"
	rows, err := db.DB.Query(query)

	if err != nil {
		return nil, err
	}

	// close the rows connection after function finishes execution
	defer rows.Close()

	var listings []Listing

	// Keep looping as long as there is another row
	for rows.Next() {
		var listing Listing

		// Reads content of the row we're currently processing
		err := rows.Scan(&listing.ID, &listing.Slug, &listing.Address, &listing.Price)

		if err != nil {
			return nil, err
		}

		listings = append(listings, listing)
	}

	return listings, nil
}

func GetListingBySlug(slug string) (*Listing, error) {
	query := "SELECT * FROM listings WHERE slug = ?"
	// we will get back only 1 row, therefore can use the QueryRow method
	row := db.DB.QueryRow(query, slug)

	var listing Listing
	err := row.Scan(&listing.ID, &listing.Slug, &listing.Address, &listing.Price)

	if err != nil {
		return nil, err
	}

	return &listing, nil
}

func (listing Listing) Update() error {
	query := `
	UPDATE listings
	SET slug = ?, address = ?, price = ?
	WHERE id = ?
	`
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(listing.Slug, listing.Address, listing.Price, listing.ID)

	return err
}

func (listing Listing) Delete() error {
	query := "DELETE FROM listings WHERE id = ?"
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(listing.ID)

	return err
}
