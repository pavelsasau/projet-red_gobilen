package main

import "fmt"

func accessInventory(player Character) {
	fmt.Println("Inventaire:")

	for _, item := range player.inventory {
		fmt.Println(item)
	}
}
