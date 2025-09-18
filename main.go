// main.go
package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"strings"
	"time"
	"unicode"

	"golang.org/x/term"
)

/*
Simple terminal RPG prototype based on the user's spec.

Controls:
- Title screen: press 'a' or 'A' to start.
- Menus: Up/Down arrows to move, Enter to validate.
- Text sequences separated visually: press Enter to advance.
*/

type Player struct {
	Name       string
	Class      string
	Weapon     string
	HP         int
	MaxHP      int
	Attack     int
	Gold       int
	Inventory  map[string]int
	WeaponTier string // "simple" or "robuste"
}

type Enemy struct {
	Name   string
	HP     int
	MaxHP  int
	Attack int
	Gold   int
	Loot   map[string]int
}

// ---------- Terminal helpers ----------
func clear() {
	// Works on Unix-like. On Windows, you can replace with "cls".
	cmd := exec.Command("/bin/sh", "-c", "clear")
	cmd.Stdout = os.Stdout
	_ = cmd.Run()
}

func gotoXY(x, y int) {
	fmt.Printf("\x1b[%d;%dH", y, x)
}

func typeWriter(s string, d time.Duration) {
	for _, r := range s {
		fmt.Printf("%c", r)
		time.Sleep(d)
	}
	fmt.Println()
}

// readKey reads one key (handles arrow keys by returning the final byte A/B/C/D)
func readKey() (rune, error) {
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		return 0, err
	}
	defer term.Restore(int(os.Stdin.Fd()), oldState)

	b := make([]byte, 3)
	n, err := os.Stdin.Read(b)
	if err != nil {
		return 0, err
	}
	if n == 1 {
		return rune(b[0]), nil
	}
	if n >= 3 && b[0] == 27 && b[1] == 91 {
		// arrows: A=up, B=down, C=right, D=left
		return rune(b[2]), nil
	}
	return rune(b[0]), nil
}

