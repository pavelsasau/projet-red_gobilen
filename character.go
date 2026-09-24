package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Définit les informations du personnage.
// Определяет данные, которые хранятся у персонажа.
type Character struct {
	Name      string
	Class     string
	Level     int
	MaxHP     int
	CurrentHP int
	Gold      int
	Skill     []string
	Inventory []string
	Equip     Equipment // Nouveau champ pour les équipements portés
	// Новое поле для надетой экипировки.
	MaxInventory int // Capacité max (par défaut : 10)
	// Максимальная вместимость инвентаря (по умолчанию 10).
	UpgradeCount int // Nombre d'améliorations déjà achetées (max : 3)
	// Количество уже купленных улучшений (максимум 3).

}

// Crée un personnage et initialise ses informations.
// Создаёт персонажа и задаёт его начальные данные.
func characterCreation() Character {

	// Crée un lecteur pour récupérer le texte entré par l'utilisateur.
	// Создаёт средство чтения текста, введённого пользователем.
	reader := bufio.NewReader(os.Stdin)

	// Stocke le nom du personnage.
	// Хранит имя персонажа.
	var name string

	// 1. Saisie et vérification du nom
	// 1. Ввод и проверка имени.
	for {
		fmt.Print("Entrez le nom de votre presonnage : ")

		// Lit le texte jusqu'à ce que l'utilisateur appuie sur Entrée.
		// Читает текст до нажатия Enter.
		input, _ := reader.ReadString('\n')

		// Supprime les espaces inutiles au début et à la fin.
		// Удаляет лишние пробелы в начале и в конце.
		name = strings.TrimSpace(input)

		// Verification : uniquement des lettres
		// Проверка: разрешены только буквы.
		valid := true

		// Parcourt chaque caractère du nom.
		// Перебирает каждый символ имени.
		for _, char := range name {

			// Vérifie si le caractère est une lettre.
			// Проверяет, является ли символ буквой.
			if (char < 'a' || char > 'z') && (char < 'A' || char > 'Z') {
				valid = false
				break
			}
		}

		// Vérifie que le nom est valide et qu'il n'est pas vide.
		// Проверяет, что имя правильное и не пустое.
		if valid && len(name) > 0 {
			break // Le nom est valide, on sort de la boucle
			// Имя правильное, поэтому выходим из цикла.
		}

		fmt.Println("Nom invalide ! utilisez uniquement des lettres sans espaces ni chiffres.")
	}

	// 2. Formatage : Première lettre en Majuscule, le reste en minuscules
	// 2. Форматирование: первая буква большая, остальные маленькие.
	name = strings.Title(strings.ToLower(name))

	// 3. Choix de la classe*
	// 3. Выбор класса персонажа.
	var className string
	var MaxHP int

	// Répète le menu jusqu'à ce qu'une classe correcte soit choisie.
	// Повторяет меню, пока не будет выбран правильный класс.
	for {
		fmt.Println("\nChoisissez votre classe :")
		fmt.Println("1. Humain (100 PV)")
		fmt.Println("2. Elfe (80 PV)")
		fmt.Println("3. Nain (120 PV)")
		fmt.Print("Votre choix : ")

		// Stocke le choix du joueur.
		// Хранит выбор игрока.
		var choice int
		fmt.Scan(&choice)

		// Définit la classe et les PV maximum selon le choix.
		// Устанавливает класс и максимальное здоровье в зависимости от выбора.
		if choice == 1 {
			className = "Humain"
			MaxHP = 100
			break // Choix valide, on sort de la boucle !
			// Выбор правильный, выходим из цикла.
		} else if choice == 2 {
			className = "Elfe"
			MaxHP = 80
			break // Choix valide, on sort de la boucle !
			// Выбор правильный, выходим из цикла.
		} else if choice == 3 {
			className = "Nain"
			MaxHP = 120
			break // Choix valide, on sort de la boucle !
			// Выбор правильный, выходим из цикла.
		} else {
			// Affiche un message si le choix est incorrect.
			// Показывает сообщение, если выбор неправильный.
			fmt.Println("\nChoix invalide ! Veuillez rechoisir une classe parmi les options indiquées.")
		}
	}

	// 4. Initialisation avec 50% des PV max et inventaire de base
	// 4. Создание персонажа с 50% максимального здоровья и начальным инвентарём.
	return Character{
		Name:      name,
		Class:     className,
		Level:     1,
		MaxHP:     MaxHP,
		CurrentHP: MaxHP / 2, // 50% des PV Max
		// Текущее здоровье равно 50% от максимального.
		Gold:         100,
		Inventory:    []string{"Potion de vie", "Potion de vie", "Potion de vie"},
		Skill:        []string{"Coup de poing"},
		MaxInventory: 10,
		UpgradeCount: 0,
	}
}

// Affiche les informations du personnage.
// Показывает информацию о персонаже.
func displayInfo(player Character) {

	// Affiche les statistiques principales du personnage.
	// Показывает основные характеристики персонажа.
	fmt.Println("Name:", player.Name)
	fmt.Println("Class:", player.Class)
	fmt.Println("Level:", player.Level)
	fmt.Println("Gold:", player.Gold)
	fmt.Println("HP:", player.CurrentHP, "/", player.MaxHP)

	// Affiche les équipements actuellement portés.
	// Показывает экипировку, которая сейчас надета.
	fmt.Println("\n--- ÉQUIPEMENTS PORTÉS ---")
	fmt.Printf("Tête  : %s\n", showEquip(player.Equip.Head))
	fmt.Printf("Torse : %s\n", showEquip(player.Equip.Body))
	fmt.Printf("Pieds : %s\n", showEquip(player.Equip.Feet))
	fmt.Println()
}

// Retourne le nom de l'équipement ou "Aucun" si l'emplacement est vide.
// Возвращает название экипировки или "Aucun", если слот пустой.
func showEquip(item string) string {

	// Vérifie si l'emplacement d'équipement est vide.
	// Проверяет, пустой ли слот экипировки.
	if item == "" {
		return "Aucun"
	}

	// Retourne le nom de l'objet équipé.
	// Возвращает название надетого предмета.
	return item
}

// Vérifie si le personnage est mort et lui rend 50 % de ses PV maximum.
// Проверяет, умер ли персонаж, и восстанавливает ему 50% максимального здоровья.
func isDead(c *Character) bool {

	// Vérifie si les PV du personnage sont inférieurs ou égaux à zéro.
	// Проверяет, меньше или равно ли здоровье персонажа нулю.
	if c.CurrentHP <= 0 {
		fmt.Println("\nVous êtes mort...")

		c.CurrentHP = c.MaxHP / 2 // Résurrection à 50% des PV max (Tâche 8)
		// Воскрешает персонажа с 50% максимального здоровья.

		fmt.Printf("Vous avez été ressuscité avec %d/%d PV !\n", c.CurrentHP, c.MaxHP)

		return true // Le joueur était bien mort
		// Возвращает true, потому что игрок был мёртв.
	}

	return false // Le joueur est encore en vie
	// Возвращает false, потому что игрок всё ещё жив.
}

// Définit les emplacements disponibles pour les équipements.
// Определяет доступные слоты для экипировки.
type Equipment struct {
	Head string // Chapeau
	// Голова — шляпа.
	Body string // Tunique
	// Тело — туника.
	Feet string // Bottes
	// Ноги — ботинки.
}
