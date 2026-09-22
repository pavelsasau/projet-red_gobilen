package main

import "fmt"

type Character struct {
	name      string
	class     string
	level     int
	hp_max    int
	hp_a      int
	inventory []string
}

func initCharacter() Character {
	player := Character{}

	player.name = "test"
	player.class = "Humains"
	player.level = 1
	player.hp_max = 100
	player.hp_a = 100
	player.inventory = []string{
		"Potion de vie",
		"Potion de vie",
		"Potion de vie",
	}

	return player
}

func displayInfo(player Character) {
	fmt.Println("Name:", player.name)
	fmt.Println("Class:", player.class)
	fmt.Println("Level:", player.level)
	fmt.Println("HP:", player.hp_a, "/", player.hp_max)
	fmt.Println()
}

// Vérifie si le personnage est mort et lui rend 50 % de ses PV maximum.
// Проверяет, умер ли персонаж, и восстанавливает ему 50% максимального здоровья.
func isDead(player *Character) {
	if player.hp_a <= 0 {
		player.hp_a = player.hp_max / 2

		fmt.Println("Le personnage est mort.")
		fmt.Println("Il revient avec", player.hp_a, "HP.")
	}
}
