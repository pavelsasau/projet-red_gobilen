package main

import "fmt"

// Parcourt et affiche l'inventaire
// Перебирает и показывает инвентарь
func accessInventory(c *Character) {

	// Répète le menu de l'inventaire jusqu'au retour.
	// Повторяет меню инвентаря, пока игрок не выйдет назад.
	for {
		fmt.Println("\n=== INVENTAIRE ===")

		// Vérifie si l'inventaire est vide.
		// Проверяет, пустой ли инвентарь.
		if len(c.Inventory) == 0 {
			fmt.Println("Votre inventaire est vide.")
			return
		}

		// Affichage des objets avec un numéro
		// Показывает предметы с их номерами.
		for i, item := range c.Inventory {
			fmt.Printf("%d. %s\n", i+1, item)
		}

		fmt.Println("0. Retour")
		fmt.Print("Choisissez un objet à utiliser/équiper : ")

		// Stocke le choix du joueur.
		// Хранит выбор игрока.
		var choice int
		fmt.Scan(&choice)

		// Retourne au menu précédent si le joueur choisit 0.
		// Возвращает назад, если игрок выбирает 0.
		if choice == 0 {
			return
		}

		// Vérification du choix valide
		// Проверка правильности выбора.
		if choice > 0 && choice <= len(c.Inventory) {

			// Récupère l'objet choisi dans l'inventaire.
			// Получает выбранный предмет из инвентаря.
			selectedItem := c.Inventory[choice-1]

			// Utilise l'objet sélectionné.
			// Использует выбранный предмет.
			useItem(c, selectedItem)
		} else {
			fmt.Println("\nChoix invalide.")
		}
	}
}

// Ajoute un objet à l'inventaire
// Добавляет предмет в инвентарь.
func addInventory(c *Character, item string) bool {

	// Vérifie si l'inventaire a atteint sa capacité maximale.
	// Проверяет, достиг ли инвентарь максимальной вместимости.
	if len(c.Inventory) >= c.MaxInventory { // Limite d'inventaire (Tâche 12)
		// Ограничение вместимости инвентаря (Задание 12).
		fmt.Println("Inventaire plein ! Impossible d'ajouter l'objet.")

		return false // Échec de l'ajout
		// Возвращает false, потому что предмет добавить не удалось.
	}

	// Ajoute l'objet à la fin de l'inventaire.
	// Добавляет предмет в конец инвентаря.
	c.Inventory = append(c.Inventory, item)

	return true // Succès de l'ajout
	// Возвращает true, потому что предмет успешно добавлен.
}

// Retire un objet de l'inventaire
// Удаляет предмет из инвентаря.
func removeInventory(c *Character, item string) bool {

	// Parcourt tous les objets de l'inventaire.
	// Перебирает все предметы в инвентаре.
	for i, slot := range c.Inventory {

		// Vérifie si l'objet actuel correspond à l'objet recherché.
		// Проверяет, совпадает ли текущий предмет с нужным.
		if slot == item {

			// Supprime l'élément à l'index i
			// Удаляет элемент с индексом i.
			c.Inventory = append(c.Inventory[:i], c.Inventory[i+1:]...)

			// L'objet a été trouvé et supprimé.
			// Предмет найден и удалён.
			return true
		}
	}

	// Affiche un message si l'objet n'a pas été trouvé.
	// Показывает сообщение, если предмет не найден.
	fmt.Printf("L'objet %s n'a pas été trouvé dans l'inventaire.\n", item)

	// Retourne false car aucun objet n'a été supprimé.
	// Возвращает false, потому что предмет не был удалён.
	return false
}

// Vérifie si le joueur possède un objet en quantité suffisante sans le retirer
// Проверяет, есть ли у игрока нужное количество предметов, не удаляя их.
func hasItem(c *Character, item string, count int) bool {

	// Compteur du nombre d'objets trouvés.
	// Счётчик найденных предметов.
	found := 0

	// Parcourt tous les objets de l'inventaire.
	// Перебирает все предметы в инвентаре.
	for _, slot := range c.Inventory {

		// Vérifie si l'objet actuel correspond à l'objet recherché.
		// Проверяет, совпадает ли текущий предмет с нужным.
		if slot == item {
			found++
		}
	}

	// Retourne true si la quantité trouvée est suffisante.
	// Возвращает true, если найденного количества достаточно.
	return found >= count
}

