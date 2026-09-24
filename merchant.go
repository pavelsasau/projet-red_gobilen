package main

import "fmt"

// Affiche le menu du marchand et permet au joueur d'acheter des objets.
// Показывает меню торговца и позволяет игроку покупать предметы.
func merchantMenu(c *Character) {

	// Répète le menu du marchand jusqu'au retour.
	// Повторяет меню торговца, пока игрок не выйдет назад.
	for {
		fmt.Println("\n=== MARCHAND ===")

		// Affiche la quantité d'or du joueur.
		// Показывает количество золота у игрока.
		fmt.Printf("Votre bourse : %d pièces d'or\n", c.Gold)

		// Affiche les objets disponibles à l'achat.
		// Показывает предметы, доступные для покупки.
		fmt.Println("1. Potion de vie (3 pièces d'or)")
		fmt.Println("2. Potion de poison (6 pièces d’or)")
		fmt.Println("3. Livre de Sort : Boule de Feu (25 pièces d’or)")
		fmt.Println("4. Fourrure de Loup (4 pièces d’or)")
		fmt.Println("5. Peau de Troll (7 pièces d’or)")
		fmt.Println("6. Cuir de Sanglier (3 pièces d’or)")
		fmt.Println("7. Plume de Corbeau (1 pièce d’or)")
		fmt.Println("8. Augmentation d'inventaire (30 pièces d'or)")
		fmt.Println("0. Retour")
		fmt.Print("Votre choix : ")

		// Stocke le choix du joueur.
		// Хранит выбор игрока.
		var choice int
		fmt.Scan(&choice)

		// Exécute l'achat correspondant au choix du joueur.
		// Выполняет покупку в зависимости от выбора игрока.
		switch choice {
		case 1:
			buyItems(c, "Potion de vie", 3)

		case 2:
			buyItems(c, "Potion de poison", 6)

		case 3:
			buyItems(c, "Livre de Sort : Boule de Feu", 25)

		case 4:
			buyItems(c, "Fourrure de Loup", 4)

		case 5:
			buyItems(c, "Peau de Troll", 7)

		case 6:
			buyItems(c, "Cuir de Sanglier", 3)

		case 7:
			buyItems(c, "Plume de Corbeau", 1)

		case 8:
			buyUpgrade(c, 30)

		case 0:
			fmt.Println("Vous quittez le marchand.")
			return // Retour au menu principal
			// Возврат в главное меню.

		default:
			// Affiche un message si le choix n'est pas valide.
			// Показывает сообщение, если выбор неправильный.
			fmt.Println("Choix invalide, réessayez.")
		}
	}
}

// Fonction générique qui gère l'achat, la déduction d'or et l'ajout à l'inventaire
// Универсальная функция, которая управляет покупкой, списанием золота и добавлением предмета в инвентарь.
func buyItems(c *Character, itemName string, price int) {

	// 1. Vérification si le joueur a assez d'or
	// 1. Проверка, достаточно ли у игрока золота.
	if c.Gold < price {
		fmt.Printf("\nVous n'avez pas assez d'or pour acheter %s ! (Prix : %d Gold, Solde : %d Gold)\n", itemName, price, c.Gold)
		return
	}

	// 2. Ajout de l'objet à l'inventaire
	// 2. Добавление предмета в инвентарь.
	if addInventory(c, itemName) {

		c.Gold -= price // 3. Déduction de l'or si l'objet peut être ajouter a l'inventaire
		// 3. Списание золота, если предмет удалось добавить в инвентарь.

		// 4. Confirmation de l'achat
		// 4. Подтверждение успешной покупки.
		fmt.Printf("\nAchat réussi : %s pour %d pièces d'or ! Solde restant : %d Gold.\n", itemName, price, c.Gold)

	} else {

		// Affiche un message si l'inventaire est plein.
		// Показывает сообщение, если инвентарь заполнен.
		fmt.Println("L'achat a été annulé car votre inventaire est plein.")
	}
}

// Achète une amélioration de la capacité de l'inventaire.
// Покупает улучшение вместимости инвентаря.
func buyUpgrade(c *Character, price int) {

	// Vérifie si le joueur a déjà acheté le nombre maximum d'améliorations.
	// Проверяет, купил ли игрок уже максимальное количество улучшений.
	if c.UpgradeCount >= 3 {
		fmt.Println("\nLe marchand vous dit : 'Je ne peux plus agrandir votre sac, vous avez atteint la limite (3/3) !'")
		return
	}

	// Vérifie si le joueur possède assez d'or.
	// Проверяет, достаточно ли у игрока золота.
	if c.Gold < price {
		fmt.Printf("\nVous n'avez pas assez d'or ! (Prix : %d Gold, Solde : %d Gold)\n", price, c.Gold)
		return
	}

	// Retire le prix de l'amélioration de l'or du joueur.
	// Вычитает стоимость улучшения из золота игрока.
	c.Gold -= price

	// Augmente la capacité de l'inventaire.
	// Увеличивает вместимость инвентаря.
	upgradeInventorySlot(c)
}
