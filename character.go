package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Character struct {
	Name      string
	Class     string
	Level     int
	MaxHP     int
	CurrentHP int
	Gold      int
	Skill     []string
	Inventory []string
	Equip     Equipment // Nouveau champ pour les équipements portés
	MaxInventory    int // Capacité max (par défaut : 10)
	UpgradeCount    int // Nombre d'améliorations déjà achetées (max : 3)
}

func characterCreation() Character {
	reader := bufio.NewReader(os.Stdin)
	var name string

	// 1. Saisie et vérification du nom
	for {
		fmt.Print("Entrez le nom de votre presonnage : ")
		input, _ := reader.ReadString('\n')
		name = strings.TrimSpace(input)

		// Verification : uniquement des lettres
		valid := true
		for _, char := range name {
			if (char < 'a' || char > 'z') && (char < 'A' || char > 'Z') {
				valid = false
				break
			}
		}
		if valid && len(name) > 0 {
			break // Le nom est valide, on sort de la boucle
		}
		fmt.Println("Nom invalide ! utilisez uniquement des lettres sans espaces ni chiffres.")
	}

	// 2. Formatage : Première lettre en Majuscule, le reste en minuscules
	name = strings.Title(strings.ToLower(name))

	// 3. Choix de la classe*
	var className string
	var MaxHP int

for {
	fmt.Println("\nChoisissez votre classe :")
	fmt.Println("1. Humain (100 PV)")
	fmt.Println("2. Elfe (80 PV)")
	fmt.Println("3. Nain (120 PV)")
	fmt.Print("Votre choix : ")

	var choice int
	fmt.Scan(&choice)

	if choice == 1 {
        className = "Humain"
        MaxHP = 100
        break // Choix valide, on sort de la boucle !
    } else if choice == 2 {
        className = "Elfe"
        MaxHP = 80
        break // Choix valide, on sort de la boucle !
    } else if choice == 3 {
        className = "Nain"
        MaxHP = 120
        break // Choix valide, on sort de la boucle !
    } else {
        fmt.Println("\nChoix invalide ! Veuillez rechoisir une classe parmi les options indiquées.")
    }
}

	// 4. Initialisation avec 50% des PV max et inventaire de base
	return Character{
		Name:      name,
		Class:     className,
		Level:     1,
		MaxHP:     MaxHP,
		CurrentHP: MaxHP / 2, // 50% des PV Max
		Gold:      100,
		Inventory: []string{"Potion de vie", "Potion de vie", "Potion de vie"},
		Skill:     []string{"Coup de poing"},
		MaxInventory: 10,
		UpgradeCount: 0,
	}
}

func displayInfo(player Character) {
	fmt.Println("Name:", player.Name)
	fmt.Println("Class:", player.Class)
	fmt.Println("Level:", player.Level)
	fmt.Println("Gold:", player.Gold)
	fmt.Println("HP:", player.CurrentHP, "/", player.MaxHP)
	fmt.Println("\n--- ÉQUIPEMENTS PORTÉS ---")
	fmt.Printf("Tête  : %s\n", showEquip(player.Equip.Head))
	fmt.Printf("Torse : %s\n", showEquip(player.Equip.Body))
	fmt.Printf("Pieds : %s\n", showEquip(player.Equip.Feet))
	fmt.Println()
}

func showEquip(item string) string {
	if item == "" {
		return "Aucun"
	}
	return item
}

// Vérifie si le personnage est mort et lui rend 50 % de ses PV maximum.
// Проверяет, умер ли персонаж, и восстанавливает ему 50% максимального здоровья.
func isDead(c *Character) bool {
	if c.CurrentHP <= 0 {
		fmt.Println("\nVous êtes mort...")
		c.CurrentHP = c.MaxHP / 2 // Résurrection à 50% des PV max (Tâche 8)
		fmt.Printf("Vous avez été ressuscité avec %d/%d PV !\n", c.CurrentHP, c.MaxHP)
		return true // Le joueur était bien mort
	}
	return false // Le joueur est encore en vie
}

type Equipment struct {
	Head string // Chapeau
	Body string // Tunique
	Feet string // Bottes
}