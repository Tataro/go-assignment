package database

import (
	"log"
	"testing"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"gitlab.com/upaphong/go-assignment/domain"
	"gitlab.com/upaphong/go-assignment/engine"
)

var (
	p          *Provider
	knightRepo engine.KnightRepository
)

func ensureTableExists() {
	const tableCreationQuery = `CREATE TABLE IF NOT EXISTS knights
	(
		id SERIAL,
		name TEXT NOT NULL,
		strength INTEGER NOT NULL,
		weapon_power  INTEGER NOT NULL,
		CONSTRAINT knights_pkey PRIMARY KEY (id)
	)`
	if _, err := p.DB.Exec(tableCreationQuery); err != nil {
		log.Fatal(err)
	}
}

func clearTable() {
	p.DB.Exec("TRUNCATE knights RESTART IDENTITY")
}

func TestBefore(t *testing.T) {
	p = NewProvider()
	knightRepo = p.GetKnightRepository()
	ensureTableExists()
	clearTable()
}

func TestSave(t *testing.T) {
	knight := domain.Knight{
		Name:        "tat",
		Strength:    3,
		WeaponPower: 4,
	}

	_, err := knightRepo.Save(&knight)
	if err != nil {
		t.Fatal(err)
	}
}

func TestFind(t *testing.T) {
	expectedResult := &domain.Knight{
		ID:          "1",
		Name:        "tat",
		Strength:    3,
		WeaponPower: 4,
	}
	knight, err := knightRepo.Find("1")
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, expectedResult, knight, "should be equal")
}

func TestFindAll(t *testing.T) {
	knights, err := knightRepo.FindAll()
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 1, len(knights), "len should equal 1")
}

func TestAfter(t *testing.T) {
	clearTable()
}
