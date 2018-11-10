package database

import (
	"database/sql"

	"gitlab.com/upaphong/go-assignment/domain"
)

type knightRepository struct {
	db *sql.DB
}

func (repository *knightRepository) Find(ID string) (*domain.Knight, error) {
	var knight domain.Knight
	query := `SELECT id, name, strength, weapon_power
		FROM knights
		WHERE id=$1`
	err := repository.db.QueryRow(query, ID).
		Scan(&knight.ID, &knight.Name, &knight.Strength, &knight.WeaponPower)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &knight, nil
}

func (repository *knightRepository) FindAll() ([]*domain.Knight, error) {
	var knights []*domain.Knight
	query := `SELECT id, name, strength, weapon_power
		FROM knights`
	rows, err := repository.db.Query(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var knight domain.Knight
		err = rows.Scan(&knight.ID, &knight.Name, &knight.Strength, &knight.WeaponPower)
		if err != nil {
			return nil, err
		}
		knights = append(knights, &knight)
	}
	return knights, nil
}

func (repository *knightRepository) Save(knight *domain.Knight) (*domain.Knight, error) {
	query := `INSERT INTO knights (name, strength, weapon_power)
		VALUES ($1, $2, $3) RETURNING id;`
	err := repository.db.QueryRow(query, knight.Name, knight.Strength, knight.WeaponPower).
		Scan(&knight.ID)

	if err != nil {
		return nil, err
	}
	return knight, nil
}
