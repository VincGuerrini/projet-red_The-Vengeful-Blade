package models

import (
	"math/rand"
)

type Monster struct { // création de ma structure Monstre
	Name         string
	HPmax        int
	Attack       int
	Initiative   int
	PatternState int
}

// création de chaque entiter

func InitSanglier() Monster {
	return Monster{
		Name:         "Sanglier",
		HPmax:        15,
		Attack:       rand.Intn(4) + 5, // Me permet a mon entiter de faire entre 5 et 9 dommage
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
		Name:         "Bûcheron",
		HPmax:        42,
		Attack:       rand.Intn(6) + 7,
		Initiative:   4,
		PatternState: 0,
	}
}

// Je met les paternes d'attaque pour chacun de mes monstres

func PatternSanglier(monster *Monster, turn int) (damage int, description string) {
	damage = monster.Attack
	description = monster.Name + " vous met un coup de museau"

	if monster.PatternState == 3 { //
		damage = monster.Attack + 4
		description = monster.Name + " s'énèrve et vous met un coup charger"
		monster.PatternState = 0
	} else {
		monster.PatternState++
	}
	return damage, description
}

func PatternLoup(monster *Monster, turn int) (damage int, description string) {
	damage = monster.Attack
	description = monster.Name + " vous mord"

	if monster.PatternState == 3 {
		damage = monster.Attack + 4
		description = monster.Name + " vous attrape a la gorge"
		monster.PatternState = 0
	} else {
		monster.PatternState++
	}
	return damage, description
}

func PatternSlime(monster *Monster, turn int) (damage int, description string) {
	damage = monster.Attack
	description = monster.Name + " se jette sur vous"

	if monster.PatternState == 2 {
		damage = monster.Attack + 4
		description = monster.Name + " vous projette ses mecueuses"
		monster.PatternState = 0
	} else {
		monster.PatternState++
	}
	return damage, description
}

func PatternTavernier(monster *Monster, turn int) (damage int, description string) {
	damage = monster.Attack
	description = monster.Name + " vous jette une pinte de bierre"

	if monster.PatternState == 3 {
		damage = monster.Attack + 4
		description = monster.Name + " vous casse une chaise sur la tête"
		monster.PatternState = 0
	} else {
		monster.PatternState++
	}
	return damage, description
}

func PatternOurs(monster *Monster, turn int) (damage int, description string) {
	damage = monster.Attack
	description = monster.Name + " vous mord"

	if monster.PatternState == 3 {
		damage = monster.Attack + 4
		description = monster.Name + " vous tranche avec ses griffes affutées"
		monster.PatternState = 0
	} else {
		monster.PatternState++
	}
	return damage, description
}

func PatternEsprit(monster *Monster, turn int) (damage int, description string) {
	damage = monster.Attack
	description = monster.Name + " vous met une giffle"

	if monster.PatternState == 3 {
		damage = monster.Attack + 2
		description = monster.Name + " vous jette une malediction"
		monster.PatternState = 0
	} else {
		monster.PatternState++
	}
	return damage, description
}

func PatternBucheron(monster *Monster, turn int) (damage int, description string) {
	damage = monster.Attack
	description = monster.Name + " vous donne un coup de manche"

	if monster.PatternState == 3 {
		damage = monster.Attack + 4
		description = monster.Name + " vous jette une hâche dans la poitrine"
		monster.PatternState = 0
	} else {
		monster.PatternState++
	}
	return damage, description
}
