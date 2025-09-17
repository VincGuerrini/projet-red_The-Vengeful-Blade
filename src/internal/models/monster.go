package models

import (
	"math/rand"
)

type Monster struct {
	Name         string
	HPmax        int
	Attack       int
	Initiative   int
	PatternState int
}

func InitSanglier() Monster {
	return Monster{
		Name:         "Sanglier",
		HPmax:        15,
		Attack:       rand.Intn(5) + 5,
		Initiative:   3,
		PatternState: 0,
	}
}

func InitLoup() Monster {
	return Monster{
		Name:         "Loup",
		HPmax:        20,
		Attack:       rand.Intn(6) + 5,
		Initiative:   3,
		PatternState: 0,
	}
}

func InitSlime() Monster {
	return Monster{
		Name:         "Slime",
		HPmax:        30,
		Attack:       3,
		Initiative:   2,
		PatternState: 0,
	}
}

func InitTavernier() Monster {
	return Monster{
		Name:         "Tavernier",
		HPmax:        40,
		Attack:       rand.Intn(6) + 7,
		Initiative:   4,
		PatternState: 0,
	}
}

func InitOurs() Monster {
	return Monster{
		Name:         "Ours",
		HPmax:        45,
		Attack:       rand.Intn(6) + 8,
		Initiative:   5,
		PatternState: 0,
	}
}

func InitEsprit() Monster {
	return Monster{
		Name:         "Esprit",
		HPmax:        25,
		Attack:       rand.Intn(4) + 5,
		Initiative:   4,
		PatternState: 0,
	}
}

func InitBucheron() Monster {
	return Monster{
		Name:         "Bucheron",
		HPmax:        42,
		Attack:       rand.Intn(6) + 7,
		Initiative:   4,
		PatternState: 0,
	}
}

func PatternSanglier(m *Monster, turn int) (int, string) {
	damage := m.Attack
	desc := m.Name + " vous met un coup de museau"

	if m.PatternState == 3 {
		damage += 4
		desc = m.Name + " s'énerve et vous met un coup chargé"
		m.PatternState = 0
	} else {
		m.PatternState++
	}
	return damage, desc
}

func PatternLoup(m *Monster, turn int) (int, string) {
	damage := m.Attack
	desc := m.Name + " vous mord"

	if m.PatternState == 3 {
		damage += 4
		desc = m.Name + " vous attrape à la gorge"
		m.PatternState = 0
	} else {
		m.PatternState++
	}
	return damage, desc
}

func PatternSlime(m *Monster, turn int) (int, string) {
	damage := m.Attack
	desc := m.Name + " se jette sur vous"

	if m.PatternState == 2 {
		damage += 4
		desc = m.Name + " vous projette ses mucosités"
		m.PatternState = 0
	} else {
		m.PatternState++
	}
	return damage, desc
}

func PatternTavernier(m *Monster, turn int) (int, string) {
	damage := m.Attack
	desc := m.Name + " vous jette une pinte de bière"

	if m.PatternState == 3 {
		damage += 4
		desc = m.Name + " vous casse une chaise sur la tête"
		m.PatternState = 0
	} else {
		m.PatternState++
	}
	return damage, desc
}

func PatternOurs(m *Monster, turn int) (int, string) {
	damage := m.Attack
	desc := m.Name + " vous mord"

	if m.PatternState == 3 {
		damage += 4
		desc = m.Name + " vous tranche avec ses griffes affûtées"
		m.PatternState = 0
	} else {
		m.PatternState++
	}
	return damage, desc
}

func PatternEsprit(m *Monster, turn int) (int, string) {
	damage := m.Attack
	desc := m.Name + " vous met une gifle"

	if m.PatternState == 3 {
		damage += 2
		desc = m.Name + " vous jette une malédiction"
		m.PatternState = 0
	} else {
		m.PatternState++
	}
	return damage, desc
}

func PatternBucheron(m *Monster, turn int) (int, string) {
	damage := m.Attack
	desc := m.Name + " vous donne un coup de manche"

	if m.PatternState == 3 {
		damage += 4
		desc = m.Name + " vous jette une hache dans la poitrine"
		m.PatternState = 0
	} else {
		m.PatternState++
	}
	return damage, desc
}
