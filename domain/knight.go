package domain

type Knight struct {
	ID          string `json:"id"`
	Name        string `json:"name" validate:"required"`
	Strength    int    `json:"strength" validate:"required"`
	WeaponPower int    `json:"weapon_power" validate:"required"`
}

func (k *Knight) GetID() string {
	return k.ID
}

func (k *Knight) GetPower() float64 {
	return float64(k.Strength + k.WeaponPower)
}
