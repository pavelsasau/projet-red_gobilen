package main

import "fmt"

// Définit les informations d'un monstre.
// Определяет данные, которые хранятся у монстра.
type Monster struct {
	Name        string
	MaxHP       int
	CurrentHP   int
	Attack      int
	PoisonTurns int // Nombre de tours de poison restants
	// Количество оставшихся ходов действия яда.
}

// Crée et retourne un Gobelin d'entraînement.
// Создаёт и возвращает тренировочного гоблина.
func initGoblin() Monster {

	// Retourne un monstre avec ses statistiques de départ.
	// Возвращает монстра с его начальными характеристиками.
	return Monster{
		Name:      "Gobelin d'entrainement",
		MaxHP:     40,
		CurrentHP: 40,
		Attack:    5,
	}
}

// Pattern de combat du Gobelin (Tâche 20)
// Логика атаки гоблина (Задание 20).
func goblinPattern(m *Monster, c *Character, turn int) {

	// Variable qui contient les dégâts du monstre.
	// Переменная, в которой хранится урон монстра.
	var damage int

	// Tous les 3 tours (tour 3, 6, 9...), le monstre inflige 200% de dégâts
	// Каждые 3 хода (3, 6, 9...) монстр наносит двойной урон.
	if turn%3 == 0 {

		// Double les dégâts de l'attaque du monstre.
		// Удваивает урон атаки монстра.
		damage = m.Attack * 2

		// Affiche l'utilisation d'une attaque puissante.
		// Показывает использование сильной атаки.
		fmt.Printf("\n⚡ %s utilise un Coup Puissant !\n", m.Name)

	} else {

		// Utilise les dégâts normaux du monstre.
		// Использует обычный урон монстра.
		damage = m.Attack
	}

	// Application des dégâts au joueur
	// Нанесение урона игроку.
	c.CurrentHP -= damage

	// Empêche les PV du joueur de devenir négatifs.
	// Не позволяет здоровью игрока стать отрицательным.
	if c.CurrentHP < 0 {
		c.CurrentHP = 0
	}

	// Affichage des informations de l'attaque
	// Показывает информацию об атаке.
	fmt.Printf("%s inflige à %s %d points de dégâts !\n", m.Name, c.Name, damage)
	fmt.Printf("%s PV : %d/%d\n", c.Name, c.CurrentHP, c.MaxHP)

	// Vérification si le joueur est mort (Tâche 8)
	// Проверяет, умер ли игрок (Задание 8).
	if isDead(c) {
		fmt.Println("\nVous avez été vaincu en combat...")
	}
}
