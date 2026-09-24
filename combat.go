package main

import "fmt"

// Gère le tour du joueur pendant le combat.
// Управляет ходом игрока во время боя.
func characterTurn(player *Character, monster *Monster) {
	// Affiche les actions disponibles pendant le combat.
	// Показывает доступные действия во время боя.
	fmt.Println("\n=== COMBAT ===")
	fmt.Println("1. Attaquer")
	fmt.Println("2. Utiliser un Sort (Magie)")
	fmt.Println("3. Inventaire")
	fmt.Print("Votre choix : ")

	// Stocke le choix du joueur.
	// Хранит выбор игрока.
	var choice int
	fmt.Scan(&choice)

	// Exécute une action selon le choix du joueur.
	// Выполняет действие в зависимости от выбора игрока.
	switch choice {
	case 1:
		// Attaque basique standard
		// Обычная базовая атака.
		damage := 5

		// Retire les dégâts des PV actuels du monstre.
		// Вычитает урон из текущего здоровья монстра.
		monster.CurrentHP -= damage

		// Empêche les PV du monstre de devenir négatifs.
		// Не позволяет здоровью монстра стать отрицательным.
		if monster.CurrentHP < 0 {
			monster.CurrentHP = 0
		}

		// Affiche les informations de l'attaque.
		// Показывает информацию об атаке.
		fmt.Println("\n👊 Vous utilisez Attaque basique !")
		fmt.Printf("Dégâts infligés : %d\n", damage)
		fmt.Printf("PV de %s : %d/%d\n", monster.Name, monster.CurrentHP, monster.MaxHP)

		return // Fin du tour du joueur
		// Конец хода игрока.

	case 2:
		// Lancer un sort (Mission 3)
		// Использование заклинания.
		if castSpell(player, monster) {
			return // Fin du tour si le sort a été lancé
			// Конец хода, если заклинание было использовано.
		}

	case 3:
		// Utiliser un objet de l'inventaire
		// Использование предмета из инвентаря.
		accessInventoryFight(player, monster)

		return // Fin du tour du joueur
		// Конец хода игрока.

	default:
		// Affiche un message si le choix n'est pas valide.
		// Показывает сообщение, если выбор неправильный.
		fmt.Println("Choix invalide")
	}
}

// Lance un combat d'entraînement contre un gobelin.
// Запускает тренировочный бой против гоблина.
func trainingFight(c *Character) {
	// Initialisation du Gobelin d'entraînement (Tâche 19)
	// Создание тренировочного гоблина.
	monster := initGoblin()

	// Commence le combat au tour numéro 1.
	// Начинает бой с первого хода.
	turn := 1

	// Affiche le début du combat.
	// Показывает начало боя.
	fmt.Println("\n=================================")
	fmt.Printf("⚔️  DEBUT DU COMBAT : %s VS %s ⚔️\n", c.Name, monster.Name)
	fmt.Println("=================================")

	// Boucle principale du combat
	// Основной цикл боя.
	for c.CurrentHP > 0 && monster.CurrentHP > 0 {
		fmt.Printf("\n--- TOUR %d ---\n", turn)

		// 1. Appliquer les dégâts de poison au monstre s'il est empoisonné
		// 1. Наносит урон от яда, если монстр отравлен.
		applyPoisonDamage(&monster)

		// Vérifie si le monstre est mort à cause du poison.
		// Проверяет, умер ли монстр от яда.
		if monster.CurrentHP <= 0 {
			fmt.Printf("\n🎉 %s a succombé au poison !\n", monster.Name)
			break
		}

		// 2. Tour du joueur
		// 2. Ход игрока.
		characterTurn(c, &monster)

		// Vérifie si le joueur a vaincu le monstre.
		// Проверяет, победил ли игрок монстра.
		if monster.CurrentHP <= 0 {
			fmt.Printf("\n🎉 Vous avez vaincu le %s !\n", monster.Name)
			break
		}

		// 3. Tour du monstre
		// 3. Ход монстра.
		goblinPattern(&monster, c, turn)

		// Vérifie si le joueur n'a plus de PV.
		// Проверяет, закончились ли у игрока очки здоровья.
		if c.CurrentHP <= 0 {
			break
		}

		// Passe au tour suivant.
		// Переходит к следующему ходу.
		turn++
	}

	// Affiche la fin du combat.
	// Показывает конец боя.
	fmt.Println("\n=== FIN DU COMBAT ===")
}

// Pour la mission bonus 3
// Для бонусного задания 3.

// Permet au joueur de choisir et lancer un sort.
// Позволяет игроку выбрать и использовать заклинание.
func castSpell(c *Character, m *Monster) bool {

	// Vérifie si le personnage connaît au moins une compétence.
	// Проверяет, знает ли персонаж хотя бы один навык.
	if len(c.Skill) == 0 {
		fmt.Println("\nVous ne connaissez aucun sort !")
		return false
	}

	// Affiche la liste des sorts disponibles.
	// Показывает список доступных заклинаний.
	fmt.Println("\n=== SORTS DISPONIBLES ===")

	// Parcourt toutes les compétences du personnage.
	// Перебирает все навыки персонажа.
	for i, spell := range c.Skill {
		fmt.Printf("%d. %s\n", i+1, spell)
	}

	fmt.Println("0. Retour")
	fmt.Print("Choisissez un sort à lancer : ")

	// Stocke le choix du joueur.
	// Хранит выбор игрока.
	var choice int
	fmt.Scan(&choice)

	// Retourne au menu précédent si le joueur choisit 0.
	// Возвращает назад, если игрок выбирает 0.
	if choice == 0 {
		return false
	}

	// Vérifie que le numéro choisi correspond à un sort existant.
	// Проверяет, что выбранный номер соответствует существующему навыку.
	if choice > 0 && choice <= len(c.Skill) {

		// Récupère le sort choisi dans la liste.
		// Получает выбранное заклинание из списка.
		spell := c.Skill[choice-1]

		// Variable qui contient les dégâts du sort.
		// Переменная, в которой хранится урон заклинания.
		damage := 0

		// Définit les dégâts selon le sort choisi.
		// Устанавливает урон в зависимости от выбранного заклинания.
		switch spell {
		case "Coup de poing":
			damage = 8 // Dégâts de Coup de poing (Mission 3)
			// Урон от «Удара кулаком».

		case "Boule de Feu":
			damage = 18 // Dégâts de Boule de Feu (Mission 3)
			// Урон от «Огненного шара».

		default:
			// Utilise 10 dégâts pour une compétence non prévue.
			// Использует 10 урона для другого навыка.
			damage = 10
		}

		// Retire les dégâts des PV du monstre.
		// Вычитает урон из здоровья монстра.
		m.CurrentHP -= damage

		// Empêche les PV du monstre de devenir négatifs.
		// Не позволяет здоровью монстра стать отрицательным.
		if m.CurrentHP < 0 {
			m.CurrentHP = 0
		}

		// Affiche le sort utilisé et les dégâts infligés.
		// Показывает использованное заклинание и нанесённый урон.
		fmt.Printf("\n✨ Vous lancez %s sur %s !\n", spell, m.Name)
		fmt.Printf("Dégâts infligés : %d\n", damage)
		fmt.Printf("%s PV : %d/%d\n", m.Name, m.CurrentHP, m.MaxHP)

		return true
	} else {
		// Affiche un message si le numéro choisi est incorrect.
		// Показывает сообщение, если выбран неправильный номер.
		fmt.Println("Choix invalide.")
		return false
	}
}
