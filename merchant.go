package main

import "fmt"

func merchantMenu(c *Character) {
	for {
		fmt.Println("\n=== MARCHAND ===")
		fmt.Println("1. Potion de vie (Gratuit)")
		fmt.Println("0. Retour")
		fmt.Print("Votre choix : ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			addInventory(c, "Potion de vie")
		case 0:
			fmt.Println("Vous quittez le marchand.")
			return // Retour au menu principal
		default:
			fmt.Println("Choix invalide, réessayez.")
		}
	}
}
