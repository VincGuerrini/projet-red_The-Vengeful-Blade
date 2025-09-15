package models

type Item struct {
	Name        string
	Type        string
	Description string
	Price       int
	Stackable   bool
	Weight      int
	Effect      string
}

var Potion = Item{
	Name:        "Potion simple",
	Type:        "Consommable",
	Description: "Petite potion qui permet de récupérer 25 PV",
	Price:       15,
	Stackable:   false,
	Weight:      2,
	Effect:      "+25pv",
}

var Poison = Item{
	Name:        "Poison simple",
	Type:        "Consommable",
	Description: "Petite fiole d’un poison très corrosif qui retire 5 PV par tour à l’adversaire",
	Price:       10,
	Stackable:   false,
	Weight:      2,
	Effect:      "-5pv par tour",
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
	Type:        "Ressources Est Matériaux.",
	Description: "Fourrure d'anieaux depeusés.",
	Price:       10,
	Stackable:   true,
	Weight:      3,
}

var Epeev1 = Item{
	Name:        "Épee simple.",
	Type:        "Arme et armure.",
	Description: "Épee trouvé en sortant de prisont, elle parait banale et meriterais un aigisage.",
	Stackable:   false,
	Weight:      3,
}

var Lancev1 = Item{}

var Daguev1 = Item{}
