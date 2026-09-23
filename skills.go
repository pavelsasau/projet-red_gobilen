package main

import "fmt"

// Vérifie si le personnage connaît déjà "Boule de Feu".
// Проверяет, знает ли персонаж уже "Boule de Feu".
func spellBook(c *Character, spell string) {
	// 1. Vérifier si le sort est déjà appris
	for _, s := range c.Skill {
		if s == spell {
			fmt.Printf("\nVous connaissez déjà le sort '%s' !\n", spell)
			return
		}
	}

	// 2. Ajouter le sort
	c.Skill = append(c.Skill, spell)
	fmt.Printf("\nVous avez appris le sort : %s !\n", spell)

	// 3. Retirer le livre de l'inventaire
	removeInventory(c, "Livre de Sort : Boule de Feu")
}
