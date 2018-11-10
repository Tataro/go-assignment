package engine

import (
	"errors"
	"fmt"

	"gitlab.com/upaphong/go-assignment/domain"
)

func (engine *arenaEngine) GetKnight(ID string) (*domain.Knight, error) {
	fighter, err := engine.knightRepository.Find(ID)
	if err != nil {
		return nil, err
	}
	if fighter == nil {
		return nil, errors.New(fmt.Sprintf("Knight #%s not found.", ID))
	}

	return fighter, nil
}

func (engine *arenaEngine) ListKnights() ([]*domain.Knight, error) {
	knights, err := engine.knightRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return knights, nil
}

func (engine *arenaEngine) Fight(fighter1ID string, fighter2ID string) (domain.Fighter, error) {
	var fighter1, fighter2 *domain.Knight
	var err error
	fighter1, err = engine.knightRepository.Find(fighter1ID)
	if err != nil {
		return nil, err
	}
	fighter2, err = engine.knightRepository.Find(fighter2ID)
	if err != nil {
		return nil, err
	}

	return engine.arena.Fight(fighter1, fighter2), nil
}

func (engine *arenaEngine) Save(knight *domain.Knight) (*domain.Knight, error) {
	knight, err := engine.knightRepository.Save(knight)
	if err != nil {
		return nil, err
	}
	return knight, nil
}
