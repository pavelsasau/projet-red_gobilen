package main

import "fmt"

func merchantMenu(c *Character) {
	for {
		fmt.Println("\n=== MARCHAND ===")
		fmt.Printf("Votre bourse : %d pièces d'or\n", c.Gold)
		fmt.Println("1. Potion de vie (3 pièces d'or)")
		fmt.Println("2. Potion de poison (6 pièces d’or)")
		fmt.Println("3. Livre de Sort : Boule de Feu (25 pièces d’or)")
		fmt.Println("4. Fourrure de Loup (4 pièces d’or)")
		fmt.Println("5. Peau de Troll (7 pièces d’or)")
		fmt.Println("6. Cuir de Sanglier (3 pièces d’or)")
		fmt.Println("7. Plume de Corbeau (1 pièce d’or)")
		fmt.Println("8. Augmentation d'inventaire (30 pièces d'or)")
		fmt.Println("0. Retour")
		fmt.Print("Votre choix : ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			buyItems(c, "Potion de vie", 3)
		case 2:
			buyItems(c, "Potion de poison", 6)
		case 3:
			buyItems(c, "Livre de Sort : Boule de Feu", 25)
		case 4:
			buyItems(c, "Fourrure de Loup", 4)
		case 5:
			buyItems(c, "Peau de Troll", 7)
		case 6:
			buyItems(c, "Cuir de Sanglier", 3)
		case 7:
			buyItems(c, "Plume de Corbeau", 1)
		case 8:
			buyUpgrade(c, 30)
		case 0:
			fmt.Println("Vous quittez le marchand.")
			return // Retour au menu principal
		default:
			fmt.Println("Choix invalide, réessayez.")
		}
	}
}

// Fonction générique qui gère l'achat, la déduction d'or et l'ajout à l'inventaire
func buyItems(c *Character, itemName string, price int) {
	// 1. Vérification si le joueur a assez d'or
	if c.Gold < price {
		fmt.Printf("\nVous n'avez pas assez d'or pour acheter %s ! (Prix : %d Gold, Solde : %d Gold)\n", itemName, price, c.Gold)
		return
	}
	// 2. Ajout de l'objet à l'inventaire
	if addInventory(c, itemName) {
		c.Gold -= price // 3. Déduction de l'or si l'objet peut être ajouter a l'inventaire
		// 4. Confirmation de l'achat
		fmt.Printf("\nAchat réussi : %s pour %d pièces d'or ! Solde restant : %d Gold.\n", itemName, price, c.Gold)
	} else {
		fmt.Println("L'achat a été annulé car votre inventaire est plein.")
	}
}

func buyUpgrade(c *Character, price int) {
	if c.UpgradeCount >= 3 {
		fmt.Println("\nLe marchand vous dit : 'Je ne peux plus agrandir votre sac, vous avez atteint la limite (3/3) !'")
		return
	}

	if c.Gold < price {
		fmt.Printf("\nVous n'avez pas assez d'or ! (Prix : %d Gold, Solde : %d Gold)\n", price, c.Gold)
		return
	}

	c.Gold -= price
	upgradeInventorySlot(c)
}
