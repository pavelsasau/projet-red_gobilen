package main

import "fmt"

func characterTurn(player *Character, monster *Monster) {
	fmt.Println("\n=== COMBAT ===")
	fmt.Println("1. Attaquer")
	fmt.Println("2. Inventaire")
	fmt.Print("Votre choix : ")

	var choice int
	fmt.Scan(&choice)

	switch choice {
	case 1:
		monster.CurrentHP -= 5

		if monster.CurrentHP < 0 {
			monster.CurrentHP = 0
		}

		fmt.Println("Attaque basique")
		fmt.Println("Dégâts infligés : 5")
		fmt.Printf("PV de %s : %d/%d\n", monster.Name, monster.CurrentHP, monster.MaxHP)

	case 2:
		accessInventoryFight(player, monster)

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