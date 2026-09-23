package main

import "fmt"

// Parcourt et affiche l'inventaire
// Перебирает и показывает инвентарь
func accessInventory(c *Character) {
	for {
		fmt.Println("\n=== INVENTAIRE ===")
		if len(c.Inventory) == 0 {
			fmt.Println("Votre inventaire est vide.")
			return
		}

		// Affichage des objets avec un numéro
		for i, item := range c.Inventory {
			fmt.Printf("%d. %s\n", i+1, item)
		}
		fmt.Println("0. Retour")
		fmt.Print("Choisissez un objet à utiliser/équiper : ")

		var choice int
		fmt.Scan(&choice)

		if choice == 0 {
			return
		}

		// Vérification du choix valide
		if choice > 0 && choice <= len(c.Inventory) {
			selectedItem := c.Inventory[choice-1]
			useItem(c, selectedItem)
		} else {
			fmt.Println("\nChoix invalide.")
		}
	}
}

// Ajoute un objet à l'inventaire
func addInventory(c *Character, item string) bool {
    if len(c.Inventory) >= 10 { // Limite d'inventaire (Tâche 12)
        fmt.Println("Inventaire plein ! Impossible d'ajouter l'objet.")
        return false // Échec de l'ajout
    }

    c.Inventory = append(c.Inventory, item)
    return true // Succès de l'ajout
}

// Retire un objet de l'inventaire
func removeInventory(c *Character, item string) bool {
	for i, slot := range c.Inventory {
		if slot == item {
			// Supprime l'élément à l'index i
			c.Inventory = append(c.Inventory[:i], c.Inventory[i+1:]...)
			return true
		}
	}
	fmt.Printf("L'objet %s n'a pas été trouvé dans l'inventaire.\n", item)
	return false
}

// Vérifie si le joueur possède un objet en quantité suffisante sans le retirer
func hasItem(c *Character, item string, count int) bool {
	found := 0
	for _, slot := range c.Inventory {
		if slot == item {
			found++
		}
	}
	return found >= count
}

// Fonction qui déclenche l'effet d'un objet selon son nom
func useItem(c *Character, item string) {
	switch item {
	case "Potion de vie":
		takePot(c)
	case "Potion de poison":
		fmt.Println("\nLa potion de poison ne peut être utilisée que pendant un combat !")
	case "Livre de Sort : Boule de Feu":
		spellBook(c)
	case "Chapeau de l'aventurier", "Tunique de l'aventurier", "Bottes de l'aventurier":
		equipItem(c, item)
	default:
		fmt.Printf("\nL'objet '%s' ne peut pas être utilisé directement.\n", item)
	}
}
