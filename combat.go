package main

import "fmt"

func characterTurn(player *Character, monster *Monster) {
	fmt.Println("\n=== COMBAT ===")
	fmt.Println("1. Attaquer")
	fmt.Println("2. Utiliser un Sort (Magie)")
	fmt.Println("3. Inventaire")
	fmt.Print("Votre choix : ")

	var choice int
	fmt.Scan(&choice)

	switch choice {
	case 1:
			// Attaque basique standard
			damage := 5
			monster.CurrentHP -= damage
			if monster.CurrentHP < 0 {
				monster.CurrentHP = 0
			}
			fmt.Println("\n👊 Vous utilisez Attaque basique !")
			fmt.Printf("Dégâts infligés : %d\n", damage)
			fmt.Printf("PV de %s : %d/%d\n", monster.Name, monster.CurrentHP, monster.MaxHP)
			return // Fin du tour du joueur
	case 2:
		// Lancer un sort (Mission 3)
		if castSpell(player, monster) {
			return // Fin du tour si le sort a été lancé
		}
	case 3:
			// Utiliser un objet de l'inventaire
			accessInventoryFight(player, monster)
			return // Fin du tour du joueur
	default:
		fmt.Println("Choix invalide")
	}
}


func trainingFight(c *Character) {
	// Initialisation du Gobelin d'entraînement (Tâche 19)
	monster := initGoblin()
	turn := 1

	fmt.Println("\n=================================")
	fmt.Printf("⚔️  DEBUT DU COMBAT : %s VS %s ⚔️\n", c.Name, monster.Name)
	fmt.Println("=================================")

	// Boucle principale du combat
	for c.CurrentHP > 0 && monster.CurrentHP > 0 {
    fmt.Printf("\n--- TOUR %d ---\n", turn)

    // 1. Appliquer les dégâts de poison au monstre s'il est empoisonné
    applyPoisonDamage(&monster)
    if monster.CurrentHP <= 0 {
        fmt.Printf("\n🎉 %s a succombé au poison !\n", monster.Name)
        break
    }

    // 2. Tour du joueur
    characterTurn(c, &monster)
    if monster.CurrentHP <= 0 {
        fmt.Printf("\n🎉 Vous avez vaincu le %s !\n", monster.Name)
        break
    }

    // 3. Tour du monstre
    goblinPattern(&monster, c, turn)
    if c.CurrentHP <= 0 {
        break
    }

    turn++
}

	fmt.Println("\n=== FIN DU COMBAT ===")
}

// Pour la mission bonus 3
func castSpell(c *Character, m *Monster) bool {
	if len(c.Skill) == 0 {
		fmt.Println("\nVous ne connaissez aucun sort !")
		return false
	}

	fmt.Println("\n=== SORTS DISPONIBLES ===")
	for i, spell := range c.Skill {
		fmt.Printf("%d. %s\n", i+1, spell)
	}
	fmt.Println("0. Retour")
	fmt.Print("Choisissez un sort à lancer : ")

	var choice int
	fmt.Scan(&choice)

	if choice == 0 {
		return false
	}

	if choice > 0 && choice <= len(c.Skill) {
		spell := c.Skill[choice-1]
		damage := 0

		switch spell {
		case "Coup de poing":
			damage = 8 // Dégâts de Coup de poing (Mission 3)
		case "Boule de Feu":
			damage = 18 // Dégâts de Boule de Feu (Mission 3)
		default:
			damage = 10
		}

		m.CurrentHP -= damage
		if m.CurrentHP < 0 {
			m.CurrentHP = 0
		}

		fmt.Printf("\n✨ Vous lancez %s sur %s !\n", spell, m.Name)
		fmt.Printf("Dégâts infligés : %d\n", damage)
		fmt.Printf("%s PV : %d/%d\n", m.Name, m.CurrentHP, m.MaxHP)
		return true
	} else {
		fmt.Println("Choix invalide.")
		return false
	}
}