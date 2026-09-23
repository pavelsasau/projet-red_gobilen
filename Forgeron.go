package main

import "fmt"

func blacksmithMenu(c *Character) {
	for {
		fmt.Println("\n=== FORGERON ===")
		fmt.Printf("Votre bourse : %d pièces d'or\n", c.Gold)
		fmt.Println("1. Chapeau de l'aventurier (1 Plume de Corbeau, 1 Cuir de Sanglier + 5 Gold)")
		fmt.Println("2. Tunique de l'aventurier (2 Fourrure de Loup, 1 Peau de Troll + 5 Gold)")
		fmt.Println("3. Bottes de l'aventurier (1 Fourrure de Loup, 1 Cuir de Sanglier + 5 Gold)")
		fmt.Println("0. Retour")
		fmt.Print("Votre choix : ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			craftItem(c, "Chapeau de l'aventurier", map[string]int{"Plume de Corbeau": 1, "Cuir de Sanglier": 1})
		case 2:
			craftItem(c, "Tunique de l'aventurier", map[string]int{"Fourrure de Loup": 2, "Peau de Troll": 1})
		case 3:
			craftItem(c, "Bottes de l'aventurier", map[string]int{"Fourrure de Loup": 1, "Cuir de Sanglier": 1})
		case 0:
			fmt.Println("Vous quittez le Forgeron.")
			return // Retour au menu principal
		default:
			fmt.Println("Choix invalide, réessayez.")
		}
	}
}

func craftItem(c *Character, itemName string, recipe map[string]int) {
	// 1. Vérification de l'or
	if c.Gold < 5 {
		fmt.Println("\nVous n'avez pas assez d'or pour forger ! (Requis : 5 Gold)")
		return
	}
	// 2. Vérification des matériaux
	missing := false
	for item, count := range recipe {
		if !hasItem(c, item, count) {
			fmt.Printf("\nIl vous manque : %d x %s !", count, item)
			missing = true
		}
	}

	// Si au moins un matériau manque, on arrête la fabrication
	if missing {
		fmt.Println()
		return
	}
	// 3. Retrait des matériaux de l'inventaire
	for item, count := range recipe {
		for i := 0; i < count; i++ {
			removeInventory(c, item)
		}
	}
	// 4. Retrait de l'or
	c.Gold -= 5
	// 5. Ajout de l'équipement fabriqué
	if !addInventory(c, itemName) {
		// En cas d'inventaire plein au dernier moment, remboursement
		c.Gold += 5
		fmt.Println("Fabrication annulée : inventaire plein.")
	} else {
		fmt.Printf("\nFabrication réussie ! Vous obtenez : %s !\n", itemName)
	}
}
