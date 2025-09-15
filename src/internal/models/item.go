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
	Name:        "Potion Simple",
	Type:        "Consommable",
	Description: "Petite Potion Qui Permet De Recupéré 25 PV ",
	Price:       15,
	Stackable:   false,
	Weight:      2,
	Effect:      "+25pv",
}

var Poison = Item{
	Name:        "Poison Simple",
	Type:        "Consommable",
	Description: "Petite Fiole D’un Poison Très Corrosif Qui Retire 5 PV Par Tour à l’Adversaire",
	Price:       10,
	Stackable:   false,
	Weight:      2,
	Effect:      "-5pv par tour",
}

var Viandedesanglier = Item{
	Name:        "Viande de sanglier",
	Type:        "Ressources Est Matériaux",
	Description: "Viande de sanglier recupéré pour le chien",
	Stackable:   false,
	Weight:      5,
}

var Os = Item{
	Name:        "Os",
	Type:        "Ressources Est Matériaux",
	Description: "os remie par le chien, allé le donné au tavernier",
}

var Peaudesanglier = Item{
	Name:        "Peau De Sanglier",
	Type:        "Ressources Est Matériaux ",
	Description: "Peau Des Sangliers Que L'on à Depeusés",
	Price:       5,
	Stackable:   true,
	Weight:      1,
}

var Fourure = Item{
	Name:        "fourure d'animeaux",
	Type:        "Ressources Est Matériaux ",
	Description: "fourure d'anieaux depeusés",
	Price:       10,
	Stackable:   true,
	Weight:      3,
}

var Epeev1 = Item{
	Name:        "Épee simple",
	Type:        "arme et armure",
	Description: "Épee trouvé en sortant de prisont, elle parait banale et meriterais un aigisage.",
	Stackable:   false,
	Weight:      3,
}

var Lancev1 = Item{}

var Daguev1 = Item{}