func readLinePrompt(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

func normalizeName(s string) string {
	s = strings.TrimSpace(s)
	if s == "" {
		return "Héros"
	}
	r := []rune(strings.ToLower(s))
	r[0] = unicode.ToUpper(r[0])
	return string(r)
}

// ---------- Menu helper (simple centered list) ----------
func menu(options []string) int {
	pos := 0
	for {
		clear()
		// small centering (not pixel-perfect)
		for i := 0; i < 6; i++ {
			fmt.Println()
		}
		width := 50
		for i, o := range options {
			prefix := "  "
			if i == pos {
				prefix = "> "
			}
			leftPad := (width - len(o)) / 2
			if leftPad < 0 {
				leftPad = 0
			}
			fmt.Printf("%s%s%s\n", strings.Repeat(" ", leftPad), prefix, o)
		}
		k, _ := readKey()
		if k == 'A' { // up
			if pos > 0 {
				pos--
			}
		} else if k == 'B' { // down
			if pos < len(options)-1 {
				pos++
			}
		} else if k == '\r' || k == '\n' || k == 13 {
			return pos
		}
	}
}

// ---------- Visual helpers ----------
func hpBar(current, max int, width int) string {
	if max <= 0 {
		return "[no hp]"
	}
	f := int(float64(current) / float64(max) * float64(width))
	if f < 0 {
		f = 0
	}
	if f > width {
		f = width
	}
	return "[" + strings.Repeat("#", f) + strings.Repeat(" ", width-f) + "]"
}

func labelTopRight(s string) {
	// place string at column ~60 row 1
	gotoXY(60, 1)
	fmt.Print(s)
	gotoXY(1, 1)
}

// ---------- Combat ----------
func combat(p *Player, e *Enemy) bool {
	reader := bufio.NewReader(os.Stdin)
	for p.HP > 0 && e.HP > 0 {
		clear()
		labelTopRight("COMBAT")
		fmt.Printf("%s   VS   %s\n\n", p.Name, e.Name)
		fmt.Printf("%s %d/%d\n", hpBar(p.HP, p.MaxHP, 20), p.HP, p.MaxHP)
		fmt.Printf("%s %d/%d\n\n", hpBar(e.HP, e.MaxHP, 20), e.HP, e.MaxHP)
		fmt.Println("Appuyez sur Entrée pour attaquer...")
		reader.ReadBytes('\n')

		// player attack
		dmg := rand.Intn(p.Attack/2+1) + p.Attack/2
		if dmg < 1 {
			dmg = 1
		}
		e.HP -= dmg
		fmt.Printf("Vous attaquez %s et infligez %d dégâts\n", e.Name, dmg)
		time.Sleep(600 * time.Millisecond)

		if e.HP <= 0 {
			fmt.Println("Ennemi vaincu !")
			// loot assign
			for k, v := range e.Loot {
				p.Inventory[k] += v
				fmt.Printf("Vous obtenez %s x%d\n", k, v)
			}
			if e.Gold > 0 {
				p.Gold += e.Gold
				fmt.Printf("Vous obtenez %d gold\n", e.Gold)
			}
			fmt.Println("Appuyez sur Entrée...")
			reader.ReadBytes('\n')
			return true
		}

		// enemy attack
		dmg2 := rand.Intn(e.Attack/2+1) + e.Attack/2
		if dmg2 < 1 {
			dmg2 = 1
		}
		p.HP -= dmg2
		fmt.Printf("%s vous attaque et inflige %d dégâts\n", e.Name, dmg2)
		time.Sleep(600 * time.Millisecond)

		if p.HP <= 0 {
			fmt.Println("Vous êtes mort... Game Over.")
			return false
		}
		fmt.Println("Appuyez sur Entrée pour le prochain tour...")
		reader.ReadBytes('\n')
	}
	return false
}

// ---------- Forge & Shop ----------
func forge(p *Player) {
	clear()
	labelTopRight("FORGE")
	typeWriter("Bienvenu dans ma forge...", 30*time.Millisecond)
	typeWriter("Veux-tu accéder à ma forge ou bien récupérer un cheval ?", 30*time.Millisecond)

	options := []string{"Accéder à la forge", "Récupérer le cheval", "Retour"}
	c := menu(options)
	if c == 0 {
		clear()
		labelTopRight("FORGE")
		typeWriter("Forge ouverte.", 25*time.Millisecond)
		if p.WeaponTier == "robuste" {
			typeWriter("Votre arme est déjà robuste.", 25*time.Millisecond)
			fmt.Println("Appuyez sur Entrée...")
			bufio.NewReader(os.Stdin).ReadBytes('\n')
			return
		}
		fmt.Println("Recette pour robuste: Os x1, Viande de Sanglier x1")
		fmt.Printf("Vous avez: Os x%d, Viande de Sanglier x%d\n", p.Inventory["Os"], p.Inventory["Viande de Sanglier"])
		if p.Inventory["Os"] >= 1 && p.Inventory["Viande de Sanglier"] >= 1 {
			p.Inventory["Os"]--
			p.Inventory["Viande de Sanglier"]--
			p.WeaponTier = "robuste"
			p.Attack += 3
			typeWriter("Forge réussite ! Votre arme est désormais robuste.", 30*time.Millisecond)
		} else {
			typeWriter("Vous n'avez pas les matériaux requis.", 30*time.Millisecond)
		}
		fmt.Println("Appuyez sur Entrée...")
		bufio.NewReader(os.Stdin).ReadBytes('\n')
	} else if c == 1 {
		typeWriter("Vous récupérez le cheval et arrivez en ville...", 30*time.Millisecond)
		options := []string{"Commercer", "Discussion avec PNJ-2", "Retour"}
		c2 := menu(options)
		if c2 == 0 {
			typeWriter("Entrée dans le commerce...", 25*time.Millisecond)
			shop(p)
		} else if c2 == 1 {
			typeWriter("Arrivé à l'Auberge", 25*time.Millisecond)
			typeWriter("Vous avez terminé l'acte 1 !", 25*time.Millisecond)
			typeWriter("À suivre...", 25*time.Millisecond)
			fmt.Println("Appuyez sur Entrée...")
			bufio.NewReader(os.Stdin).ReadBytes('\n')
		}
	}
}

func shop(p *Player) {
	for {
		clear()
		labelTopRight("COMMERCE")
		typeWriter("Bienvenue au shop. Vous pouvez acheter une potion de soin pour 50 gold.", 20*time.Millisecond)
		fmt.Printf("Gold: %d\n", p.Gold)
		options := []string{"Acheter Potion (50)", "Retour"}
		c := menu(options)
		if c == 0 {
			if p.Gold >= 50 {
				p.Gold -= 50
				p.Inventory["Potion"]++
				typeWriter("Achat effectué.", 20*time.Millisecond)
			} else {
				typeWriter("Pas assez d'or.", 20*time.Millisecond)
			}
			fmt.Println("Appuyez sur Entrée...")
			bufio.NewReader(os.Stdin).ReadBytes('\n')
		} else {
			return
		}
	}
}

// ---------- Game sequences ----------
func introSequence(p *Player) {
	clear()
	labelTopRight("INTRO")
	l := []string{
		"Un passé sombre...",
		"Une histoire appartenant au passé...",
		"Un petit chevalier vaincra...",
		"Aujourd'hui enfermé...",
		"Bientôt libéré...",
	}
	r := bufio.NewReader(os.Stdin)
	for _, line := range l {
		typeWriter(line, 40*time.Millisecond)
		fmt.Println("(Appuyez sur Entrée pour continuer)")
		r.ReadBytes('\n')
		clear()
		labelTopRight("INTRO")
	}

	typeWriter("Arrivé dans la salle des armes...", 40*time.Millisecond)
	fmt.Println("(Appuyez sur Entrée)")
	r.ReadBytes('\n')
	clear()
	labelTopRight("INTRO")

	classes := []string{"Humain", "Elfe", "Nain"}
	typeWriter("Choix de classe :", 30*time.Millisecond)
	choice := menu(classes)
	ch := classes[choice]
	p.Class = ch
	if ch == "Humain" {
		p.Weapon = "Epée simple"
		p.Attack = 6
	} else if ch == "Elfe" {
		p.Weapon = "Dagues simples"
		p.Attack = 5
	} else {
		p.Weapon = "Lance simple"
		p.Attack = 7
	}
	p.WeaponTier = "simple"
	p.MaxHP = 30
	p.HP = p.MaxHP
	p.Inventory = map[string]int{}
	p.Gold = 0

	typeWriter("Vous avez choisi votre arme, et êtes désormais libéré...", 30*time.Millisecond)
	fmt.Println("Appuyez sur Entrée...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
	clear()
	labelTopRight("INTRO")

	// Prairie lore with Enter between lines (mimics comma behaviour)
	lore := []string{
		"Vous arrivez dans une prairie, le monde semble différent, mais vous n'avez pas changé...",
		"Un jour...",
		"Le roi connu pour sa malfaisance vous jeta au cachot...",
		"Il fit tuer votre famille et vos proche sans aucune pitié...",
		"Et pire...",
		"Sans aucune raison...",
	}
	for _, line := range lore {
		typeWriter(line, 30*time.Millisecond)
		fmt.Println("(Appuyez sur Entrée)")
		bufio.NewReader(os.Stdin).ReadBytes('\n')
		clear()
		labelTopRight("INTRO")
	}
}

func act1(p Player) {
	// Rencontre PNJ-1
	clear()
	labelTopRight("ACTE 1")
	typeWriter(fmt.Sprintf("Vous avez choisis de faire parti des %s ainsi votre arme principale sera : %s", p.Class, p.Weapon), 25*time.Millisecond)
	fmt.Println("Appuyez sur Entrée...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')

	// Prairie loop
	for {
		clear()
		labelTopRight("ACTE 1")
		typeWriter("Vous vous réveillez au milieu d'une prairie, le soleil vous ébouillit...", 25*time.Millisecond)

		options := []string{
			"Chasser (Combat Sanglier)",
			"Parler au chien (quête)",
			"Aller à la taverne",
			"Aller voir le forgeron",
			"Quitter la partie (save & exit)",
		}
		c := menu(options)
		if c == 0 {
			e := &Enemy{Name: "Sanglier", HP: 12, MaxHP: 12, Attack: 4, Gold: 0, Loot: map[string]int{"Viande de Sanglier": 1}}
			// copy player pointer to allow HP change
			pp := &p
			combat(pp, e)
			// return spawn
			continue
		} else if c == 1 {
			// Dog quest: needs Viande de Sanglier
			if p.Inventory["Viande de Sanglier"] >= 1 {
				typeWriter("Vous donnez la Viande de Sanglier au chien.", 25*time.Millisecond)
				p.Inventory["Viande de Sanglier"]--
				p.Inventory["Os"]++
				typeWriter("Quête accomplie : Obtention d'Os x1", 25*time.Millisecond)
				fmt.Println("Appuyez sur Entrée...")
				bufio.NewReader(os.Stdin).ReadBytes('\n')
			} else {
				typeWriter("Le chien voudrait de la viande...", 25*time.Millisecond)
				fmt.Println("Appuyez sur Entrée...")
				bufio.NewReader(os.Stdin).ReadBytes('\n')
			}
		} else if c == 2 {
			tavern(&p)
		} else if c == 3 {
			forge(&p)
		} else {
			typeWriter("Sauvegarde fictive... Au revoir.", 20*time.Millisecond)
			os.Exit(0)
		}
	}
}

func tavern(p *Player) {
	clear()
	labelTopRight("ACTE 1")
	typeWriter("Vous entrez à la taverne.", 25*time.Millisecond)
	options := []string{"Parler au tavernier (quête Os)", "Insulter le Tavernier (combat)", "Sortir de la taverne"}
	c := menu(options)
	if c == 0 {
		typeWriter("Le tavernier demande: Apportez-moi 1 Os.", 25*time.Millisecond)
		if p.Inventory["Os"] >= 1 {
			p.Inventory["Os"]--
			p.Gold += 100
			typeWriter("Quête accomplie ! Vous obtenez Bourse +100 gold", 25*time.Millisecond)
		} else {
			typeWriter("Vous n'avez pas d'Os.", 25*time.Millisecond)
		}
		fmt.Println("Appuyez sur Entrée...")
		bufio.NewReader(os.Stdin).ReadBytes('\n')
	} else if c == 1 {
		e := &Enemy{Name: "Tavernier", HP: 18, MaxHP: 18, Attack: 6, Gold: 100, Loot: map[string]int{}}
		if combat(p, e) {
			typeWriter("Vous récupérez la bourse (+100 gold)", 25*time.Millisecond)
		}
	} else {
		typeWriter("Vous sortez de la taverne...", 25*time.Millisecond)
		fmt.Println("Appuyez sur Entrée...")
		bufio.NewReader(os.Stdin).ReadBytes('\n')

		clear()
		typeWriter("Un slime apparaît !", 25*time.Millisecond)
		e := &Enemy{Name: "Slime", HP: 8, MaxHP: 8, Attack: 3, Gold: 10, Loot: map[string]int{}}
		combat(p, e)

		// after combat: option continue or go to forge
		for {
			options := []string{"Continuer la chasse", "Rendre visite au forgeron", "Retour à la prairie"}
			c2 := menu(options)
			if c2 == 0 {
				r := randomEncounter()
				combat(p, r)
				fmt.Println("Appuyez sur Entrée...")
				bufio.NewReader(os.Stdin).ReadBytes('\n')
				continue
			} else if c2 == 1 {
				forge(p)
				return
			} else {
				return
			}
		}
	}
}

func randomEncounter() *Enemy {
	r := rand.Intn(100)
	if r < 55 {
		return &Enemy{Name: "Sanglier", HP: 12, MaxHP: 12, Attack: 4, Gold: 5, Loot: map[string]int{"Viande de Sanglier": 1}}
	} else if r < 85 {
		return &Enemy{Name: "Loup", HP: 14, MaxHP: 14, Attack: 6, Gold: 8, Loot: map[string]int{"Os": 1}}
	}
	return &Enemy{Name: "Slime", HP: 8, MaxHP: 8, Attack: 3, Gold: 3, Loot: map[string]int{}}
}

// ---------- Title screen & main ----------
func titleScreen() {
	clear()
	ascii := `   _____ _           __     __                     __       _          |~       
  |_   _| |__   ___  \ \   / /__ _ __   __ _  ___ / _|_   _| |        /\//      
    | | | '_ \ / _ \  \ \ / / _ \ '_ \ / _` + "`" + ` |/ _ \ |_| | | | |       /  \\XX    
    | | | | | |  __/   \ V /  __/ | | | (_| |  __/  _| |_| | |       :~~@|      
    |_|_|_| |_|\___|  _ \_/ \___|_| |_|\__, |\___|_|  \__,_|_|       :__]|.X    
   | __ )| | __ _  __| | ___           |___/                        :=====|X    
   |  _ \| |/ _` + "`" + ` |/ _` + "`" + ` |/ _ \                                        :~~@| X    
   | |_) | | (_| | (_| |  __/                                        :  ]| X    
   |____/|_|\__,_|\__,_|\___|                                        :  ]| X    
                                                         / \   |~    :  ]| X    
                                                        /   \ / \    :__]|_.    
     XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX           :   :/ / \_  :=====|    
     X  ___ _  _ _____ ___ ___ ___     _    X          :=========/ \ :~~~@|X    
     X | __| \| |_   _| _ \ __| _ \   /_\   X           :~~~~@~~/\  \:   ]|X    
     X | _|| .` + "`" + ` | | | |   / _||   /  / _ \  X           :  |   /  \ .]/\ ]|_    
     X |___|_|\_| |_| |_|_\___|_|_\ /_/ \_\ X           :  /\_/[][]\./..\ ]|_   
     X                                      X           :  [= |_MM_| [==]# ]|   
     XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX       /\  :  [-H- -- -{__i]#_]|__.
                                                   /\ \-:  [    H   ][#][#] ] ~|
                                          \  \ \  /  \ \:  [ H  H  H; I#I |]] I 
                                        \   \  \ : II \ |__[        ;     |]]   
                                      \   \      : mm  I|++  H  H  H;     |]/\  
                                    \   \        :  n  i|++   ,     ; NHN |/_-\ 
                                     \        |##      i|   _HHH__ _;_HNH |[ ##-
                                    \         |##  ~~~~==== HHHHH     :HN  [ ###`
	fmt.Println(ascii)
	fmt.Println("\nPour démarrer : Appuyer sur la touche 'a' (si d'autres touches sont appuyées, ne rien faire).")
	for {
		k, _ := readKey()
		if k == 'a' || k == 'A' {
			return
		}
	}
}

func mainMenu() int {
	clear()
	for i := 0; i < 6; i++ {
		fmt.Println()
	}
	options := []string{"Démarrer une nouvelle partie", "Quitter le jeu"}
	return menu(options)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	titleScreen()
	choice := mainMenu()
	if choice == 1 {
		typeWriter("Au revoir...", 30*time.Millisecond)
		return
	}
	clear()
	labelTopRight("INTRO")
	name := readLinePrompt("Entrez votre nom: ")
	player := Player{}
	player.Name = normalizeName(name)
	introSequence(&player)
	act1(player)
}
