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
	player.inventory = []string{}

	return player
}

func displayInfo(player Character) {
	fmt.Println("Name:", player.name)
	fmt.Println("Class:", player.class)
	fmt.Println("Level:", player.level)
	fmt.Println("HP:", player.hp_a, "/", player.hp_max)
}
