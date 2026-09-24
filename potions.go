package main

import "fmt"

//Posion de Vie
// Зелье здоровья.
func takePot(player *Character) {

	// Parcourt tous les objets de l'inventaire.
	// Перебирает все предметы в инвентаре.
	for i, item := range player.Inventory {

		// Vérifie si l'objet actuel est une Potion de vie.
		// Проверяет, является ли текущий предмет зельем здоровья.
		if item == "Potion de vie" {

			// Ajoute 50 PV au personnage.
			// Добавляет персонажу 50 единиц здоровья.
			player.CurrentHP = player.CurrentHP + 50

			// Vérifie que les PV ne dépassent pas les PV maximum.
			// Проверяет, чтобы здоровье не превышало максимальное.
			if player.CurrentHP > player.MaxHP {

				// Limite les PV actuels aux PV maximum.
				// Ограничивает текущее здоровье максимальным значением.
				player.CurrentHP = player.MaxHP
			}

			// Supprime la potion utilisée de l'inventaire.
			// Удаляет использованное зелье из инвентаря.
			player.Inventory = append(
				player.Inventory[:i],
				player.Inventory[i+1:]...,
			)

			// Arrête la fonction après l'utilisation de la potion.
			// Останавливает функцию после использования зелья.
			return
		}
	}
}

// Posion de Poison
// Зелье яда.

// Applique l'effet visuel du poison sur le monstre pendant 3 secondes
// Накладывает эффект яда на монстра.
func poisonPot(m *Monster) {

	// Définit la durée du poison à 3 tours.
	// Устанавливает длительность яда на 3 хода.
	m.PoisonTurns = 3

	// Affiche un message indiquant que le monstre est empoisonné.
	// Показывает сообщение о том, что монстр отравлен.
	fmt.Printf("\n🧪 Vous lancez une Potion de poison sur %s !\n", m.Name)
	fmt.Printf("%s est empoisonné pour 3 tours !\n", m.Name)
}

// Applique les dégâts du poison à chaque tour.
// Наносит урон от яда каждый ход.
func applyPoisonDamage(m *Monster) {

	// Vérifie si le poison est encore actif.
	// Проверяет, действует ли ещё яд.
	if m.PoisonTurns > 0 {

		// Définit les dégâts du poison à 10 PV.
		// Устанавливает урон яда в 10 единиц здоровья.
		damage := 10

		// Retire les dégâts des PV actuels du monstre.
		// Вычитает урон из текущего здоровья монстра.
		m.CurrentHP -= damage

		// Empêche les PV du monstre de devenir négatifs.
		// Не позволяет здоровью монстра стать отрицательным.
		if m.CurrentHP < 0 {
			m.CurrentHP = 0
		}

		// Réduit d'un tour la durée restante du poison.
		// Уменьшает оставшуюся длительность яда на один ход.
		m.PoisonTurns--

		// Affiche les dégâts et le nombre de tours de poison restants.
		// Показывает урон и количество оставшихся ходов яда.
		fmt.Printf("\n🟢 Le poison brûle %s... (-%d PV) | PV restants : %d/%d (Tours de poison restants : %d)\n",
			m.Name, damage, m.CurrentHP, m.MaxHP, m.PoisonTurns)
	}
}

// Fonction principale pour UTILISER la potion : retire l'objet ET lance l'effet
// Главная функция для ИСПОЛЬЗОВАНИЯ зелья: удаляет предмет и запускает его эффект.
func usePoisonPotion(c *Character, m *Monster) {

	// 1. On tente de retirer la potion de l'inventaire du joueur
	// 1. Пытаемся удалить зелье яда из инвентаря игрока.
	if removeInventory(c, "Potion de poison") {

		// 2. Si le retrait a réussi, on applique les dégâts au monstre
		// 2. Если зелье удалось удалить, применяем эффект яда к монстру.
		poisonPot(m)

	} else {

		// 3. Si le joueur n'en a pas, on affiche une erreur
		// 3. Если у игрока нет такого зелья, показываем ошибку.
		fmt.Println("Vous n'avez pas de Potion de poison dans votre inventaire !")
	}
}
