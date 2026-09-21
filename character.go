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
	fmt.Println(player.name)
}

func main() {
	player := initCharacter()

	displayInfo(player)
}
