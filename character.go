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

	// 3. Choix de la classe
	fmt.Println("\nChoisissez votre classe :")
	fmt.Println("1. Humain (100 PV)")
	fmt.Println("2. Elfe (80 PV)")
	fmt.Println("3. Nain (120 PV)")
	fmt.Print("Votre choix : ")

	var choice int
	fmt.Scan(&choice)

	MaxHP := 100
	className := "Humain"

	switch choice {
	case 2:
		className = "Elfe"
		MaxHP = 80
	case 3:
		className = "Nain"
		MaxHP = 120
	default:
		fmt.Println("Choix par défaut : Humain")
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
	}
}

func displayInfo(player Character) {
	fmt.Println("Name:", player.Name)
	fmt.Println("Class:", player.Class)
	fmt.Println("Level:", player.Level)
	fmt.Println("Gold:", player.Gold)
	fmt.Println("HP:", player.CurrentHP, "/", player.MaxHP)
	fmt.Println()
}

// Vérifie si le personnage est mort et lui rend 50 % de ses PV maximum.
// Проверяет, умер ли персонаж, и восстанавливает ему 50% максимального здоровья.
func isDead(player *Character) {
	if player.CurrentHP <= 0 {
		player.CurrentHP = player.MaxHP / 2

		fmt.Println("Le personnage est mort.")
		fmt.Println("Il revient avec", player.CurrentHP, "HP.")
	}
}
