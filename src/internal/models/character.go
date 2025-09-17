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
	Equipment             Equipment
	Skills                []string
	ExpCurrent            int
	ExpMax                int
	InventoryCapacity     int
	InventoryUpgradesUsed int
	Initiative            int
}

func InitCharacter(name, class string) *Character {
	return &Character{}
}

func DisplayInfo(c *Character) {
}

func IsDead(c *Character) bool {
	return c.HPCurrent <= 0
}

func Revive(c *Character) {
	if IsDead(c) {
		c.HPCurrent = c.HPMax / 2
	}
}

func AdjustHP(c *Character, delta int) {
	c.HPCurrent += delta
	if c.HPCurrent > c.HPMax {
		c.HPCurrent = c.HPMax
	}
	if c.HPCurrent < 0 {
		c.HPCurrent = 0
	}
}

func AddGold(c *Character, amount int) error {
	if amount <= 0 {
		return errors.New("montant invalide")
	}
	c.Gold += amount
	return nil
}

func SpendGold(c *Character, amount int) error {
	if amount > c.Gold {
		return errors.New("vous n'avez pas assez d'argent")
	}
	c.Gold -= amount
	return nil
}

func LearnSkill(c *Character, skill string) error {
	for _, s := range c.Skills {
		if s == skill {
			return errors.New("le personnage connaît déjà cette compétence")
		}
	}
	c.Skills = append(c.Skills, skill)
	return nil
}
