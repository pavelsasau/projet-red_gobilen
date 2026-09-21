package main

import "fmt"

func accessInventory(player Character) {
	fmt.Println("Inventaire:")

	for _, item := range player.inventory {
		fmt.Println(item)
	}
}

func takePot(player *Character) {
	for _, item := range player.inventory {
		fmt.Println(item)
		// Parcourt et affiche l'inventaire
		// Перебирает и показывает инвентарь
	}
}
