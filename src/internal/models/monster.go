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
		Attack:       rand.Intn(4) + 5,
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
