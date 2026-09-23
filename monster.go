package main

import "fmt"

type Monster struct {
	Name      string
	MaxHP     int
	CurrentHP int
	Attack    int
}

func initGoblin() Monster {
	return Monster{
		Name:      "Gobelin d'entrainement",
		MaxHP:     40,
		CurrentHP: 40,
		Attack:    5,
	}
}

// Pattern de combat du Gobelin (Tâche 20)
func goblinPattern(m *Monster, c *Character, turn int) {
	var damage int

	// Tous les 3 tours (tour 3, 6, 9...), le monstre inflige 200% de dégâts
	if turn%3 == 0 {
		damage = m.Attack * 2
		fmt.Printf("\n⚡ %s utilise un Coup Puissant !\n", m.Name)
	} else {
		damage = m.Attack
	}

	// Application des dégâts au joueur
	c.CurrentHP -= damage
	if c.CurrentHP < 0 {
		c.CurrentHP = 0
	}

	// Affichage des informations de l'attaque
	fmt.Printf("%s inflige à %s %d points de dégâts !\n", m.Name, c.Name, damage)
	fmt.Printf("%s PV : %d/%d\n", c.Name, c.CurrentHP, c.MaxHP)

	// Vérification si le joueur est mort (Tâche 8)
	if isDead(c) {
		fmt.Println("\nVous avez été vaincu en combat...")
	}
}
