package main

import "fmt"

func accessInventory(player Character) {
	fmt.Println("Inventaire:")

	for _, item := range player.inventory {
		fmt.Println(item)
	}
}

<<<<<<< HEAD
func takePot(player *Character) {
	for _, item := range player.inventory {
		fmt.Println(item)
		// Parcourt et affiche l'inventaire
		// Перебирает и показывает инвентарь
	}
}
=======
// Ajoute un objet à l'inventaire
func addInventory(c *Character, item string) {
	c.inventory = append(c.inventory, item)
	fmt.Printf("Vous avez obtenu : %s !\n", item)
}

// Retire un objet de l'inventaire
func removeInventory(c *Character, item string) bool {
	for i, slot := range c.inventory {
		if slot == item {
			// Supprime l'élément à l'index i
			c.inventory = append(c.inventory[:i], c.inventory[i+1:]...)
			return true
		}
	}
	fmt.Printf("L'objet %s n'a pas été trouvé dans l'inventaire.\n", item)
	return false
}
>>>>>>> 540d1cdf4d8f74d8a0ec3242d9a99f2b6321e18d
