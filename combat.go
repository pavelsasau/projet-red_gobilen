package main

import "fmt"

func characterTurn(player *Character, monster *Monster) {
	fmt.Println("\n=== COMBAT ===")
	fmt.Println("1. Attaquer")
	fmt.Println("2. Inventaire")
	fmt.Print("Votre choix : ")

	var choice int
	fmt.Scan(&choice)

	switch choice {
	case 1:
		monster.CurrentHP -= 5

		if monster.CurrentHP < 0 {
			monster.CurrentHP = 0
		}

		fmt.Println("Attaque basique")
		fmt.Println("Dégâts infligés : 5")
		fmt.Printf("PV de %s : %d/%d\n", monster.Name, monster.CurrentHP, monster.MaxHP)

	case 2:
		accessInventory(player)

	default:
		fmt.Println("Choix invalide")
	}
}
