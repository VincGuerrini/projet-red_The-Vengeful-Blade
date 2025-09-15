package models

import (
	"errors"
)

type Character struct {
	Name                  string
	Class                 string
	Level                 int
	HPMax                 int
	HPCurrent             int
	ManaMax               int
	ManaCurrent           int
	Gold                  int
	Inventory             []Item
	MaxSlots              int
	Equipment             []string
	Weapon                []Equipment
	Skills                []string
	ExpCurrent            int
	ExpMax                int
	InventoryCapacity     int
	InventoryUpgradesUsed string
	Initiative            string
}

func InitCharacter(name, class string) *Character

func DisplayInfo(c *Character)

func IsDead(c *Character) bool {
	if c.HPCurrent <= 0 {
		return true
	} else {
		return false
	}
}

func Revive(c *Character) {
	if IsDead(c) == true {
		c.HPCurrent = c.HPMax / 2
	}
}

func AdjustHP(c *Character, delta int) {
	if delta != 0 {
		switch {
		case delta > 0:
			c.HPCurrent = c.HPCurrent + delta
		case delta < 0:
			c.HPCurrent = c.HPCurrent - delta
		}
	} else {
		c.HPCurrent = c.HPCurrent + 0
	}
}

func AddGold(c *Character, amount int) error {
	var err error
	if amount <= 0 {
		err = errors.New("Montant invalide...")
	} else {
		c.Gold = c.Gold + amount
		err = nil
	}
	return err
}

func SpendGold(c *Character, amount int) error {
	var err error
	if amount > c.Gold {
		err = errors.New("Vous n'avez pas assez d'argent !")
	} else {
		c.Gold = c.Gold - amount
		err = nil
	}
	return err
}

func LearnSkill(c *Character, skill string) error
