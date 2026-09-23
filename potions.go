package main 

import "fmt"

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
func poisonPot(m *Monster) {
	m.PoisonTurns = 3
	fmt.Printf("\n🧪 Vous lancez une Potion de poison sur %s !\n", m.Name)
	fmt.Printf("%s est empoisonné pour 3 tours !\n", m.Name)
}

func applyPoisonDamage(m *Monster) {
	if m.PoisonTurns > 0 {
		damage := 10
		m.CurrentHP -= damage
		if m.CurrentHP < 0 {
			m.CurrentHP = 0
		}

		m.PoisonTurns--
		fmt.Printf("\n🟢 Le poison brûle %s... (-%d PV) | PV restants : %d/%d (Tours de poison restants : %d)\n",
			m.Name, damage, m.CurrentHP, m.MaxHP, m.PoisonTurns)
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
