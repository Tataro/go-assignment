package engine

import "gitlab.com/upaphong/go-assignment/domain"

type Engine interface {
	GetKnight(ID string) (*domain.Knight, error)
	ListKnights() ([]*domain.Knight, error)
	Fight(fighter1ID string, fighter2ID string) (domain.Fighter, error)
	Save(knight *domain.Knight) (*domain.Knight, error)
}

type KnightRepository interface {
	Find(ID string) (*domain.Knight, error)
	FindAll() ([]*domain.Knight, error)
	Save(knight *domain.Knight) (*domain.Knight, error)
}

type DatabaseProvider interface {
	GetKnightRepository() KnightRepository
}

type arenaEngine struct {
	arena            *domain.Arena
	knightRepository KnightRepository
}

func NewEngine(db DatabaseProvider) Engine {
	return &arenaEngine{
		arena:            &domain.Arena{},
		knightRepository: db.GetKnightRepository(),
	}
}
