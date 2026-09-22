package main

import "fmt"

func accessInventory(player Character) {
	fmt.Println("Inventaire:")

	for _, item := range player.Inventory {
		fmt.Println(item)
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

// Parcourt et affiche l'inventaire
// Перебирает и показывает инвентарь
