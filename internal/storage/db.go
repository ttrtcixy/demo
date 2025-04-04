package storage

import (
	"database/sql"
	"errors"
	_ "github.com/mattn/go-sqlite3"
	"github.com/ttrtcixy/demo/internal/models"
	"log"
)

type DB struct {
	connect *sql.DB
}
type Query struct {
	query string
	args  []any
}

func NewDB() *DB {
	d, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		log.Fatalln(err)
	}

	return &DB{connect: d}
}

var getPartners = `select PartnerId, PartnerType, PartnerName, Director, Phone, Rating, Email, LegalAddress From Partners;`
var ErrPartnersNoFound = errors.New("партнеры не найдены")

func (d *DB) GetPartners() (*models.Partners, error) {
	query := Query{query: getPartners}
	rows, err := d.connect.Query(query.query)
	if err != nil {
		return nil, err
	}

	if !rows.Next() {
		return &models.Partners{}, ErrPartnersNoFound
	}

	partners := models.Partners{}
	for {
		var partner models.Partner
		err := rows.Scan(&partner.Id, &partner.PartnerType, &partner.CompanyName, &partner.Director, &partner.Phone, &partner.Rating, &partner.Email, &partner.Address)
		if err != nil {
			return nil, err
		}
		partners = append(partners, partner)

		if !rows.Next() {
			break
		}
	}

	return &partners, nil
}

var addPartner = `insert into Partners(PartnerType, PartnerName, Director, Phone, Rating, Email, LegalAddress) values(?, ?, ?, ?, ?, ?, ?)`

func (d *DB) AddPartner(partner models.Partner) error {
	args := []any{partner.PartnerType, partner.CompanyName, partner.Director, partner.Phone, partner.Rating, partner.Email, partner.Address}
	query := Query{query: addPartner, args: args}
	_, err := d.connect.Exec(query.query, query.args...)
	if err != nil {
		return err
	}
	return nil
}

var deletePartner = `delete from Partners where PartnerId = ?`

func (d *DB) DeletePartner(id int) error {
	query := Query{query: deletePartner, args: []any{id}}
	_, err := d.connect.Exec(query.query, query.args...)
	if err != nil {
		return err
	}
	return nil
}

var updatePartner = `update Partners set PartnerType = ?, PartnerName = ?, Director = ?, Phone = ?, Rating = ? where PartnerId = ?;`

func (d *DB) UpdatePartner(partner models.Partner) error {
	args := []any{partner.PartnerType, partner.CompanyName, partner.Director, partner.Phone, partner.Rating, partner.Id}
	query := Query{query: updatePartner, args: args}
	_, err := d.connect.Exec(query.query, query.args...)
	if err != nil {
		return err
	}
	return nil
}
