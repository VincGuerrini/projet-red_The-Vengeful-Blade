package models

import (
	"math/rand"
)

type Item struct {
	Name        string
	Type        string
	Description string
	Price       int
	Stackable   bool
	Weight      int
	Effect      int
	Qty         int
}

var Potion = Item{
	Name:        "Potion simple",
	Type:        "Consommable",
	Description: "Petite potion qui permet de récupérer 25 PV",
	Price:       15,
	Stackable:   true,
	Weight:      2,
	Effect:      25,
}

var Poison = Item{
	Name:        "Poison simple",
	Type:        "Consommable",
	Description: "Fiole de poison qui retire 5 PV par tour à l’adversaire",
	Price:       10,
	Stackable:   true,
	Weight:      2,
	Effect:      -5,
}

var ViandeDeSanglier = Item{
	Name:        "Viande de sanglier",
	Type:        "Ressources et Matériaux",
	Description: "Viande de sanglier récupérée pour le chien",
	Price:       0,
	Stackable:   true,
	Weight:      5,
}

var Os = Item{
	Name:        "Os",
	Type:        "Ressources et Matériaux",
	Description: "Os remis par le chien, à donner au tavernier",
	Price:       0,
	Stackable:   true,
}

var PeauDeSanglier = Item{
	Name:        "Peau de sanglier",
	Type:        "Ressources et Matériaux",
	Description: "Peau de sanglier dépeçée",
	Price:       5,
	Stackable:   true,
	Weight:      1,
}

var Fourrure = Item{
	Name:        "Fourrure d'animaux",
	Type:        "Ressources et Matériaux",
	Description: "Fourrure d'animaux dépeçés",
	Price:       10,
	Stackable:   true,
	Weight:      3,
}

var EpeeV1 = Item{
	Name:        "Épée simple",
	Type:        "Arme et armure",
	Description: "Épée trouvée en sortant de prison, elle paraît banale et mériterait un aiguisage",
	Stackable:   false,
	Weight:      3,
	Effect:      rand.Intn(3) + 3,
}

var EpeeV2 = Item{
	Name:        "Épée robuste",
	Type:        "Arme et armure",
	Description: "Épée restaurée grâce au cuir de sanglier, tranchante mais encore améliorable",
	Stackable:   false,
	Weight:      4,
	Effect:      rand.Intn(4) + 5,
}

var EpeeV3 = Item{
	Name:        "Épée spirituelle",
	Type:        "Arme et armure",
	Description: "Épée banale mais cachant un grand pouvoir, pas parfaite mais puissante",
	Stackable:   false,
	Weight:      5,
	Effect:      rand.Intn(6) + 8,
}

var EpeeV4 = Item{
	Name:      "Épée vengeresse",
	Type:      "Arme et armure",
	Stackable: false,
	Weight:    5,
	Effect:    50,
}

var LanceV1 = Item{
	Name:        "Lance simple",
	Type:        "Arme et armure",
	Description: "Lance banale trouvée en sortant de prison, rouillée et mériterait un affûtage",
	Stackable:   false,
	Weight:      3,
	Effect:      rand.Intn(3) + 3,
}

var LanceV2 = Item{
	Name:        "Lance robuste",
	Type:        "Arme et armure",
	Description: "Lance aiguisée grâce au cuir de sanglier, encore améliorable",
	Stackable:   false,
	Weight:      4,
	Effect:      rand.Intn(4) + 5,
}

var LanceV3 = Item{
	Name:        "Lance spirituelle",
	Type:        "Arme et armure",
	Description: "Lance puissante, pas encore parfaite",
	Stackable:   false,
	Weight:      5,
	Effect:      rand.Intn(6) + 8,
}

var LanceV4 = Item{
	Name:      "Lance vengeresse",
	Type:      "Arme et armure",
	Stackable: false,
	Weight:    5,
	Effect:    50,
}

var DagueV1 = Item{
	Name:        "Dague simple",
	Type:        "Arme et armure",
	Description: "Dague approximative, mériterait un affûtage",
	Stackable:   false,
	Weight:      3,
	Effect:      rand.Intn(3) + 3,
}

var DagueV2 = Item{
	Name:        "Dague robuste",
	Type:        "Arme et armure",
	Description: "Dague légèrement améliorée mais encore perfectible",
	Stackable:   false,
	Weight:      4,
	Effect:      rand.Intn(4) + 5,
}

var DagueV3 = Item{
	Name:        "Dague spirituelle",
	Type:        "Arme et armure",
	Description: "Dague puissante mais pas parfaite",
	Stackable:   false,
	Weight:      5,
	Effect:      rand.Intn(6) + 8,
}

var DagueV4 = Item{
	Name:      "Dague vengeresse",
	Type:      "Arme et armure",
	Stackable: false,
	Weight:    5,
	Effect:    50,
}
