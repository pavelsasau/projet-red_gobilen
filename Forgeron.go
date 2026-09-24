package main

import "fmt"

// Affiche le menu du forgeron et permet au joueur de fabriquer des équipements.
// Показывает меню кузнеца и позволяет игроку создавать экипировку.
func blacksmithMenu(c *Character) {
	for {
		// Affiche le titre du menu, l'or du joueur et les recettes disponibles.
		// Показывает название меню, золото игрока и доступные рецепты.
		fmt.Println("\n=== FORGERON ===")
		fmt.Printf("Votre bourse : %d pièces d'or\n", c.Gold)
		fmt.Println("1. Chapeau de l'aventurier (1 Plume de Corbeau, 1 Cuir de Sanglier + 5 Gold)")
		fmt.Println("2. Tunique de l'aventurier (2 Fourrure de Loup, 1 Peau de Troll + 5 Gold)")
		fmt.Println("3. Bottes de l'aventurier (1 Fourrure de Loup, 1 Cuir de Sanglier + 5 Gold)")
		fmt.Println("0. Retour")
		fmt.Print("Votre choix : ")

		// Lit le choix du joueur.
		// Считывает выбор игрока.
		var choice int
		fmt.Scan(&choice)

		// Exécute une action selon le choix du joueur.
		// Выполняет действие в зависимости от выбора игрока.
		switch choice {
		case 1:
			// Fabrique le chapeau avec les matériaux indiqués dans la recette.
			// Создаёт шляпу из материалов, указанных в рецепте.
			craftItem(c, "Chapeau de l'aventurier", map[string]int{"Plume de Corbeau": 1, "Cuir de Sanglier": 1})

		case 2:
			// Fabrique la tunique avec les matériaux indiqués dans la recette.
			// Создаёт тунику из материалов, указанных в рецепте.
			craftItem(c, "Tunique de l'aventurier", map[string]int{"Fourrure de Loup": 2, "Peau de Troll": 1})

		case 3:
			// Fabrique les bottes avec les matériaux indiqués dans la recette.
			// Создаёт ботинки из материалов, указанных в рецепте.
			craftItem(c, "Bottes de l'aventurier", map[string]int{"Fourrure de Loup": 1, "Cuir de Sanglier": 1})

		case 0:
			fmt.Println("Vous quittez le Forgeron.")
			return // Retour au menu principal
			// Возврат в главное меню.

		default:
			// Affiche un message si le choix n'existe pas.
			// Показывает сообщение, если такого варианта выбора нет.
			fmt.Println("Choix invalide, réessayez.")
		}
	}
}

// Fabrique un objet si le joueur possède assez d'or et de matériaux.
// Создаёт предмет, если у игрока достаточно золота и материалов.
func craftItem(c *Character, itemName string, recipe map[string]int) {
	// 1. Vérification de l'or
	// 1. Проверка количества золота.
	if c.Gold < 5 {
		fmt.Println("\nVous n'avez pas assez d'or pour forger ! (Requis : 5 Gold)")
		return
	}

	// 2. Vérification des matériaux
	// 2. Проверка необходимых материалов.
	missing := false

	// Parcourt chaque matériau de la recette et sa quantité nécessaire.
	// Перебирает каждый материал из рецепта и необходимое количество.
	for item, count := range recipe {

		// Vérifie si le joueur possède assez de ce matériau.
		// Проверяет, есть ли у игрока достаточное количество этого материала.
		if !hasItem(c, item, count) {
			fmt.Printf("\nIl vous manque : %d x %s !", count, item)
			missing = true
		}
	}

	// Si au moins un matériau manque, on arrête la fabrication
	// Если не хватает хотя бы одного материала, создание предмета останавливается.
	if missing {
		fmt.Println()
		return
	}

	// 3. Retrait des matériaux de l'inventaire
	// 3. Удаление использованных материалов из инвентаря.
	for item, count := range recipe {

		// Retire le matériau autant de fois que demandé dans la recette.
		// Удаляет материал столько раз, сколько указано в рецепте.
		for i := 0; i < count; i++ {
			removeInventory(c, item)
		}
	}

	// 4. Retrait de l'or
	// 4. Списание золота.
	c.Gold -= 5

	// 5. Ajout de l'équipement fabriqué
	// 5. Добавление созданной экипировки в инвентарь.
	if !addInventory(c, itemName) {

		// En cas d'inventaire plein au dernier moment, remboursement
		// Если инвентарь оказался заполнен, золото возвращается игроку.
		c.Gold += 5
		fmt.Println("Fabrication annulée : inventaire plein.")
	} else {

		// Affiche un message si la fabrication a réussi.
		// Показывает сообщение, если создание предмета прошло успешно.
		fmt.Printf("\nFabrication réussie ! Vous obtenez : %s !\n", itemName)
	}
}
