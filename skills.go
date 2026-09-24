package main

import "fmt"

// Vérifie si le personnage connaît déjà "Boule de Feu".
// Проверяет, знает ли персонаж уже "Boule de Feu".
func spellBook(c *Character, spell string) {

	// 1. Vérifier si le sort est déjà appris
	// 1. Проверяет, изучено ли заклинание уже.
	for _, s := range c.Skill {

		// Compare chaque compétence connue avec le sort demandé.
		// Сравнивает каждый известный навык с нужным заклинанием.
		if s == spell {

			// Affiche un message si le sort est déjà connu.
			// Показывает сообщение, если заклинание уже изучено.
			fmt.Printf("\nVous connaissez déjà le sort '%s' !\n", spell)

			// Arrête la fonction.
			// Останавливает функцию.
			return
		}
	}

	// 2. Ajouter le sort
	// 2. Добавляет заклинание в список навыков.
	c.Skill = append(c.Skill, spell)

	// Affiche un message confirmant l'apprentissage du sort.
	// Показывает сообщение, что заклинание успешно изучено.
	fmt.Printf("\nVous avez appris le sort : %s !\n", spell)

	// 3. Retirer le livre de l'inventaire
	// 3. Удаляет книгу заклинаний из инвентаря.
	removeInventory(c, "Livre de Sort : Boule de Feu")
}