// Fonction qui déclenche l'effet d'un objet selon son nom
// Функция запускает действие предмета в зависимости от его названия.
func useItem(c *Character, item string) {

	// Choisit l'action selon le nom de l'objet.
	// Выбирает действие в зависимости от названия предмета.
	switch item {

	case "Potion de vie":
		// Utilise une potion de vie.
		// Использует зелье здоровья.
		takePot(c)

	case "Potion de poison":
		// La potion de poison ne peut pas être utilisée hors combat.
		// Зелье яда нельзя использовать вне боя.
		fmt.Println("\nLa potion de poison ne peut être utilisée que pendant un combat !")

	case "Livre de Sort : Boule de Feu":
		// Apprend le sort Boule de Feu.
		// Изучает заклинание «Огненный шар».
		spellBook(c, "Boule de Feu")

	case "Chapeau de l'aventurier", "Tunique de l'aventurier", "Bottes de l'aventurier":
		// Équipe l'objet sélectionné.
		// Надевает выбранный предмет.
		equipItem(c, item)

	default:
		// Affiche un message si l'objet ne peut pas être utilisé directement.
		// Показывает сообщение, если предмет нельзя использовать напрямую.
		fmt.Printf("\nL'objet '%s' ne peut pas être utilisé directement.\n", item)
	}
}

// Augmente la capacité maximale de l'inventaire.
// Увеличивает максимальную вместимость инвентаря.
func upgradeInventorySlot(c *Character) {

	// Vérifie si le joueur a déjà acheté trois améliorations.
	// Проверяет, купил ли игрок уже три улучшения.
	if c.UpgradeCount >= 3 {
		fmt.Println("\nVous avez atteint la limite maximale d'améliorations d'inventaire (3/3).")
		return
	}

	// Ajoute 10 emplacements à l'inventaire.
	// Добавляет 10 мест в инвентарь.
	c.MaxInventory += 10

	// Augmente le compteur des améliorations.
	// Увеличивает счётчик улучшений.
	c.UpgradeCount++

	// Affiche la nouvelle capacité de l'inventaire.
	// Показывает новую вместимость инвентаря.
	fmt.Printf("\nInventaire agrandi ! Nouvelle capacité : %d emplacements (Améliorations : %d/3).\n", c.MaxInventory, c.UpgradeCount)
}

// Pour l'utilisation d'objet en combat
// Для использования предметов во время боя.
func useItemFight(c *Character, m *Monster, item string) {

	// Choisit l'effet selon l'objet utilisé.
	// Выбирает действие в зависимости от используемого предмета.
	switch item {

	case "Potion de vie":
		// Utilise la potion de vie.
		// Использует зелье здоровья.
		takePot(c)

		// Retire la potion utilisée de l'inventaire.
		// Удаляет использованное зелье из инвентаря.
		removeInventory(c, item)

	case "Potion de poison":
		// Applique le poison au monstre.
		// Накладывает яд на монстра.
		poisonPot(m)

		// Retire la potion utilisée de l'inventaire.
		// Удаляет использованное зелье из инвентаря.
		removeInventory(c, item)

	default:
		// Affiche un message si l'objet ne peut pas être utilisé en combat.
		// Показывает сообщение, если предмет нельзя использовать в бою.
		fmt.Printf("\nL'objet '%s' ne peut pas être utilisé en combat.\n", item)
	}
}

// Affiche l'inventaire disponible pendant un combat.
// Показывает инвентарь, доступный во время боя.
func accessInventoryFight(c *Character, m *Monster) {

	// Vérifie si l'inventaire est vide.
	// Проверяет, пустой ли инвентарь.
	if len(c.Inventory) == 0 {
		fmt.Println("\nVotre inventaire est vide.")
		return
	}

	// Affiche tous les objets disponibles pendant le combat.
	// Показывает все предметы, доступные во время боя.
	fmt.Println("\n=== INVENTAIRE (COMBAT) ===")
	for i, item := range c.Inventory {
		fmt.Printf("%d. %s\n", i+1, item)
	}

	fmt.Println("0. Retour")
	fmt.Print("Choisissez un objet à utiliser : ")

	// Stocke le choix du joueur.
	// Хранит выбор игрока.
	var choice int
	fmt.Scan(&choice)

	// Retourne au combat si le joueur choisit 0.
	// Возвращает в бой, если игрок выбирает 0.
	if choice == 0 {
		return
	}

	// Vérifie que le numéro choisi existe dans l'inventaire.
	// Проверяет, существует ли выбранный номер в инвентаре.
	if choice > 0 && choice <= len(c.Inventory) {

		// Récupère l'objet correspondant au numéro choisi.
		// Получает предмет, соответствующий выбранному номеру.
		selectedItem := c.Inventory[choice-1]

		// Utilise l'objet pendant le combat.
		// Использует предмет во время боя.
		useItemFight(c, m, selectedItem)
	} else {
		fmt.Println("Choix invalide.")
	}
}
