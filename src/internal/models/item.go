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
}

var Potion = Item{
	Name:        "Potion simple",
	Type:        "Consommable",
	Description: "Petite potion qui permet de récupérer 25 PV",
	Price:       15,
	Stackable:   false,
	Weight:      2,
	Effect:      +25,
}

var Poison = Item{
	Name:        "Poison simple",
	Type:        "Consommable",
	Description: "Petite fiole d’un poison très corrosif qui retire 5 PV par tour à l’adversaire",
	Price:       10,
	Stackable:   false,
	Weight:      2,
	Effect:      -5,
}

var Viandedesanglier = Item{
	Name:        "Viande de sanglier",
	Type:        "Ressources et Matériaux",
	Description: "Viande de sanglier recupéré pour le chien",
	Stackable:   false,
	Weight:      5,
}

var Os = Item{
	Name:        "Os",
	Type:        "Ressources Est Matériaux",
	Description: "Os remie par le chien, aller le donner au tavernier",
}

var Peaudesanglier = Item{
	Name:        "Peau de sanglier",
	Type:        "Ressources est matériaux",
	Description: "Peau des sangliers que l'on a dépeusé",
	Price:       5,
	Stackable:   true,
	Weight:      1,
}

var Fourure = Item{
	Name:        "Fourrure d'animeaux",
	Type:        "Ressources Est Matériaux",
	Description: "Fourrure d'anieaux depeusés",
	Price:       10,
	Stackable:   true,
	Weight:      3,
}

var Epeev1 = Item{
	Name:        "Épee simple",
	Type:        "Arme et armure",
	Description: "Épee trouvé en sortant de prison, elle parait banale et meriterais un aigisage",
	Stackable:   false,
	Weight:      3,
	Effect:      rand.Intn(3) + 3,
}

var Epeev2 = Item{
	Name:        "Épee robuste",
	Type:        "Arme et armure",
	Description: "Épee qui a été restoré grace au cuire de sanglier, elle est tranchante mais elle meriterais encore une affinage",
	Stackable:   false,
	Weight:      4,
	Effect:      rand.Intn(4) + 5,
}

var Epeev3 = Item{
	Name:        "Épee spirituelle",
	Type:        "Arme et armure",
	Description: "cette épee d'apparance si banale caché un si grand pouvoire, elle n'est pas perfect mais elle est bien plus puissante cas l'originie",
	Stackable:   false,
	Weight:      5,
	Effect:      rand.Intn(6) + 8,
}

var Épeev4 = Item{
	Name:      "Épee vengeresse",
	Type:      "Arme et armure",
	Stackable: false,
	Weight:    5,
	Effect:    50,
}

var Lancev1 = Item{
	Name:        "Lance simple",
	Type:        "Arme et armure",
	Description: "Lance banale trouvée en sortant de prison, complètement rouillée et qui mériterait un affûtage ",
	Stackable:   false,
	Weight:      3,
	Effect:      rand.Intn(3) + 3,
}

var Lancev2 = Item{
	Name:        "Lance robuste",
	Type:        "Arme et armure",
	Description: "Lance qui a eu un bonne aigisage bien merrité, elle meriterais encore une ameliration au niveau du manche",
	Stackable:   false,
	Weight:      4,
	Effect:      rand.Intn(4) + 5,
}

var Lancev3 = Item{
	Name:        "Lance spirituelle",
	Type:        "Arme et armure",
	Description: "cette lance caché un si grand pouvoire, cela est inimajinable, elle n'est pas encore parfaite mais elle est devenue si puissante",
	Stackable:   false,
	Weight:      5,
	Effect:      rand.Intn(6) + 8,
}

var Lancev4 = Item{
	Name:      "Lance vengeresse",
	Type:      "Arme et armure",
	Stackable: false,
	Weight:    5,
	Effect:    50,
}

var Daguev1 = Item{
	Name:        "Dague simple",
	Type:        "Arme et armure",
	Description: "Dague au tranchant aproximatif trouvé en sortant de prison, elle parais bien banale est mériterait un affûtage",
	Stackable:   false,
	Weight:      3,
	Effect:      rand.Intn(3) + 3,
}

var Daguev2 = Item{
	Name:        "Dague robuste",
	Type:        "Arme et armure",
	Description: "Dague au tranchant aproximatif trouvé en sortant de prison, elle parais bien banale est mériterait un affûtage",
	Stackable:   false,
	Weight:      4,
	Effect:      rand.Intn(4) + 5,
}

var Daguev3 = Item{
	Name:        "Dague spirituelle",
	Type:        "Arme et armure",
	Description: "cette dague caché un pouvoir uncomensurable c'est complettement fou, elle n'est pas parfaite mais incroyable quand meme",
	Stackable:   false,
	Weight:      5,
	Effect:      rand.Intn(6) + 8,
}

var Daguev4 = Item{
	Name:      "Dague vengeresse",
	Type:      "Arme et armure",
	Stackable: false,
	Weight:    5,
	Effect:    50,
}
