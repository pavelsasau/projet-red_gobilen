package main 

import (
	"fmt"
	"time"
)
//Posion de Vie
func takePot(player *Character) {
	for i, item := range player.Inventory {

		if item == "Potion de vie" {
			player.CurrentHP = player.CurrentHP + 50

			if player.CurrentHP > player.MaxHP {
				player.CurrentHP = player.MaxHP
			}

			player.Inventory = append(
				player.Inventory[:i],
				player.Inventory[i+1:]...,
			)

			return
		}
	}
}
// Posion de Poison
// Applique l'effet visuel du poison sur le monstre pendant 3 secondes
func poisonPot(m *Monster){
	fmt.Printf("Vous lancer une potion de poison sur %s !\n", m.Name)

	// Boucle de 3 tours (1 fois par seconde)
	for i := 1; i <= 3; i++ {
		time.Sleep(1 * time.Second) // Pause d'une seconde

		m.CurrentHP -= 10 // Retire 10 PV au MONSTRE
		// Sécurité pour ne pas descendre en dessous de 0 PV
		if m.CurrentHP < 0 {
			m.CurrentHP = 0
		}

		fmt.Printf("Le poison brûle %s... (-10) | PV restant : %d/%d\n", m.Name, m.CurrentHP, m.MaxHP)
	}
}

// Fonction principale pour UTILISER la potion : retire l'objet ET lance l'effet
func usePoisonPotion(c *Character, m *Monster){
	// 1. On tente de retirer la potion de l'inventaire du joueur
	if removeInventory(c, "Potion de poison"){
		// 2. Si le retrait a réussi, on applique les dégâts au monstre
		poisonPot(m)	
	} else {
		// 3. Si le joueur n'en a pas, on affiche une erreur
		fmt.Println("Vous n'avez pas de Potion de poison dans votre inventaire !")
	}
}
