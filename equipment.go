package main

import "fmt"

func equipItem(c *Character, item string) {
	switch item {
	case "Chapeau de l'aventurier":
		// 1. Si un chapeau est déjà équipé, on le remet dans l'inventaire
		if c.Equip.Head != "" {
			c.Inventory = append(c.Inventory, c.Equip.Head)
			c.MaxHP -= 10 // On retire le bonus du chapeau précédent
		}
		// 2. On équipe le nouveau chapeau
		c.Equip.Head = item
		c.MaxHP += 10
		fmt.Println("\nVous équipez : Chapeau de l'aventurier (+10 PV Max) !")

	case "Tunique de l'aventurier":
		if c.Equip.Body != "" {
			c.Inventory = append(c.Inventory, c.Equip.Body)
			c.MaxHP -= 25 // On retire le bonus de la tunique précédente
		}
		c.Equip.Body = item
		c.MaxHP += 25
		fmt.Println("\nVous équipez : Tunique de l'aventurier (+25 PV Max) !")

	case "Bottes de l'aventurier":
		if c.Equip.Feet != "" {
			c.Inventory = append(c.Inventory, c.Equip.Feet)
			c.MaxHP -= 15 // On retire le bonus des bottes précédentes
		}
		c.Equip.Feet = item
		c.MaxHP += 15
		fmt.Println("\nVous équipez : Bottes de l'aventurier (+15 PV Max) !")

	default:
		fmt.Println("\nCet objet ne peut pas être équipé.")
		return
	}

	// On retire l'équipement de l'inventaire une fois équipé
	removeInventory(c, item)
}
