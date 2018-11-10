package domain

type Arena struct{}

func (arena *Arena) Fight(fighter1 Fighter, fighter2 Fighter) Fighter {
	f1Power := fighter1.GetPower()
	f2Power := fighter2.GetPower()
	if f1Power > f2Power {
		return fighter1
	} else if f2Power > f1Power {
		return fighter2
	}
	return nil
}
